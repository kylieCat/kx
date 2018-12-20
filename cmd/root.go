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
	"os"
	"path"

	"github.com/mitchellh/go-homedir"

	"github.com/kylie-a/kx/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	placeholder       = "-"
	defaultKubeConfig = "~/.kube/config"
	kubeConfigEnvVar  = "KUBECONFIG"
)

var (
	cfgFile     string
	contextName string
	ns          string
	favorite    string
)
var (
	all     bool
	set     bool
	noColor bool
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
		if len(args) == 0 {
			return nil
		}
		if len(args) > 2 {
			return errors.New("too many args")
		}
		if len(args) == 1 && args[0] == placeholder {
			return errors.New("invalid args")
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
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if all {
				list(kubeConf)
			} else {
				if c, n, err := kubeConf.GetCurrentContextAndNamespace(); err != nil {
					return err
				} else {
					fmt.Println(c, n)
				}
			}
		} else {
			if contextName != placeholder {
				if set {
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
						path := getKubePath()
						return kubeConf.Save(path)
					}
				}
			}
			if ns != "" {
				if err := kubeConf.SetNamespaceForContext(contextName, ns); err != nil {
					return err
				}
			}
		}
		if cmd.Flag("favorite").Changed {
			if err := kxConf.AddFavorite(favorite, contextName, ns); err != nil {
				return err
			}
		}
		path, err := homedir.Expand(defaultKubeConfig)
		if err != nil {
			return err
		}
		return kubeConf.Save(path)
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		if kxConf.Changed() {
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			kxPath := path.Join(home, ".kx.yaml")
			if err = kxConf.Save(kxPath); err != nil {
				return err
			}
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
	cobra.OnInitialize(initConfig)
	initKubeConfig()
	initKxConfig()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kx.yaml)")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "Use to print list of contexts and namespace pairs")
	rootCmd.Flags().BoolVarP(&set, "set", "s", false, "Act as 'kubectl config set-context' instead of 'kubectl config use-context'")
	rootCmd.Flags().BoolVarP(&noColor, "no-color", "c", true, "Turn off color")
	rootCmd.Flags().StringVarP(&favorite, "favorite", "f", "", "set a favorite context namespace pair")
}

func getKubePath() string {
	var confPath string
	var ok bool
	var err error

	if confPath, ok = os.LookupEnv(kubeConfigEnvVar); !ok {
		confPath = defaultKubeConfig
	}

	if confPath, err = homedir.Expand(confPath); err != nil {
		fmt.Printf("error loading kube config from path %s: %s", confPath, err.Error())
		os.Exit(1)
	}
	return confPath
}

func initKubeConfig() {
	var confPath string

	confPath = getKubePath()


	kubeConf, _ = pkg.GetKubeConfig(confPath)
}

func initKxConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	kxPath := path.Join(home, ".kx.yaml")
	kxConf, err = pkg.GetKxConfig(kxPath)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".kx" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kx")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("broken config file:", err.Error())
	}
}
