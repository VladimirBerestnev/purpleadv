package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	statusCode int
}

func (writer *WrapperWriter) WriteHeader(code int) {
	writer.statusCode = code
	writer.ResponseWriter.WriteHeader(code)
}
