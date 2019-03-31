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
	"fmt"
	"os"

	"github.com/kylie-a/kx/pkg"
	"github.com/spf13/cobra"
)

const (
	placeholder    = "-"
	previousMarker = "--"
)

var (
	cfgFile     string
	contextName string
	ns          string
	favorite    string
	kubePath    string
	kxPath      string
	newName     string
)
var (
	all      bool
	set      bool
	noColor  bool
	previous bool
)

var kubeConf *pkg.KubeConfig
var kxConf *pkg.KxConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kx",
	Short: "A more convenient method for changing contexts and name spaces for kubectl",
	Long: `A more convenient method for changing contexts and name spaces for kubectl

Usage:
kx [context [namespace]] [-a|--all] [-s|--set] [--no-color]

To add to prompt:

$ export PROMPT_COMMAND=$(kx prompt)

Examples:

Change context and namespace
$ kx new-context new-namespace

Use a dash as the context name to stay in the same context but change namespace
$ kx - new-namespace

Change context, keep the namespace set in .kube/config
$ kctx new-context

Show current context and namespace
$ kx
current-context current-namespace

Show all contexts and namespaces
$ kx -a
Current	Context          	Namespace
-->    	foo-eks-prd  	    foo
       	foo-gke-prd      	foo
       	foo-gke-stg      	foo
`,
	Args: func(cmd *cobra.Command, args []string) error {
		return validateArgs(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if all {
				list(kubeConf)
			} else {
				if ctxPair, err := kubeConf.GetCurrentContextAndNamespace(); err != nil {
					return err
				} else {
					fmt.Println(ctxPair)
				}
			}
			return nil
		} else {
			if err := updateContext(contextName, ns, set); err != nil {
				return err
			}
		}

		if cmd.Flag("favorite").Changed {
			if err := kxConf.AddFavorite(favorite, contextName, ns); err != nil {
				return err
			}
		}

		if cmd.Flag("rename").Changed {
			if err := kubeConf.RenameContext(contextName, newName); err != nil {
				return err
			}
			if err := updateContext(newName, "", set); err != nil {
				return err
			}
		}
		return kubeConf.Save(kubePath)
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		if kxConf.Changed() {
			return kxConf.Save(kxPath)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfigPaths, initKubeConfig, initKxConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kx.yaml)")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "Use to print list of contexts and namespace pairs")
	rootCmd.Flags().BoolVarP(&noColor, "no-color", "c", true, "Turn off color")
	rootCmd.Flags().StringVarP(&favorite, "favorite", "f", "", "set a favorite context namespace pair")
	rootCmd.Flags().StringVarP(&newName, "rename", "r", "", "Change the name of a context")
}

func initConfigPaths() {
	kubePath = pkg.GetKubeConfigPath()
	kxPath = pkg.GetKxConfigPath()
}

func initKubeConfig() {
	var err error

	if kubeConf, err = pkg.GetKubeConfig(kubePath); err != nil {
		fmt.Printf("error loading kubeconfig from path %s: %s\n", kubePath, err.Error())
		os.Exit(1)
	}
}

func initKxConfig() {
	kxConf = pkg.GetKxConfig(kxPath)
}
