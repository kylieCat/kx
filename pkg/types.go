package pkg

import (
	"github.com/kylie-a/kx/pkg/colors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ClusterConf struct {
	CertificateAuthorityData interface{} `yaml:"certificate-authority-data"`
	Server                   interface{} `yaml:"server"`
}

type Cluster struct {
	ClusterConf ClusterConf `yaml:"cluster"`
	Name        string      `yaml:"name"`
}

type Clusters []Cluster

func (cl Clusters) Len() int           { return len(cl) }
func (cl Clusters) Swap(i, j int)      { cl[i], cl[j] = cl[j], cl[i] }
func (cl Clusters) Less(i, j int) bool { return cl[i].Name < cl[j].Name }

type ContextConf struct {
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
	User      string `yaml:"user"`
}

type Context struct {
	ContextConf ContextConf `yaml:"context"`
	Name        string      `yaml:"name"`
}

type Contexts []*Context

func (contexts Contexts) Len() int           { return len(contexts) }
func (contexts Contexts) Swap(i, j int)      { contexts[i], contexts[j] = contexts[j], contexts[i] }
func (contexts Contexts) Less(i, j int) bool { return contexts[i].Name < contexts[j].Name }

func (contexts Contexts) GetContext(name string) (*Context, error) {
	for _, ctx := range contexts {
		if ctx.Name == name {
			return ctx, nil
		}
	}
	return nil, NewNoSuchContextError(name)
}

func (contexts Contexts) SetNamespaceForContext(ctxName, ns string) error {
	var ctx *Context
	var err error

	if ctx, err = contexts.GetContext(ctxName); err != nil {
		return err
	}
	ctx.ContextConf.Namespace = ns
	return nil
}

func (contexts Contexts) GetNamespaceForContext(name string) (string, error) {
	var ctx *Context
	var err error

	if ctx, err = contexts.GetContext(name); err != nil {
		return "", err
	}
	return ctx.ContextConf.Namespace, nil
}

type Preferences struct{}

type AuthProviderConfig struct {
	AccessToken string `yaml:"access-token"`
	CmdArgs     string `yaml:"cmd-args"`
	CmdPath     string `yaml:"cmd-path"`
	Expiry      string `yaml:"expiry"`
	ExpiryKey   string `yaml:"expiry-key"`
	TokenKey    string `yaml:"token-key"`
}

type AuthProvider struct {
	Config AuthProviderConfig `yaml:"config"`
	Name   string             `yaml:"name"`
}

type UserConf struct {
	AuthProvider      AuthProvider `yaml:"auth-provider"`
	Password          string       `yaml:"password"`
	Username          string       `yaml:"username"`
	Token             string       `yaml:"token"`
	ClientCertificate string       `yaml:"clientCertificate"`
	ClientKey         string       `yaml:"clientKey"`
}

type User struct {
	Name     string   `yaml:"name"`
	UserConf UserConf `yaml:"user"`
}

type Users []User

func (u Users) Len() int           { return len(u) }
func (u Users) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u Users) Less(i, j int) bool { return u[i].Name < u[j].Name }

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

type Prompt struct {
	Separator      string `yaml:"separator" json:"separator"`
	LeftWrapper    string `yaml:"leftWrapper" json:"leftWrapper"`
	RightWrapper   string `yaml:"rightWrapper" json:"rightWrapper"`
	ContextColor   string `yaml:"contextColor" json:"contextColor"`
	NamespaceColor string `yaml:"namespaceColor" json:"namespaceColor"`
	SeparatorColor string `yaml:"separatorColor" json:"separatorColor"`
	ColorOff       string `json:"colorOff" yaml:"colorOff"`
}

func (p Prompt) FillColors() Prompt {
	p.ContextColor = colors.Colors.Get(p.ContextColor)
	p.NamespaceColor = colors.Colors.Get(p.NamespaceColor)
	p.SeparatorColor = colors.Colors.Get(p.SeparatorColor)
	return p
}

type Favorite struct {
	Context   string `json:"context" yaml:"context"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

type Favorites map[string]Favorite

type KxConfig struct {
	Prompt    Prompt    `yaml:"prompt" json:"prompt"`
	Favorites Favorites `json:"favorites" yaml:"favorites" `
}
