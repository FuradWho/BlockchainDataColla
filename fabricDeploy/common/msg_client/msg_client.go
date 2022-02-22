package msg_client

import (
	"fmt"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	lcpackager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/lifecycle"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/policydsl"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	channelId        = "mychannel"
	connectConfigDir = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/configs/connect-config/sdk-connection-config.yaml"
	chaincodeId      = "mycc_0"
	chaincodePath    = "newchaincode/test"
	ccVersion        = "0"
	Admin            = "Admin"
)

type FabricClient struct {
	ConnectionFile []byte
	NetworkConfig  fab.NetworkConfig
	ChannelId      string
	GoPath         string

	userName string
	userOrg  string

	resmgmtClients map[string]*resmgmt.Client
	sdk            *fabsdk.FabricSDK
	retry          resmgmt.RequestOption

	ChannelClient *channel.Client
}

func (f *FabricClient) GetNetworkConfig() (fab.NetworkConfig, error) {

	configBackend, _ := f.sdk.Config()
	networkConfig := fab.NetworkConfig{}

	err := lookup.New(configBackend).UnmarshalKey("organizations", &networkConfig.Organizations)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	err = lookup.New(configBackend).UnmarshalKey("orderers", &networkConfig.Orderers)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	err = lookup.New(configBackend).UnmarshalKey("channels", &networkConfig.Channels)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	err = lookup.New(configBackend).UnmarshalKey("peers", &networkConfig.Peers)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	return networkConfig, nil
}

func (f *FabricClient) Init() error {

	connectConfig, _ := ioutil.ReadFile(connectConfigDir)

	f.ConnectionFile = connectConfig
	f.ChannelId = channelId
	f.GoPath = os.Getenv("GOPATH")

	sdk, err := fabsdk.New(config.FromRaw(f.ConnectionFile, "yaml"))
	if err != nil {
		log.Error("Failed to setup main sdk ")
		return err
	}
	f.sdk = sdk

	networkConfig, err := f.GetNetworkConfig()
	if err != nil {
		log.Error("Failed to Get Network Config")
		return err
	}

	resmgmtClients := make(map[string]*resmgmt.Client)
	f.NetworkConfig = networkConfig

	for k, _ := range f.NetworkConfig.Organizations {
		resmgmtClient, err := resmgmt.New(sdk.Context(fabsdk.WithUser(Admin), fabsdk.WithOrg(k)))
		if err != nil {
			log.Errorf("Failed to create channel management client : %s \n", err)
			return err
		}
		resmgmtClients[k] = resmgmtClient
	}

	channelContext := sdk.ChannelContext(channelId, fabsdk.WithUser(Admin))
	channelClient, err := channel.New(channelContext)
	if err != nil {
		log.Printf("Failed to create an new channelClient:%s\n", err)
	}

	f.ChannelClient = channelClient
	f.resmgmtClients = resmgmtClients
	f.retry = resmgmt.WithRetry(retry.DefaultResMgmtOpts)

	return nil
}

func (f *FabricClient) InstallCC(chaincodeId, chaincodePath string) (string, error) {

	desc := &lcpackager.Descriptor{
		Path:  chaincodePath,
		Type:  pb.ChaincodeSpec_GOLANG,
		Label: chaincodeId,
	}

	ccPkg, err := lcpackager.NewCCPackage(desc)
	if err != nil {
		log.Errorf("Failed to NewCCPackage client : %s \n", err)
		return "", err
	}

	installCCReq := resmgmt.LifecycleInstallCCRequest{
		Label:   chaincodeId,
		Package: ccPkg,
	}

	packageID := lcpackager.ComputePackageID(installCCReq.Label, installCCReq.Package)
	ccPolicy := policydsl.SignedByAnyMember([]string{"Org1MSP", "Org2MSP"})

	approveCCReq := resmgmt.LifecycleApproveCCRequest{
		Name:              chaincodeId,
		Version:           "v1.0",
		PackageID:         packageID,
		Sequence:          1,
		EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",
		SignaturePolicy:   ccPolicy,
		InitRequired:      true,
	}

	fmt.Println(packageID)

	for org, client := range f.resmgmtClients {
		if !strings.Contains(org, "order") {
			peers := f.NetworkConfig.Organizations[org].Peers
			for _, peer := range peers {
				fmt.Printf("%s  %s \n", org, peer)
				resp, err := client.LifecycleInstallCC(installCCReq, resmgmt.WithTargetEndpoints(peer), f.retry)
				time.Sleep(time.Second * 10)
				if err != nil {
					log.Errorf("Failed to install chaincode : %s \n", err)
					return "", err
				}
				log.Infof("Package ID: %+v \n", resp)

				txnID, err := client.LifecycleApproveCC(channelId, approveCCReq, resmgmt.WithOrdererEndpoint("orderer.example.com"), resmgmt.WithTargetEndpoints(peer))
				if err != nil {
					log.Errorf("Failed to ApproveCC chaincode : %s \n", err)
					return "", err
				}
				log.Infoln(txnID)
			}
		}
	}

	req1 := resmgmt.LifecycleCheckCCCommitReadinessRequest{
		Name:              chaincodeId,
		Version:           "v1.0",
		EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",
		SignaturePolicy:   ccPolicy,
		Sequence:          1,
		InitRequired:      true,
	}
	resp1, err := f.resmgmtClients["org1"].LifecycleCheckCCCommitReadiness(channelId, req1, resmgmt.WithTargetEndpoints("peer0.org1.example.com"), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		log.Errorf("Failed to LifecycleQueryCommittedCC : %s \n", err)
	}
	log.Infof("%+v \n", resp1)

	time.Sleep(time.Second * 5)

	commitReq := resmgmt.LifecycleCommitCCRequest{
		Name:              chaincodeId,
		Version:           "v1.0",
		Sequence:          1,
		EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",
		SignaturePolicy:   ccPolicy,
		InitRequired:      true,
	}
	txnID, err := f.resmgmtClients["org1"].LifecycleCommitCC(channelId, commitReq, resmgmt.WithTargetEndpoints("peer0.org1.example.com"), resmgmt.WithOrdererEndpoint("orderer.example.com"), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		log.Errorf("Failed to LifecycleCommitCC : %s \n", err)
		return "", err
	}
	log.Infof("%+v \n", txnID)

	return packageID, nil
}
