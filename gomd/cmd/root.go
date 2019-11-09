/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kris-nova/public-speaking/gomd/mdterm"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomd",
	Short: "Play markdown files in your terminal",
	Long: `

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logger.Critical("Invalid directory. Usuage: gomd <presentation-directory> <flags>")
			os.Exit(-2)
		}
		o.dirToParse = args[0]
		err := RunGoMD(o)
		if err != nil {
			logger.Critical("Failure: %v", err)
			os.Exit(-1)
		}
		os.Exit(0)
	},
}

type GoMDOptions struct {
	dirToParse string
	verbosity  int
}

var o = &GoMDOptions{}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&logger.Level, "verbosity", "v", 4, "Verbosity between 0 (off) to 4 (everything)")
}

func RunGoMD(o *GoMDOptions) error {
	p, err := mdterm.LoadPresentation(o.dirToParse)
	if err != nil {
		return fmt.Errorf("unable to load deck: %v", err)
	}
	return p.Run()
}
