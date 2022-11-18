package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Automate checklist to ensure you are ready to go",
	Long: `A small CLI that will run some commands for you, depending on the chosen config. 
Written with ❤️ in Go.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
