package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "lps",
		Short: "Usefull tool set",
		Long:  `Usefull tool set`,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Execute()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		},
	}
	rootCmd.AddCommand(GetDepplCmd())
	rootCmd.AddCommand(GetDictionaryCmd())
}

func GetCommand() *cobra.Command {
	return rootCmd
}
