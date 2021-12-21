package cert_apply

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"os"
)

const (
	PrivateKeyFilesPath = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/keystore/"
	PublicKeyFilesPath  = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/keystore/"
	SignCertFilesPath   = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/signcerts/"
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
	file, err := os.Create(PrivateKeyFilesPath + "client_private_key.pem")
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

	file, err = os.Create(PublicKeyFilesPath + "client_public_key.pem")
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
			Country:            []string{"CH"},
			Province:           []string{"Beijing"},
			Locality:           []string{"Beijing"},
			Organization:       []string{"colla"},
			OrganizationalUnit: []string{"node"},
			CommonName:         "127.0.0.1",
		},
		//	IPAddresses:    []net.IP{net.IPv4(127, 0, 0, 1)},
		// DNSNames:       []string{"colla.node.com"},
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
