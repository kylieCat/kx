// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"sort"

	"github.com/gosuri/uitable"
	"github.com/kylie-a/kx/pkg"
)

func validateArgs(args []string) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) > 2 {
		return errors.New("too many args")
	}
	if len(args) == 1 && args[0] == placeholder {
		contextName = kxConf.Previous.Context
		ns = kxConf.Previous.Namespace
		return nil
	}
	if len(args) >= 1 {
		if args[0] == placeholder {
			contextName = kubeConf.CurrentContext
		} else {
			contextName = args[0]
		}
	}
	if len(args) == 2 {
		if args[1] != placeholder {
			ns = args[1]
		} else {
			var err error
			if ns, err = kubeConf.GetNamespaceForContext(contextName); err != nil {
				return err
			}
		}
	}
	return nil
}

func list(conf *pkg.KubeConfig) {
	table := uitable.New()
	table.AddRow("Current", "Context", "Namespace", "Favorites")
	marker := ""
	sort.Sort(conf.Contexts)
	for _, ctx := range conf.Contexts {
		if ctx.Name == conf.CurrentContext {
			marker = "-->"
		} else {
			marker = ""
		}
		favesForCtx := kxConf.FavoritesForContext(ctx.Name)
		table.AddRow(marker, ctx.Name, ctx.ContextConf.Namespace, favesForCtx.String())
	}
	fmt.Println(table)
}

func updateContext(contextName, ns string, useSet bool) error {
	var currentCtx pkg.CtxNsPair
	var err error

	if currentCtx, err = kubeConf.GetCurrentContextAndNamespace(); err != nil {
	    return err
	}
	if useSet {
		kubeConf.SetContext(contextName)
	} else {
		if err := kubeConf.UseContext(contextName); err != nil {
			ctxPair, ok := kxConf.GetFavorite(contextName)
			if !ok {
				return err
			}
			if err = kubeConf.UseContext(ctxPair.Context); err != nil {
				return errors.New("error setting context from favorite: " + err.Error())
			}
			if err = kubeConf.SetNamespaceForContext(ctxPair.Context, ctxPair.Namespace); err != nil {
				return errors.New("error setting namespace from favorite: " + err.Error())
			}
		}
	}
	if ns != "" {
		if err := kubeConf.SetNamespaceForContext(contextName, ns); err != nil {
			return err
		}
	}
	if err = kxConf.SetPrevious(currentCtx); err != nil {
		return err
	}
	return kubeConf.Save(kubePath)
}


func returnToPrevious() error {
	var ctxPair pkg.CtxNsPair
	var err error

	if ctxPair, err = kxConf.GetPrevious(); err != nil {
		return err
	}
	if err = kubeConf.UseContext(ctxPair.Context); err != nil {
		return errors.New("error setting context from previous: " + err.Error())
	}
	if err = kubeConf.SetNamespaceForContext(ctxPair.Context, ctxPair.Namespace); err != nil {
		return errors.New("error setting namespace from previous: " + err.Error())
	}
	return kubeConf.Save(kubePath)
}
