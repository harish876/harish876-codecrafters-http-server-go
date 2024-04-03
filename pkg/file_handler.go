package pkg

import (
	"log"
	"os"
	"path/filepath"
)

func HandleFile(fileName, directory string) (string, error) {
	absolutePath := filepath.Join(directory, fileName)
	log.Println("Path is", absolutePath)
	contents, err := os.ReadFile(absolutePath)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
