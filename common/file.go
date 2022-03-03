package common

import (
	"mime/multipart"
	"net/http"
)

func DetectContentType(file *multipart.FileHeader) (string, error) {
	reader, err := file.Open()
	if err != nil {
		return "", err
	}
	buff := make([]byte, 512)
	if _, err := reader.Read(buff); err != nil {
		return "", err
	}
	return http.DetectContentType(buff), nil
}
