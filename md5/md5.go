package md5

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

//go:generate counterfeiter . Summer

type Summer interface {
	Sum() (string, error)
}

type fileContentsSummer struct {
	filepath string
}

func NewFileContentsSummer(filepath string) Summer {
	return &fileContentsSummer{
		filepath: filepath,
	}
}

func (f fileContentsSummer) Sum() (string, error) {
	fileToSum, err := os.Open(f.filepath)
	if err != nil {
		return "", err
	}
	defer fileToSum.Close()

	fileReader := bufio.NewReader(fileToSum)
	hash := md5.New()
	_, err = io.Copy(hash, fileReader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
