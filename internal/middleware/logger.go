package middleware

import (
	"fmt"
	"time"

	"github.com/wb-go/wbf/ginext"
)

type Logger struct {
	Method string
	URL    string
	Time   time.Time
}

func LoggerMiddleware(logCh chan<- Logger) ginext.HandlerFunc {
	return func(c *ginext.Context) {
		start := time.Now()

		c.Next()

		data := Logger{
			Method: c.Request.Method,
			URL:    c.Request.URL.String(),
			Time:   start,
		}
		select {
		case logCh <- data:
		default:
		}
	}
}

func RunLogger(logCh <-chan Logger) {
	go func() {
		for data := range logCh {
			fmt.Printf("[%s] %s %s\n", data.Time.Format(time.RFC3339), data.Method, data.URL)
		}
	}()
}
