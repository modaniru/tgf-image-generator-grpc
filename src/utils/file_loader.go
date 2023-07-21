package utils

import (
	"image"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

// Загружает все .png файлы в папке
func LoadPngImages(folderPath string) (map[string]image.Image, error) {
	files, err := filepath.Glob(folderPath + "/*.png")
	if err != nil {
		return nil, err
	}
	res := make(map[string]image.Image)
	for _, file := range files {
		fileName := strings.Split(file, "/")[1]
		img, err := gg.LoadImage(file)
		if err != nil {
			return nil, err
		}
		res[fileName[0:len(fileName)-4]] = img
	}
	return res, nil
}
