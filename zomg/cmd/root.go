/*
Copyright © 2019 Kris Nova <kris@nivenly.com> 2019

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
	"github.com/fatih/color"
	"os"

	"github.com/kris-nova/public-speaking/zomg/mdterm"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zomg",
	Short: "Play markdown files in your terminal",
	Long: `
zomg <directory-of-zomg-files> <zomg-flags>
Example: zomg ~/meeps/myslides --theme basic
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			RunZOMG(o)
		}
		o.dirToParse = args[0]
		err := RunZOMG(o)
		if err != nil {
			logger.Critical("Failure: %v", err)
			os.Exit(-1)
		}
		os.Exit(0)
	},
}

type ZOMGOptions struct {
	dirToParse string
	themeName  string
	theme      mdterm.Theme
	verbosity  int
}

var (
	o        = &ZOMGOptions{}
	themeMap = map[string]struct {
		info       string
		cthemeName mdterm.ThemeName
		theme      mdterm.Theme
	}{
		"nova": {
			info:       "The default theme, specifically designed for Kris Nova's archlinux laptop.",
			cthemeName: mdterm.NovaTheme,
			theme:      &mdterm.ThemeNova{},
		},
		"hack": {
			info:       "Theme designed for dark terminals",
			cthemeName: mdterm.HackTheme,
			theme:      &mdterm.ThemeHack{},
		},
		"basic": {
			info:       "A basic theme, that does no coloring and using your terminals default settings.",
			cthemeName: mdterm.BasicTheme,
			theme:      &mdterm.ThemeBasic{},
		},
	}
)

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
	rootCmd.Flags().StringVarP(&o.themeName, "theme-name", "t", "nova", "The theme to use")

	args := os.Args
	if len(args) == 0 {
		printHelp()
	}
	// ---- Fix Help
	themeStr := "Possible themes:\n\n"
	for n, m := range themeMap {
		newStr := fmt.Sprintf("%s:\n", color.BlueString(n))
		newStr = fmt.Sprintf("%s%s\n\n", newStr, m.info)
		themeStr = fmt.Sprintf("%s%s", themeStr, newStr)
	}
	rootCmd.SetHelpTemplate(fmt.Sprintf(`
 _____   ____  __  _________
/__  /  / __ \/  |/  / ____/
  / /  / / / / /|_/ / / __  
 / /__/ /_/ / /  / / /_/ /  
/____/\____/_/  /_/\____/  Terminal Slides by Kris Nova 

{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}

%s

Author  : Kris Nova <kris@nivenly.com>
Bugs    : github.com/kris-nova/public-speaking
`, themeStr))
}

func printHelp() {
	rootCmd.Help()
	os.Exit(1)
}

// RunZOMG is the main entry point of the program
func RunZOMG(o *ZOMGOptions) error {
	if o.dirToParse == "" {
		logger.Critical("Invalid directory. Usage: zomg <presentation-directory> <flags>")
		logger.Critical("Run zomg -h / --help for more info")
		os.Exit(1)
	}
	// Process theme
	found := false
	for ii, tmap := range themeMap {
		if ii == o.themeName {
			o.theme = tmap.theme
			found = true
		}
	}
	if !found {
		// Default to Nova
		o.theme = &mdterm.ThemeNova{}
	}
	// Init the presentation
	p, err := mdterm.LoadPresentation(o.dirToParse)
	if err != nil {
		return fmt.Errorf("unable to load presentation: %v", err)
	}
	p.SetTheme(o.theme)

	// Render is what reasons about the slide data
	err = p.RenderPresentation()
	if err != nil {
		return fmt.Errorf("unable to process presentation: %v", err)
	}
	// Display will actually write the processed information to the TTY buffer
	err = p.DisplayPresentation()
	if err != nil {
		return fmt.Errorf("unable to display presentation: %v", err)
	}
	return nil
}
