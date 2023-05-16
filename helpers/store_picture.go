package helpers

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func CheckIfExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return true
	}
	return false
}

func GetFilePathAndFileName(folder string) (string, string, error) {
	id := uuid.New().String()
	fileName := fmt.Sprintf("%s.%s", id, "jpg")
	currentFolder, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	folderName := filepath.Join(currentFolder, folder)
	filePath := filepath.Join(folderName, fileName)

	return filePath, fileName, nil
}

func StorePictureInLocalFolder(file multipart.File, folder, filePath string) error {

	// Verificar se pasta existe, se não existir, criar uma nova pasta
	if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(folder, os.ModePerm); err != nil {
			return err
		}
	}

	// Criar um arquivo dentro da pasta, porém o arquivo ainda está vázio
	localPicture, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer localPicture.Close()

	// Cria uma cópia do arquivo recebido pela requisição (file) e
	// cola essa cópia dentro do novo arquivo criado na função anterior (localPicture)
	if _, err = io.Copy(localPicture, file); err != nil {
		return err
	}

	return nil
}
