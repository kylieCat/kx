package pkg

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var defaultKxConfig = KxConfig{
	Prompt: Prompt{
		Separator:      ":",
		LeftWrapper:    "[",
		RightWrapper:   "]",
		ContextColor:   "K8sBlue",
		NamespaceColor: "K8sBlue",
		SeparatorColor: "White",
		ColorOff:       "\033[0m",
	},
	Favorites: make(Favorites),
	Previous:  CtxNsPair{},
	// Set as true by default so the config is saved after first use
	changed: true,
}

func GetKubeConfig(path string) (*KubeConfig, error) {
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

func GetKxConfig(path string) *KxConfig {
	var rawFile []byte
	rawFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println()
		return &defaultKxConfig
	}
	var kxConfig KxConfig
	err = yaml.Unmarshal(rawFile, &kxConfig)
	if err != nil {
		return &defaultKxConfig
	}
	if kxConfig.Favorites == nil {
		kxConfig.Favorites = make(Favorites)
	}
	return &kxConfig
}
