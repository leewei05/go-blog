package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)

	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "goldfolye",
		Short: "Goldfolye is a SQL database",
		Run:   func(cmd *cobra.Command, args []string) { fmt.Println("Hello Goldfoyle") },
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
