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

const completeScript = `function _listkx()
{
	# https://www.gnu.org/software/bash/manual/html_node/Programmable-Completion-Builtins.html
    local cur

    cur=${COMP_WORDS[COMP_CWORD]}
    case ${COMP_CWORD} in
        1)
            COMPREPLY=($(compgen -W "$(kx -a | tail -n +2 |awk 'BEGIN {FS="\t"}; {print $2}')" -- ${cur}))
            ;;
    esac
}

complete -F _listkx kx`

// setCmd represents the set command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Bash completion",
	Long: `To load completion run

	$ eval "$(kx complete)"

To configure your bash shell to load completions for each session add to your bashrc

	# ~/.bashrc or ~/.profile
	eval "$(kx complete)"

Or

	kx complete > ./.kxcomplete; source ./kxcomplete
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(completeScript)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
