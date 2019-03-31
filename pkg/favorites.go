package pkg

import (
	"fmt"
	"strings"
)

type CtxNsPair struct {
	Context   string `json:"context" yaml:"context"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

type Favorites map[string]CtxNsPair

func (favorites Favorites) AddFavorite(name, ctx, ns string) {
	favorites[name] = CtxNsPair{Context: ctx, Namespace: ns}
}

func (favorites Favorites) GetFavorite(name string) (CtxNsPair, bool) {
	pair, ok := favorites[name]
	return pair, ok
}

func (favorites Favorites) IsFavorite(ctx, ns string) (string, bool) {
	for name, favorite := range favorites {
		if favorite.Context == ctx && favorite.Namespace == ns {
			return name, true
		}
	}
	return "", false
}

func (favorites Favorites) FavoritesForContext(name string) Favorites {
	faves := make(Favorites)
	for faveName, favorite := range favorites {
		if favorite.Context == name {
			faves[faveName] = favorite
		}
	}
	return faves
}

func (favorites Favorites) String() string {
	var faves []string

	for faveName, fave := range favorites {
		faves = append(faves, fmt.Sprintf("%s=%s:%s", faveName, fave.Context, fave.Namespace))
	}
	return strings.Join(faves, ",")
}
