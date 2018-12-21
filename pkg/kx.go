package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

const (
	KubeConfigEnvVar      = "KUBECONFIG"
	DefaultKubeConfigPath = "~/.kube/config"
	KxConfigEnvVar        = "KXCONFIG"
	DefaultKxConfigPath   = "~/.kx.yaml"
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

func GetKubeConfigPath() string {
	return getPathFromEnv(KubeConfigEnvVar, DefaultKubeConfigPath)
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

func GetDefaultKubeConfig() (*KubeConfig, error) {
	path := getPathFromEnv("", DefaultKubeConfigPath)
	return GetKubeConfig(path)
}

func GetKxConfigPath() string {
	return getPathFromEnv(KxConfigEnvVar, DefaultKxConfigPath)
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

func getEnvOrDefault(envVar, defaultVal string) string {
	var val string
	var ok bool

	if val, ok = os.LookupEnv(envVar); !ok {
		val = defaultVal
	}
	return val
}

func getPathFromEnv(envVar, defaultValue string) string {
	var filePath string
	var err error

	filePath = getEnvOrDefault(envVar, defaultValue)
	if filePath, err = homedir.Expand(filePath); err != nil {
		fmt.Printf("error loading config from path %s set in %s: %s\n", filePath, envVar, err.Error())
		os.Exit(1)
	}
	return filePath
}
