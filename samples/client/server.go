package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HelloClientImpl struct {
}

func (i *HelloClientImpl) SayHello(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "SayHello", name)
}

func (i *HelloClientImpl) SayHelloAgain(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "SayHelloAgain", name)
}

func Invoke(ctx context.Context, method string, in string) (string, error) {
	client := http.DefaultClient
	request, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/%s", method), strings.NewReader(in))
	do, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	res, _ := ioutil.ReadAll(do.Body)
	return string(res), nil
}

func main() {
	hello := &HelloClientImpl{}
	sya(hello, "world")
	sya(hello, "moto")

	syaAgain(hello, "world")
	syaAgain(hello, "moto")
}

func sya(hello *HelloClientImpl, in string) {
	sayHello, err := hello.SayHello(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}

func syaAgain(hello *HelloClientImpl, in string) {
	sayHello, err := hello.SayHelloAgain(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}
