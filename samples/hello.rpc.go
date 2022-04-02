package samples

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// generate code

var serviceMap = map[string]interface{}{}

func RegisterService(s Rpc) {
	//service export
	serviceMap[s.Refer()] = s
}

func GetService(name string) interface{} {
	return serviceMap[name]
}

func DefaultHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		in, _ := ioutil.ReadAll(r.Body)
		method := r.URL.Path

		//protocol
		outs, err := Handler(method, string(in))
		if err != nil {
			return
		}
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(outs)
		_, _ = w.Write([]byte(outs))
	}
}

func Handler(method string, in string) (string, error) {
	methods := strings.Split(method, "/")

	service := GetService(methods[1])
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

type HelloClientImpl struct {
}

func (i *HelloClientImpl) SayHello(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/Hello/SayHello", name)
}

func (i *HelloClientImpl) SayHelloAgain(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/Hello/SayHelloAgain", name)
}

func (i *HelloClientImpl) Refer() string {
	return "Hello"
}

type WorldClientImpl struct {
}

func (i *WorldClientImpl) SayWorld(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/World/SayWorld", name)
}

func (i *WorldClientImpl) SayWorldAgain(ctx context.Context, name string) (string, error) {
	return Invoke(ctx, "/World/SayWorldAgain", name)
}

func (i *WorldClientImpl) Refer() string {
	return "World"
}

func Invoke(ctx context.Context, method string, in string) (string, error) {
	client := http.DefaultClient
	//router & service naming
	request, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080%s", method), strings.NewReader(in))
	do, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	res, _ := ioutil.ReadAll(do.Body)
	return string(res), nil
}
