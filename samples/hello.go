package samples

import (
	"context"
)

type Rpc interface {
	Refer() string
}

type Hello interface {
	Rpc
	SayHello(context.Context, string) (string, error)
	SayHelloAgain(context.Context, string) (string, error)
}

type World interface {
	Rpc
	SayWorld(context.Context, string) (string, error)
	SayWorldAgain(context.Context, string) (string, error)
}
