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

func StorePictureInLocalFolder(file multipart.File, hfile *multipart.FileHeader, folder string) (string, error) {
	id := uuid.New().String()
	fileName := fmt.Sprintf("%s.%s", id, filepath.Ext(hfile.Filename))
	currentFolder, err := os.Getwd()
	if err != nil {
		return "", err
	}

	folderName := filepath.Join(currentFolder, folder)

	// Verificar se pasta existe, se não existir, criar uma nova pasta
	if _, err := os.Stat(folderName); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(folderName, os.ModePerm); err != nil {
			return "", err
		}
	}

	picturePath := filepath.Join(folderName, fileName)
	// Criar um arquivo dentro da pasta, porém o arquivo ainda está vázio
	localPicture, err := os.Create(picturePath)
	if err != nil {
		return "", err
	}

	defer localPicture.Close()

	// Cria uma cópia do arquivo recebido pela requisição (file) e
	// cola essa cópia dentro do novo arquivo criado na função anterior (localPicture)
	if _, err = io.Copy(localPicture, file); err != nil {
		return "", err
	}

	return picturePath, nil
}
