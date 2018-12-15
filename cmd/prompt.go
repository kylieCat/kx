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
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Bash function to use as prompt command to show current context and namespace in your shell.",
	Long: `Bash function to use as prompt command to show current context and namespace in your shell.

Usage:
export PROMPT_COMMAND=$(kx prompt)
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`__kctx() {
    if [ -z "$KCTX_COLOR" ]; then
        KCTX_COLOR="\033[0;38;5;38m"
    fi
    printf "[${KCTX_COLOR}$(kx| awk '{print $1":"$2}')${COLOR_OFF}]\n"
}`)
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
