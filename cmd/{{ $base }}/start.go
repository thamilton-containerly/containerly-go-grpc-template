package main

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type Service struct {
	Server *grpc.Server
}

type Server struct {
	containerly.UnimplementedContainerlyServer
}

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
