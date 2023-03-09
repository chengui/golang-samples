package gyn

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Account struct {
	User string
	Pass string
}

func BasicAuth(auths []string) HandlerFunc {
	pairs := make(map[string]*Account)
	for _, auth := range auths {
		if !strings.Contains(auth, ":") {
			continue
		}
		code := base64.StdEncoding.EncodeToString([]byte(auth))
		pair := strings.SplitN(auth, ":", 2)
		pairs[code] = &Account{User: pair[0], Pass: pair[1]}
	}

	return func(c *Context) {
		auth := c.Req.Header.Get("Authorization")
		code := strings.TrimSpace(strings.TrimPrefix(auth, "Basic"))
		if val, ok := pairs[code]; !ok {
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
			c.Fail(401, fmt.Errorf("Unauthorized"))
		} else {
			c.Set("auth", val)
		}
	}
}
