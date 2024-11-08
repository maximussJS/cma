package services

import (
	"cma/packages/config"
	"fmt"
	"os"
	"path/filepath"
)

type Filesystem struct {
	BasePath string
}

func NewFilesystem() (*Filesystem, error) {
	basePath := config.GlobalConfig.BasePath

	err := os.MkdirAll(basePath, os.ModePerm)

	if err != nil {
		return nil, fmt.Errorf("Filesystem.NewFilesystem() mkdir all error %w", err)
	}

	return &Filesystem{BasePath: basePath}, nil
}

func (fs *Filesystem) WriteString(filename, content string) error {
	filePath := filepath.Join(fs.BasePath, filename)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("Filesystem.WriteString() write file error %w", err)
	}

	return nil
}

func (fs *Filesystem) ReadString(filename string) (string, error) {
	filePath := filepath.Join(fs.BasePath, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", nil
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", fmt.Errorf("Filesystem.ReadString() read file error %w", err)
	}

	return string(data), nil
}
