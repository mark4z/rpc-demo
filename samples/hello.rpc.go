package samples

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var serviceMap = map[string]interface{}{}

func RegisterService(s Rpc) {
	serviceMap[s.Refer()] = s
}

func GetService(name string) interface{} {
	return serviceMap[name]
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
	request, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080%s", method), strings.NewReader(in))
	do, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	res, _ := ioutil.ReadAll(do.Body)
	return string(res), nil
}
