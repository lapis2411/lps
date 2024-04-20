package main

import (
	"github.com/lapis2411/lps/cmd"
)

func main() {
	rootCmd := cmd.GetCommand()
	rootCmd.Execute()
}
