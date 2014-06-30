package shade

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/kisielk/raven-go/raven"
	"net/http"
	"runtime"
)

func Middleware(dsn string) martini.Handler {
	if dsn == "" {
		panic("Error: No DSN detected!\n")
	}
	client, _ := raven.NewClient(dsn)

	return func(res http.ResponseWriter, context martini.Context) {
		defer func() {
			if err := recover(); err != nil {
				const size = 1 << 12
				buf := make([]byte, size)
				n := runtime.Stack(buf, false)
				client.CaptureMessage(fmt.Sprintf("%v\nStacktrace:\n%s", err, buf[:n]))
				res.WriteHeader(http.StatusInternalServerError)
			}
		}()
		context.Next()
	}
}
