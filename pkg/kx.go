package pkg

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func GetConfig(path string) (*KubeConfig, error) {
	var rawFile []byte
	rawFile, err := ioutil.ReadFile(path)
	if err != nil {
	    return nil, err
	}
	var kubeConfig KubeConfig
	err = yaml.Unmarshal(rawFile, &kubeConfig)
	if err != nil {
		return nil, err
	}
	return &kubeConfig, nil
}
