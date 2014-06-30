martini-sentry
==============

Sentry middleware for Martini. Remember to use it before all other middlewares.

### How to use
```
import "github.com/dz0ny/martini-sentry"

m := martini.Classic()
m.Use(shade.Middleware(os.GetEnv("SENTRY_DSN")))
```
