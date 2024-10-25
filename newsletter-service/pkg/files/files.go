package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func IsValidFileType(fileHeader *multipart.FileHeader) bool {
	mimeType := fileHeader.Header.Get("Content-Type")
	if mimeType != "application/pdf" && mimeType != "image/png" {
		return false
	}

	extension := filepath.Ext(fileHeader.Filename)
	if extension != ".pdf" && extension != ".png" {
		return false
	}

	return true
}

func SaveFile(file io.Reader, fileHeader *multipart.FileHeader, name, path string) (string, error) {
	extension := filepath.Ext(fileHeader.Filename)
	filePath := path + "/" + name + extension

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to copy file contents: %w", err)
	}

	return filePath, nil
}
