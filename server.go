package main

import (
	"net"
	"log"
	"google.golang.org/grpc"

	pb "github.com/rayyildiz/dmt/pb/compnay"
	"golang.org/x/net/context"
	"google.golang.org/grpc/reflection"
)

type server struct{}

const (
	port = ":50051"
)

func (s *server) List(ctx context.Context, in *pb.CompanyListRrequest) (*pb.CompanyListResponse, error) {

	//log.Printf("Request comes for List %v", *in)

	return &pb.CompanyListResponse{
		Companies: &pb.CompanyData{
			Id:          1,
			Name:        "Name",
			Descrpition: "Desc",
			LogoUrl:     "",
		},
	}, nil
}

func (s *server) GetDetail(ctx context.Context, in *pb.CompanyDetailRequest) (*pb.CompanyDetailResponse, error) {
	//log.Printf("Request comes for GetDetail %v", *in)

	return &pb.CompanyDetailResponse{
		Company: &pb.CompanyData{
			Id:          1,
			Name:        "Name Details",
			Descrpition: "Desc",
			LogoUrl:     "",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCompanyServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
