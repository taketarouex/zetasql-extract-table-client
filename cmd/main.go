package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tktkc72/zetasql-extract-table-client/github.com/tktkc72/sqlanalyzer"
	"google.golang.org/grpc"
)

func main() {
	target := os.Getenv("TARGET")
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer conn.Close()
	client := sqlanalyzer.NewExtractTableNamesClient(conn)
	flag.Parse()
	state := flag.Arg(0)
	request := sqlanalyzer.ExtractTableNamesRequest{
		Statement: state,
	}
	res, err := client.Do(context.Background(), &request)
	if err != nil {
		log.Fatalf("err: %v, response: %v", err, res)
	}
	for _, t := range res.TableNames {
		fmt.Printf("table: %s\n", t)
	}
}
