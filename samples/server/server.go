package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var serviceMap = map[string]interface{}{}

func main() {
	serviceMap["Hello"] = &HelloServerImpl{}
	serviceMap["World"] = &WorldServerImpl{}

	mux := http.DefaultServeMux
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in, _ := ioutil.ReadAll(r.Body)
		method := r.URL.Path

		outs, err := Invoke(method, string(in))
		if err != nil {
			return
		}
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(outs)
		_, _ = w.Write([]byte(outs))
	}))
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func Invoke(method string, in string) (string, error) {
	methods := strings.Split(method, "/")

	service := serviceMap[methods[1]]
	methodName := methods[2]

	//find method by name
	methodValue := reflect.ValueOf(service).MethodByName(methodName)
	if !methodValue.IsValid() {
		return "", fmt.Errorf("method %s not found", methodName)
	}

	ins := make([]reflect.Value, 2)
	ins[0] = reflect.ValueOf(context.Background())
	ins[1] = reflect.ValueOf(in)

	outs := methodValue.Call(ins)
	if outs[1].Interface() != nil {
		return outs[0].Interface().(string), outs[1].Interface().(error)
	}

	return outs[0].Interface().(string), nil
}

type HelloServerImpl struct {
}

func (i *HelloServerImpl) SayHello(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("hello %s", req), nil
}

func (i *HelloServerImpl) SayHelloAgain(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("hello %s again", req), nil
}

type WorldServerImpl struct {
}

func (i *WorldServerImpl) SayWorld(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("%s World", req), nil
}

func (i *WorldServerImpl) SayWorldAgain(ctx context.Context, req string) (string, error) {
	return fmt.Sprintf("%s World again", req), nil
}
