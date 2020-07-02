package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/tautcony/qart/models"
)

func Read(path string) ([]byte, *models.FileInfo, error) {
	p, err := filepath.Abs(path)
	log.Printf("Read <- %s\n", p)
	if err != nil {
		panic(err)
	}
	dir, err := os.Stat(p)
	if err != nil {
		return nil, nil, err
	}
	fi := &models.FileInfo{
		Name:    dir.Name(),
		ModTime: dir.ModTime(),
		Size:    dir.Size(),
		IsDir:   dir.IsDir(),
	}
	data, err := ioutil.ReadFile(p)
	return data, fi, err
}

func Write(path string, data []byte) error {
	p, err := filepath.Abs(path)
	log.Printf("Write ->: %s\n", p)
	if err != nil {
		panic(err)
	}

	return ioutil.WriteFile(p, data, 0666)
}

func Remove(path string) error {
	p, err := filepath.Abs(path)
	log.Printf("Remove x %s\n", p)
	if err != nil {
		panic(err)
	}
	return os.Remove(p)
}
