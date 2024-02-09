package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/nfnt/resize"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: \n Resize Image : resize <image path> <width> <height> \n Check Size : resize check <image path>")
		return
	}

	if os.Args[1] == "check" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: resize check <image path>")
			return
		}

		imagePath := filepath.Join(wd, os.Args[2])
		file, err := os.Open(imagePath)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		defer file.Close()

		imageSize, _, err := image.DecodeConfig(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
		}

		fmt.Println("size:", imageSize.Width, "x", imageSize.Height)
		return
	}
	if len(os.Args) < 4 {
		fmt.Println("Usage: resize <image path> <width> <height>")
		return
	}

	if len(os.Args) > 4 {
		fmt.Println("Too many arguments")
		fmt.Println("Usage: resize <image path> <width> <height>")
		return
	}

	imagePath := filepath.Join(wd, os.Args[1])
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Error: %s", "File not found")
		return
	}
	defer file.Close()

	widthStr := os.Args[2]
	heightStr := os.Args[3]
	imageFilename := filepath.Base(imagePath)

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		fmt.Printf("Error: %s", "Invalid width")
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		fmt.Printf("Error: %s", "Invalid height")
		return
	}

	startTime := time.Now()

	fmt.Println("Processing image...")

	var img image.Image
	img, _, err = image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %s", "Invalid image format, only jpg, jpeg and png are supported")
		return
	}

	resizer := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	outputImageFormat := "png"
	if filepath.Ext(imagePath) == ".jpg" || filepath.Ext(imagePath) == ".jpeg" {
		outputImageFormat = "jpg"
	}

	resizedImageFilename := fmt.Sprintf("%s_%s_%s.%s", imageFilename[:len(imageFilename)-len(filepath.Ext(imageFilename))], widthStr, heightStr, outputImageFormat)

	output, err := os.Create(resizedImageFilename)
	if err != nil {
		fmt.Printf("Error: %s", "Error creating file")
		return
	}
	defer output.Close()

	switch outputImageFormat {
	case "jpg":
		err = jpeg.Encode(output, resizer, nil)
	case "png":
		err = png.Encode(output, resizer)
	}
	if err != nil {
		fmt.Printf("Error: %s", "Error encoding image")
		return
	}

	endTime := time.Now()

	duration := endTime.Sub(startTime)

	fmt.Printf("Time spent: %s\n", duration)
	fmt.Printf("Image saved successfully: \x1b]8;;file://%s\x07%s\x1b]8;;\x07\n", resizedImageFilename, resizedImageFilename)
}
