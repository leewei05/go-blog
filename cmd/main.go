package main

import (
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
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
