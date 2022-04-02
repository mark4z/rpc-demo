package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HelloServerImpl struct {
}

func (i *HelloServerImpl) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("hello %s", name), nil
}

func (i *HelloServerImpl) SayHelloAgain(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("hello %s again", name), nil
}

func main() {
	server := &HelloServerImpl{}

	mux := http.DefaultServeMux
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in, _ := ioutil.ReadAll(r.Body)
		method := r.URL.Path
		outs, err := Invoke(server, method, in)
		if err != nil {
			return
		}
		fmt.Println(outs)
		_, _ = w.Write([]byte(outs))
	}))
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func Invoke(server *HelloServerImpl, method string, in []byte) (string, error) {
	methods := strings.Split(method, "/")

	method = methods[1]
	if method == "SayHello" {
		return server.SayHello(context.Background(), string(in))
	} else if method == "SayHelloAgain" {
		return server.SayHelloAgain(context.Background(), string(in))
	}
	return "", fmt.Errorf("method not found")
}
