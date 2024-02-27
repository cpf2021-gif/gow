package middleware

import (
	"time"

	"github.com/cpf2021-gif/gow"
	"github.com/cpf2021-gif/gow/log"
)

func Logger() gow.HandlerFunc {
	return func(c *gow.Context) {
		t := time.Now()
		c.Next()
		log.GetDefaultLog().Infof("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
