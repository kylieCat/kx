package pkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type KxConfig struct {
	Path      string
	Prompt    Prompt    `yaml:"prompt" json:"prompt"`
	Favorites Favorites `json:"favorites" yaml:"favorites"`
	Previous  CtxNsPair `json:"previous" yaml:"previous"`
	changed   bool
}

func (kx KxConfig) Changed() bool {
	return kx.changed
}

func (kx KxConfig) Save(path string) error {
	var fileContents []byte
	var err error

	if fileContents, err = yaml.Marshal(kx); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, fileContents, 0644); err != nil {
		return err
	}
	return nil
}

func (kx *KxConfig) AddFavorite(name, ctx, ns string) error {
	kx.Favorites.AddFavorite(name, ctx, ns)
	kx.changed = true
	return nil
}

func (kx *KxConfig) GetFavorite(name string) (CtxNsPair, bool) {
	if kx == nil {
		return CtxNsPair{}, false
	}
	return kx.Favorites.GetFavorite(name)
}

func (kx *KxConfig) IsFavorite(ctx, ns string) (string, bool) {
	if kx == nil {
		return "", false
	}
	return kx.Favorites.IsFavorite(ctx, ns)
}

func (kx *KxConfig) FavoritesForContext(name string) Favorites {
	if kx == nil {
		return make(Favorites)
	}
	return kx.Favorites.FavoritesForContext(name)
}
