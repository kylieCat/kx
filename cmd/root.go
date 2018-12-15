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

	"github.com/kylie-a/kx/pkg"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	placeholder          = "-"
	defaultKubeConfig    = "~/.kube/config"
	kblue                = "\033[0;38;5;38m"
	colorOptionEnvVar    = "KX_COLOR"
	kubeConfigEnvVar     = "KUBECONFIG"
)

var (
	cfgFile     string
	contextName string
	ns          string
)
var (
	all     bool
	set     bool
	noColor bool
)
var conf *pkg.KubeConfig

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
				contextName = conf.CurrentContext
			} else {
				contextName = args[0]
			}
		}
		if len(args) == 2 {
			if args[1] != placeholder {
				ns = args[1]
			} else {
				var err error
				if ns, err = conf.GetNamespaceForContext(contextName); err != nil {
					return err
				}
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if all {
				list(conf)
			} else {
				if c, n, err := conf.GetCurrentContextAndNamespace(); err != nil {
					return err
				} else {
					fmt.Println(c, n)
				}
			}
			os.Exit(0)
		}
		if contextName != placeholder {
			if set {
				conf.SetContext(contextName)
			} else {
				if err := conf.UseContext(contextName); err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}
		}
		if ns != "" {
			if err := conf.SetNamespaceForContext(contextName, ns); err != nil {
				return err
			}
		}
		path, err := homedir.Expand(defaultKubeConfig)
		if err != nil {
			return err
		}
		return conf.Save(path)
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kx.yaml)")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "Use to print list of contexts and namespace pairs")
	rootCmd.Flags().BoolVarP(&set, "set", "s", false, "Act as 'kubectl config set-context' instead of 'kubectl config use-context'")
	rootCmd.Flags().BoolVarP(&noColor, "no-color", "c", true, "Turn off color")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var confPath string
	var err error

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
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("broken config file:", viper.ConfigFileUsed())
	}

	if confPath, err = homedir.Expand(defaultKubeConfig); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conf, _ = pkg.GetConfig(confPath)
}
