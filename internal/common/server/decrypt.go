package server

import (
	"bytes"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"io"
	"net/http"
)

// RequestDecryptMiddleware — decryption middleware for incoming http requests.
func RequestDecryptMiddleware(next http.Handler, enc *encryption.Symmetric) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		decrypted, eErr := enc.Decrypt(b)
		if eErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(decrypted))

		next.ServeHTTP(rw, r.WithContext(r.Context()))

	})
}
