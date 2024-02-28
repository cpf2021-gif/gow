package gow

import (
	"time"

	"github.com/cpf2021-gif/gow/log"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.GetDefaultLog().Infof("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
