package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/lapis2411/tools/cmd/deepl"
)

var rootCmd = &cobra.Command{
	Use:   "lps",
	Short: "Usefull tool set",
	Long:  `Usefull tool set`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(deepl.DeeplCmd)
}
