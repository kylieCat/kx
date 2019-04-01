package pkg

import "github.com/kylie-a/kx/pkg/colors"

type Prompt struct {
	Separator      string `yaml:"separator" json:"separator"`
	LeftWrapper    string `yaml:"leftWrapper" json:"leftWrapper"`
	RightWrapper   string `yaml:"rightWrapper" json:"rightWrapper"`
	ContextColor   string `yaml:"contextColor" json:"contextColor"`
	NamespaceColor string `yaml:"namespaceColor" json:"namespaceColor"`
	SeparatorColor string `yaml:"separatorColor" json:"separatorColor"`
	ColorOff       string
}

func (p Prompt) FillColors() Prompt {
	p.ContextColor = colors.Colors.Get(p.ContextColor)
	p.NamespaceColor = colors.Colors.Get(p.NamespaceColor)
	p.SeparatorColor = colors.Colors.Get(p.SeparatorColor)
	p.ColorOff = colors.ColorOff
	return p
}


