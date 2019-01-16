package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"test/pb/demo"
	"time"

	"google.golang.org/grpc"
)

func printMessage(clnt demo.HelloClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	m, err := clnt.SayHello(
		ctx,
		&demo.Who{
			Name: message,
		},
	)
	if err != nil {
		fmt.Printf("%v.SayHello(_) = _, %v: \n", clnt, err)
	} else {
		fmt.Println(m.GetMessage())
	}
}

func main() {
	name := flag.String("n", "李国泰", "name aa")
	port := flag.String("p", "8888", "port")
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	fmt.Println("request service")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", *port), opts...)
	if err != nil {
		log.Fatalf("fail to dial service: %v", err)
	} else {
		defer conn.Close()
		client := demo.NewHelloClient(conn)

		waitc := make(chan struct{})
		tick := 1
		go func() {
			for tick <= 3 {
				printMessage(
					client,
					fmt.Sprintf("%s, %d times", *name, tick),
				)
				tick++
				time.Sleep(time.Second * 1)
			}
			close(waitc)
		}()
		<-waitc
	}
}
