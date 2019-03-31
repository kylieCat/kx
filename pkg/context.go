package pkg

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
