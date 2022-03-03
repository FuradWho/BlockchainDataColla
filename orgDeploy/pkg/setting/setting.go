package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type network struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}
type service struct {
	FabricServerName string `yaml:"fabricServerName"`
	CaServerName     string `yaml:"caServerName"`
}
type path struct {
	PrivateKeyFilesPath string `yaml:"PrivateKeyFilesPath"`
	PublicKeyFilesPath  string `yaml:"PublicKeyFilesPath"`
	SignCertFilesPath   string `yaml:"SignCertFilesPath"`
	ServerCert          string `yaml:"serverCert"`
	ClientKey           string `yaml:"clientKey"`
	ClientCert          string `yaml:"clientCert"`
}

type Config struct {
	Network network `yaml:"network"`
	Service service `yaml:"service"`
	Path    path    `yaml:"path"`
}

var Conf Config

func init() {

	data, err := ioutil.ReadFile("/home/furad/GolandProjects/BlockchainDataColla/orgDeploy/config/conf.yaml")
	if err != nil {
		return
	}
	if err := yaml.Unmarshal(data, &Conf); err != nil { //解析yaml文件
		return
	}
	return
}
