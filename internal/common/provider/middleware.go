package provider

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type MiddlewareFunc func(c *resty.Client, req *resty.Request) error

// GzipAndEncryptMiddleware is a middleware for encrypting data before sending to server.
func GzipAndEncryptMiddleware(c *resty.Client, req *resty.Request, enc encryption.SymmetricEncryption) error {
	if req.Header.Get("Content-Encoding") != "gzip" {
		buf := bytes.NewBuffer(nil)
		zb := gzip.NewWriter(buf)

		js, err := json.Marshal(req.Body)
		if err != nil {
			return err
		}

		_, err = zb.Write(js)
		if err != nil {
			return err
		}

		err = zb.Close()
		if err != nil {
			return err
		}

		compressed := buf.Bytes()
		req.Header.Set(
			"Content-Encoding", "gzip")

		//req.SetBody(compressed)

		// encryption
		encrypted, err := enc.Encrypt(compressed)
		if err != nil {
			logrus.Error(err)
		}
		req.SetBody(encrypted)
	}

	return nil
}
