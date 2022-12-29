package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	rpc "{{ $base }}/rpc"
)

type Server struct {
	rpc.UnimplementedContainerlyServer
}

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start",
		Long:  `start`,
		Run: func(cmd *cobra.Command, args []string) {
			s := grpc.NewServer()
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", "8080"))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			rpc.RegisterContainerlyServer(s, &Server{})

			log.Printf("server listening at %v", lis.Addr())
			err = s.Serve(lis)
			if err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
