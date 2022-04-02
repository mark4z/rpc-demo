package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

	world := &WorldClientImpl{}
	syw(world, "world")
	syw(world, "moto")

	sywAgain(world, "world")
	sywAgain(world, "moto")
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

func syw(hello *WorldClientImpl, in string) {
	sayHello, err := hello.SayWorld(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}

func sywAgain(hello *WorldClientImpl, in string) {
	sayHello, err := hello.SayWorldAgain(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}

type HelloClientImpl struct {
}

func (i *HelloClientImpl) SayHello(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/Hello/SayHello", name)
}

func (i *HelloClientImpl) SayHelloAgain(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/Hello/SayHelloAgain", name)
}

type WorldClientImpl struct {
}

func (i *WorldClientImpl) SayWorld(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/World/SayWorld", name)
}

func (i *WorldClientImpl) SayWorldAgain(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/World/SayWorldAgain", name)
}
