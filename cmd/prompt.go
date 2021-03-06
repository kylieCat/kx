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
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

const promptFuncTmpl = `#! /usr/bin/env bash

__kctx() {
    if [ -z "$KCTX_COLOR" ]; then
        KCTX_COLOR="\033[0;38;5;38m"
    fi
	info=$(kx)
	context="{{ .ContextColor }}${info%% *}{{ .ColorOff }}"
	separator="{{ .SeparatorColor }}{{ .Separator }}{{ .ColorOff }}"
	namespace="{{ .NamespaceColor }}${info#* }{{ .ColorOff }}"
    printf "{{ .LeftWrapper}}${context}${separator}${namespace}{{ .RightWrapper }}\n"
}`

// setCmd represents the set command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Bash function to use as prompt command to show current context and namespace in your shell.",
	Long: `Bash function to use as prompt command to show current context and namespace in your shell.

Usage:
export PROMPT_COMMAND=$(kx prompt)
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("promptFunc").Parse(promptFuncTmpl))
		return t.Execute(os.Stdout, kxConf.Prompt.FillColors())
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
