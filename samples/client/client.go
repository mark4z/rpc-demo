package main

import (
	"context"
	"github.com/mark4z/rpc-demo/samples"
)

func main() {
	hello := &samples.HelloClientImpl{}
	sya(hello, "world")
	sya(hello, "moto")

	syaAgain(hello, "world")
	syaAgain(hello, "moto")

	world := &samples.WorldClientImpl{}
	syw(world, "world")
	syw(world, "moto")

	sywAgain(world, "world")
	sywAgain(world, "moto")
}

func sya(hello *samples.HelloClientImpl, in string) {
	sayHello, err := hello.SayHello(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}

func syaAgain(hello *samples.HelloClientImpl, in string) {
	sayHello, err := hello.SayHelloAgain(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}

func syw(hello *samples.WorldClientImpl, in string) {
	sayHello, err := hello.SayWorld(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}

func sywAgain(hello *samples.WorldClientImpl, in string) {
	sayHello, err := hello.SayWorldAgain(context.Background(), in)
	if err != nil {
		panic(err)
	}
	println(sayHello)
}
