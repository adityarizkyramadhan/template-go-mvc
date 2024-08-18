package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveFile saves the file from *multipart.FileHeader to the desired directory.
// It returns the file path of the saved file and an error if there is any.
// The saved file path is unique by using a timestamp.
func SaveFile(file *multipart.FileHeader, path string) (string, error) {
	if file == nil {
		return "", nil
	}

	// Ensure the base directory exists
	basePath := "storage"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return "", err
	}

	// Create a unique file name using timestamp
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	fileName = strings.ReplaceAll(fileName, " ", "_")
	file.Filename = fileName
	filePath := filepath.Join(basePath, path, fileName)

	// Ensure the directory exists
	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", err
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file content from src to dst
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	baseLink := fmt.Sprintf("%s/api/v1/storage/%s?filename=%s", os.Getenv("BASE_URL"), path, file.Filename)

	return baseLink, nil
}
