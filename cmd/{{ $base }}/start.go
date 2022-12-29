package main

import (
	containerly "containerly-go-grpc-template/rpc"
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
			//s := grpc.NewServer()
			//lis, err := net.Listen("tcp", fmt.Sprintf(":%v", "50051"))
			//if err != nil {
			//	log.Fatalf("failed to listen: %v", err)
			//}
			//containerly.RegisterContainerlyServer(s, &Server{})
			//log.Printf("server listening at %v", lis.Addr())
			//err = s.Serve(lis)
			//if err != nil {
			//	log.Fatalf("failed to serve: %v", err)
			//}
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
