// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/kylie-a/kx/pkg/colors"
	"github.com/spf13/cobra"
)

var (
	ctxColor       string
	nsColor        string
	separatorColor string
)

// setCmd represents the context command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var ok bool

		prompt := kxConf.GetPrompt()
		if cmd.Flag("context").Changed {
			if _, ok = colors.Colors.GetSafe(ctxColor); !ok {
				return errors.New("no such color: " + ctxColor)
			}
			prompt.ContextColor = ctxColor
		}
		if cmd.Flag("ns").Changed {
			if _, ok = colors.Colors.GetSafe(nsColor); !ok {
				return errors.New("no such color: " + nsColor)
			}
			prompt.NamespaceColor = nsColor
		}
		if cmd.Flag("separator").Changed {
			if _, ok = colors.Colors.GetSafe(separatorColor); !ok {
				return errors.New("no such color: " + separatorColor)
			}
			prompt.SeparatorColor = separatorColor
		}

		kxConf.Prompt = *prompt
		return kxConf.Save(kxPath)
	},
}

func init() {
	colorsCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&ctxColor, "context", "c", "", "Change the color of the context in kx config")
	setCmd.Flags().StringVarP(&nsColor, "ns", "n", "", "Change the color of the namespace in kx config")
	setCmd.Flags().StringVarP(&separatorColor, "separator", "s", "", "Change the color of the separator in kx config")
}
