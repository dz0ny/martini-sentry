package shade

import (
  "github.com/kisielk/raven-go/raven"
	"github.com/go-martini/martini"
)

func Middleware(dsn string) martini.Handler {
	if dsn == "" {
		panic("Error: No DSN detected!\n")
	}
	client, _ := raven.NewClient(dsn)
	
	return func(context martini.Context) {
		defer func() {
			if err := recover(); err != nil {
				client.CaptureMessage(err.Error())
			}
		}()
		context.Next()
	}
}
