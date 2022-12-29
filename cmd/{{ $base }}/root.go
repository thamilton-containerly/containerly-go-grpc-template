package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "server is example cmd app build with golang",
		Long:  `server is example golang web service using gin and gorm`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

}
