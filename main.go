package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vvm",
		Short: "vvm is a version manager CLI tool",
	}

	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install a version (placeholder)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Install command placeholder")
		},
	}

	rootCmd.AddCommand(installCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
