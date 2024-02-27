package middleware

import (
	"log"
	"time"

	"github.com/cpf2021-gif/gow"
)

func Logger() gow.HandlerFunc {
	return func(c *gow.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
