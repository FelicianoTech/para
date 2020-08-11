package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "para",
	//Short:   "Fetch release metrics for your project",
	Short: "Get information on packages and releases",
	Long: `para is short for Packages and Releases Analytics. It allows you to view download and install metrics for your software distributables on GitHub, Snap, Brew, and more coming soon.

Starting a new project? You can check for name availability on these platforms.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
