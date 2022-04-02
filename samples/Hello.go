package samples

import "context"

type Hello interface {
	SayHello(context.Context, string) (string, error)
	SayHelloAgain(context.Context, string) (string, error)
}

type World interface {
	SayWorld(context.Context, string) (string, error)
	SayWorldAgain(context.Context, string) (string, error)
}
