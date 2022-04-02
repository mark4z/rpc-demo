package main

import (
	"context"
	"fmt"
	"github.com/mark4z/rpc-demo/samples"
	"log"
	"net/http"
)

func main() {
	samples.RegisterService(&HelloServerImpl{})
	samples.RegisterService(&WorldServerImpl{})

	mux := http.DefaultServeMux
	mux.HandleFunc("/", samples.DefaultHandler())

	log.Fatalln(http.ListenAndServe(":8080", mux))
}

type HelloServerImpl struct {
}

func (i *HelloServerImpl) SayHello(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("hello %s", req), nil
}

func (i *HelloServerImpl) SayHelloAgain(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("hello %s again", req), nil
}

func (i *HelloServerImpl) Refer() string {
	return "Hello"
}

type WorldServerImpl struct {
}

func (i *WorldServerImpl) SayWorld(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("%s World", req), nil
}

func (i *WorldServerImpl) SayWorldAgain(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("%s World again", req), nil
}

func (i *WorldServerImpl) Refer() string {
	return "World"
}
