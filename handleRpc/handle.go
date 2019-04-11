import (
	"context"
)

func Handler(r *Request) {
	timeout := r.Value("timeout")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	done := make(chan struct{}, 1)
	go func(){
		RPC(ctx, ...)
		done <- struct {}
	}
	select {
	case <- done:
		//nice ...
	case <-ctx.Done():
		//timeout ...
	}
}
