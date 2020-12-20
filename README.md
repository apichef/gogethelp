# go.get.help

### Env(key, default)

The Env function retrieves the value of an environment variable or returns a default.

```go
package main

import h "github.com/apichef/gogethelp"

func main() {
	env := h.Env("APP_ENV", "local")
}
```

