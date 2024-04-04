package disel

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func HandleGetFile(fileName, directory string) (string, error) {
	absolutePath := filepath.Join(directory, fileName)
	log.Println("Read Path is", absolutePath)
	contents, err := os.ReadFile(absolutePath)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func HandlePostFile(fileName, directory, fileContents string) error {
	absolutePath := filepath.Join(directory, fileName)
	log.Println("Write Path is", absolutePath)
	err := os.WriteFile(absolutePath, []byte(fileContents), fs.FileMode(os.O_WRONLY))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Writtent to File Successfully at %s", absolutePath)
	return nil
}
