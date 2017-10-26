package main

import (
	"log"

	"google.golang.org/grpc"

	"flag"
	"fmt"
	"time"

	pb "github.com/rayyildiz/dmt-go/pb/company"
	"golang.org/x/net/context"
)

func mustDial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		panic(err)
	}
	return conn
}

func main() {
	count := 100000

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	searchAddr := flag.String("serveradd", ":50051", "Search service addr")

	flag.Parse()

	client := pb.NewCompanyClient(mustDial(*searchAddr))

	start := time.Now()

	for i := 0; i < count; i++ {
		go func() {
			_, err := client.List(ctx, &pb.CompanyListRrequest{})

			if err != nil {
				fmt.Printf("Error occured while invoking rpc %v\n", err)
			}
		}()
	}

	// fmt.Printf("Result %v\n", r)

	elapsed := time.Since(start)
	log.Printf("Binomial took  %s for %d times", elapsed, count)

}
