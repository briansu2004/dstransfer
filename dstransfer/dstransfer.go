package main

import (
	"github.com/adrianwit/dstransfer"
	"flag"
	"fmt"
	_ "github.com/adrianwit/fbc"
	_ "github.com/adrianwit/fsc"
	_ "github.com/adrianwit/mgc"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/viant/asc"
	_ "github.com/viant/bgc"
	"github.com/google/gops/agent"
	"log"
	"os"
)

var port = flag.Int("port", 8080, "service port")

func main() {

	go func() {
		if err := agent.Listen(agent.Options{}); err != nil {
			log.Fatal(err)
		}
	}()
	flag.Parse()
	service := dstransfer.New(false, nil)
	server := dstransfer.NewServer(service, *port)
	go server.StopOnSiginals(os.Interrupt)
	fmt.Printf("dstransfer listening on :%d\n", *port)
	server.ListenAndServe()
}
