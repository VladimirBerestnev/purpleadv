package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := logrus.WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.Path,
			"query":       r.URL.RawQuery,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
			"referer":     r.Referer(),
		})

		writer := &WrapperWriter{w, http.StatusOK}

		start := time.Now()

		log.Info("HTTP request started")
		next.ServeHTTP(writer, r)
		duration := time.Since(start)

		log.WithFields(logrus.Fields{
			"status_code": writer.statusCode,
			"duration_ms": duration.Milliseconds(),
		}).Info("HTTP request completed")
	})
}
