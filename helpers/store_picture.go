package helpers

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func CheckIfExists(filePath string) bool {
	err := filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path, filePath)
		if path == filePath {
			return nil
		}
		return errors.New("NotExist")
	})
	if err != nil {
		return false
	}
	return true
}

func DeleteFile(filepath string) error {
	if err := os.RemoveAll(filepath); err != nil {
		return err
	}
	return nil
}

func GetFilePathAndFileName(folder string) (string, string, string, error) {
	id := uuid.New().String()
	fileName := fmt.Sprintf("%s.%s", id, "jpg")
	currentFolder, err := os.Getwd()
	if err != nil {
		return "", "", "", err
	}

	folderName := filepath.Join(currentFolder, folder)
	filePath := filepath.Join(folderName, fileName)
	relativePath := strings.Split(filePath, "cmd\\")[1]

	return filePath, fileName, relativePath, nil
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
