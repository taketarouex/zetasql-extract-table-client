package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tktkc72/zetasql-extract-table-client/github.com/tktkc72/sqlanalyzer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	target := os.Getenv("TARGET")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(target))
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial(target, opts...)
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
