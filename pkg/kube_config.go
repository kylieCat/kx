package pkg

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Preferences struct{}

type KubeConfig struct {
	APIVersion     string      `yaml:"apiVersion"`
	Clusters       Clusters    `yaml:"clusters"`
	Contexts       Contexts    `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context"`
	Kind           string      `yaml:"kind"`
	Preferences    Preferences `yaml:"preferences"`
	Users          Users       `yaml:"users"`
	currCtx        string
	currNS         string
}

func (kc KubeConfig) GetCurrentContextName() string {
	return kc.CurrentContext
}

func (kc KubeConfig) GetContext(name string) (*Context, error) {
	return kc.Contexts.GetContext(name)
}

func (kc KubeConfig) GetCurrentContext() (*Context, error) {
	return kc.Contexts.GetContext(kc.CurrentContext)
}

func (kc KubeConfig) GetNamespaceForCurrentContext() (string, error) {
	ctx, err := kc.GetCurrentContext()
	if err != nil {
		return "", err
	}
	return ctx.ContextConf.Namespace, nil
}

func (kc KubeConfig) GetCurrentContextAndNamespace() (string, string, error) {
	ctxName := kc.GetCurrentContextName()
	ns, err := kc.GetNamespaceForCurrentContext()
	if err != nil {
		return "", "", err
	}
	return ctxName, ns, nil
}

func (kc *KubeConfig) SetContext(name string) {
	kc.CurrentContext = name
}

func (kc *KubeConfig) UseContext(name string) error {
	var err error
	if _, err = kc.GetContext(name); err != nil {
		return NewCannotSetContextError(name)
	}
	kc.SetContext(name)
	return nil
}

func (kc *KubeConfig) SetNamespaceForContext(ctxName, ns string) error {
	if err := kc.Contexts.SetNamespaceForContext(ctxName, ns); err != nil {
		return err
	}
	kc.currNS = ns
	return nil
}

func (kc KubeConfig) GetNamespaceForContext(name string) (string, error) {
	return kc.Contexts.GetNamespaceForContext(name)
}

func (kc KubeConfig) Save(path string) error {
	var fileContents []byte
	var err error

	if fileContents, err = yaml.Marshal(kc); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, fileContents, 0644); err != nil {
		return err
	}
	return nil
}

type ContextFilter func(ctx *Context) bool

func (kc KubeConfig) FilterContexts(filter ContextFilter) Contexts {
	contexts := make(Contexts, 0, len(kc.Contexts))
	for _, ctx := range kc.Contexts {
		if filter(ctx) {
			contexts = append(contexts, ctx)
		}
	}
	return contexts
}

type ClusterUserData struct {
	User    User      `yaml:"user"`
	Cluster Cluster   `yaml:"cluster"`
	Ctx     CtxNsPair `yaml:"ctx"`
}

func (c ClusterUserData) YAML() ([]byte, error) {
	var yml []byte
	var err error

	if yml, err = yaml.Marshal(c); err != nil {
		return nil, err
	}
	return yml, nil
}

func (c ClusterUserData) JSON() ([]byte, error) {
	var js []byte
	var err error

	if js, err = json.Marshal(c); err != nil {
		return nil, err
	}
	return js, nil
}

func (kc KubeConfig) GetClusterUserData(ctxName string) (*ClusterUserData, error) {
	var ctx *Context
	var user User
	var cluster Cluster
	var err error

	if ctx, err = kc.Contexts.GetContext(ctxName); err != nil {
		return nil, err
	}

	if user, err = kc.Users.Get(ctx.ContextConf.User); err != nil {
		return nil, err
	}

	if cluster, err = kc.Clusters.Get(ctx.ContextConf.Cluster); err != nil {
		return nil, err
	}
	userClusterData := &ClusterUserData{
		User:    user,
		Cluster: cluster,
		Ctx:     CtxNsPair{Context: ctx.Name, Namespace: ctx.ContextConf.Namespace},
	}
	return userClusterData, nil
}
