package issue

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

type IssueCert struct {
	Cert       *x509.Certificate
	PrivateKey *rsa.PrivateKey
}

/*
	const CrtsFilePath = "/home/fabric/GolandProjects/BlockchainDataColla/ca_server/msp/clientcrt/"
*/

const (

	// CrtsFilePath = "E:\\projects\\BlockchainDataColla\\ca_server\\msp\\clientcrt\\"
	CrtsFilePath = "/home/furad/GolandProjects/BlockchainDataColla/caServer/msp/clientcrt/"
	// CaFilePath = "E:\\projects\\BlockchainDataColla\\ca_server\\msp\\signcert\\ca.pem"
	CaFilePath = "/home/furad/GolandProjects/BlockchainDataColla/caServer/msp/signcert/ca.pem"
)

func (i *IssueCert) GetPublicKey(PublicFile string) *x509.Certificate {

	publicKeyFile, err := ioutil.ReadFile(PublicFile)
	if err != nil {
		return nil
	}

	pemBlock, _ := pem.Decode(publicKeyFile)
	if pemBlock == nil {
		return nil
	}

	caCRT, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		panic(err)
	}

	i.Cert = caCRT
	return caCRT

}

func (i *IssueCert) GetPrivateKey(PrivateFile string) *rsa.PrivateKey {
	privateKeyFile, err := ioutil.ReadFile(PrivateFile)
	if err != nil {
		return nil
	}
	pemBlock, _ := pem.Decode(privateKeyFile)
	if pemBlock == nil {
		return nil
	}
	caPrivateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	i.PrivateKey = caPrivateKey
	return caPrivateKey
}

func (i *IssueCert) GetCrsToBytes(CrsFilePath string) (error, []byte) {

	crsFile, err := ioutil.ReadFile(CrsFilePath)
	if err != nil {
		return err, nil
	}
	return nil, crsFile
}

func (i *IssueCert) CrsCreateCrt(crsFile []byte, cn string) (error, []byte) {

	clientCSR, err := x509.ParseCertificateRequest(crsFile)
	if err != nil {
		return err, nil
	}
	if err = clientCSR.CheckSignature(); err != nil {
		return err, nil
	}

	// create client certificate template
	clientCRTTemplate := x509.Certificate{
		Signature:          clientCSR.Signature,
		SignatureAlgorithm: clientCSR.SignatureAlgorithm,

		PublicKeyAlgorithm: clientCSR.PublicKeyAlgorithm,
		PublicKey:          clientCSR.PublicKey,

		SerialNumber:   big.NewInt(3),
		Issuer:         i.Cert.Subject,
		Subject:        clientCSR.Subject,
		EmailAddresses: clientCSR.EmailAddresses,
		NotBefore:      time.Now(),
		NotAfter:       time.Now().Add(365 * 24 * time.Hour),
		IsCA:           false,
		IPAddresses:    clientCSR.IPAddresses,
		DNSNames:       clientCSR.DNSNames,
		KeyUsage:       x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}

	clientCRTRaw, err := x509.CreateCertificate(rand.Reader, &clientCRTTemplate, i.Cert, clientCSR.PublicKey, i.PrivateKey)
	if err != nil {
		return err, nil
	}

	// save the certificate
	crsFilePath := CrtsFilePath + cn + ".crt"
	clientCRTFile, err := os.Create(crsFilePath)
	defer clientCRTFile.Close()
	if err != nil {
		return err, nil
	}
	pem.Encode(clientCRTFile, &pem.Block{Type: "CERTIFICATE", Bytes: clientCRTRaw})

	return nil, clientCRTRaw

}

func (i IssueCert) GetCaCrt() (error, []byte) {
	caBytes, err := ioutil.ReadFile(CaFilePath)
	if err != nil {
		return err, nil
	}
	return nil, caBytes
}
