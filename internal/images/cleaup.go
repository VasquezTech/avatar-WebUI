package images

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Folder struct {
	F fs.File
}

var Dist fs.FS
var Static Folder

func MoveDir(sourceDir, destinationDir string) error {
	// Get the list of files and subdirectories in the source directory
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(destinationDir, 0755); err != nil {
		return err
	}
	fmt.Println(len(files))
	// Move each file and subdirectory to the destination directory
	for _, file := range files {
		sourcePath := filepath.Join(sourceDir, file.Name())
		destinationPath := filepath.Join(destinationDir, file.Name())

		// If it's a directory, call the function recursively
		if file.IsDir() {
			if err := MoveDir(sourcePath, destinationPath); err != nil {
				return err
			}
		} else {
			// Move the file
			fmt.Println("Moving files", sourcePath, destinationPath)
			if err := os.Rename(sourcePath, destinationPath); err != nil {
				//	if err := CleanupIMG(sourcePath, destinationPath); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	// Remove the source directory after moving its contents
	if err := os.Remove(sourceDir); err != nil {
		return err
	}

	return nil
}

func CleanupIMG(bodyImg, traitImg string) error {
	// Open the two images
	img1File, err := os.Open(bodyImg)
	if err != nil {
		fmt.Println("Error opening body:", err)
		return err
	}
	defer img1File.Close()

	img2File, err := os.Open(traitImg)
	if err != nil {
		fmt.Println("Error opening trait:", err)
		return err
	}
	defer img2File.Close()

	// Decode the images
	img1, _, err := image.Decode(img1File)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	img2, _, err := image.Decode(img2File)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Get the dimensions of the images
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	// Ensure the images are of the same size
	if bounds1 != bounds2 {
		fmt.Println("Error: Images are not of the same size")
		return fmt.Errorf("Error: Images are not of the same size")
	}

	// Create a new image to store the result
	result := image.NewRGBA(bounds1)

	//var results []image.RGBA
	// Compare each pixel and remove identical ones
	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			pixel1 := img1.At(x, y)
			pixel2 := img2.At(x, y)

			// If pixels are not identical, keep the pixel from the first image
			if pixel1 != pixel2 {
				result.Set(x, y, pixel2)
				//	results = append(results, *result)
			}
		}
	}
	saveResult(traitImg, result)
	return nil
}
func saveResult(traitImg string, result *image.RGBA) {

	// Save the result to a new image file
	os.Remove(traitImg)
	outFile, err := os.Create(traitImg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Comparison complete. Result saved as", traitImg)
}
