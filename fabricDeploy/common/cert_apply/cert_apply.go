package cert_apply

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"net"
	"os"
)

const (
	/*
		linux path:

		PrivateKeyFilesPath = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/keystore/"
		PublicKeyFilesPath  = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/keystore/"
		SignCertFilesPath   = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/signcert/"
	*/

	PrivateKeyFilesPath = "E:\\projects\\BlockchainDataColla\\fabricDeploy\\msp\\keystore\\"
	PublicKeyFilesPath  = "E:\\projects\\BlockchainDataColla\\fabricDeploy\\msp\\keystore\\"
	SignCertFilesPath   = "E:\\projects\\BlockchainDataColla\\fabricDeploy\\msp\\signcert\\"
)

type Crt struct {
	privateKey     *rsa.PrivateKey
	privateKeyFile string
	publicKey      *rsa.PublicKey
	publicKeyFile  string
}

func (c *Crt) CreatePairKey() error {
	certKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	privateKeyStream := x509.MarshalPKCS1PrivateKey(certKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyStream,
	}
	file, err := os.Create(PrivateKeyFilesPath + "fabric_private_key.pem")
	defer file.Close()
	if err != nil {
		return err
	}

	err = pem.Encode(file, privateKeyBlock)
	if err != nil {
		return err
	}

	publicKey := &certKey.PublicKey
	publicKeyStream, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyStream,
	}

	file, err = os.Create(PublicKeyFilesPath + "fabric_public_key.pem")
	defer file.Close()
	if err != nil {
		return err
	}

	err = pem.Encode(file, publicKeyBlock)
	if err != nil {
		return err
	}

	c.privateKey = certKey
	c.publicKey = publicKey

	return nil
}

func (c *Crt) CreateCSR() ([]byte, error) {

	if c.privateKey == nil || c.publicKey == nil {
		return nil, errors.New("please create the private key and public key first")
	}
	template := &x509.CertificateRequest{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          c.publicKey,
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Province:           []string{"Beijing"},
			Locality:           []string{"Beijing"},
			Organization:       []string{"colla"},
			OrganizationalUnit: []string{"fabric"},
			CommonName:         "colla.fabric.com",
		},
		IPAddresses: []net.IP{
			net.IPv4(127, 0, 0, 1),
			net.IPv4(192, 168, 175, 133),
			net.IPv4(192, 168, 152, 1),
			net.IPv4(192, 168, 0, 95),
		},
		DNSNames:       []string{"colla.fabric.com", "localhost"},
		EmailAddresses: []string{"liu1337543811@gmail.com"},
	}

	csrDER, err := x509.CreateCertificateRequest(rand.Reader, template, c.privateKey)
	if err != nil {
		return nil, err
	}
	return csrDER, nil
}

func (c Crt) SaveCSR(crtBytes []byte) error {
	if c.privateKey == nil || c.publicKey == nil {
		return errors.New("please create the private key and public key first")
	}

	crsFilePath := SignCertFilesPath + "client-ca-cert.crt"
	clientCRTFile, err := os.Create(crsFilePath)
	defer clientCRTFile.Close()
	if err != nil {
		return err
	}
	err = pem.Encode(clientCRTFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: crtBytes,
	})
	if err != nil {
		return err
	}

	return nil
}
