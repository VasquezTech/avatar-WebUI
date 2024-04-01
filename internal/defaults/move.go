package defaults

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

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
			if err := os.Rename(sourcePath, destinationPath); err != nil {
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
