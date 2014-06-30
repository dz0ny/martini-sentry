package shade

import (
	"github.com/go-martini/martini"
	"github.com/kisielk/raven-go/raven"
)

func Middleware(dsn string) martini.Handler {
	if dsn == "" {
		panic("Error: No DSN detected!\n")
	}
	client, _ := raven.NewClient(dsn)

	return func(context martini.Context) {
		defer func() {
			if err := recover(); err != nil {
				const size = 1 << 12
				buf := make([]byte, size)
				n := runtime.Stack(buf, false)
				client.CaptureMessage(err, string(buf[:n]))
			}
		}()
		context.Next()
	}
}
