package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/xprst/whd-grpc-base/test/pb/demo"
	"log"
	"time"

	"google.golang.org/grpc"
)

func printMessage(clnt demo.HelloClient, message string) {
	// 设置请求的超时时间
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
	name := flag.String("n", "ligt", "name aa")
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
