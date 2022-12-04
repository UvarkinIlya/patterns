package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Converter struct{}

func (converter Converter) Convert(sourceFile *os.File, path, format string) (err error) {
	fileExtension, err := FileExtension(sourceFile)
	if err != nil {
		return
	}

	var img image.Image
	switch fileExtension {
	case "png":
		img, err = png.Decode(sourceFile)
	case "jpg":
		img, err = jpeg.Decode(sourceFile)
	default:
		return fmt.Errorf("Format not supported. ")
	}

	if !converter.isFormatSupport(format) {
		return fmt.Errorf("Format not supported. ")
	}
	file, err := os.Create(fmt.Sprintf("%s.%s", path, format))
	if err != nil {
		return
	}

	switch format {
	case "png":
		err = png.Encode(file, img)
	case "jpg":
		err = jpeg.Encode(file, img, nil)
	default:
		return fmt.Errorf("Format not supported. ")
	}

	return
}

func (converter Converter) isFormatSupport(format string) bool {
	return format == "png" || format == "jpg"
}

func FileExtension(file *os.File) (extension string, err error) {
	fileName := file.Name()
	arr := strings.Split(fileName, ".")
	if len(arr) == 0 {
		return "", fmt.Errorf("Extension not defined. ")
	}

	return arr[len(arr)-1], nil
}

func main() {
	converter := Converter{}

	file, err := os.Open("facade/corgi.jpg")
	if err != nil {
		panic(err)
	}

	err = converter.Convert(file, "facade/newFile", "png")
	if err != nil {
		panic(err)
	}
}
