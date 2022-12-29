package main

import (
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start",
		Long:  `start`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
