// Define the package
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("file path provided")
		os.Exit(1)
	}

	// Accept a file path to a folder as an argument  For now, it has to be an exact path
	var directoryPath string = os.Args[1]

	log.Println("Changing directory to", directoryPath)
	err := os.Chdir(directoryPath)
	// Return an error if the error does not come from a filepath known to exist.
	if err != nil {
		log.Fatal("Input is not a valid filepath")
		os.Exit(1)
	}

	dataTypesToFolders := map[string]string{
		".png":  "images",
		".pdf":  "pdfs",
		".jpg":  "images",
		".txt":  "texts",
		".jpeg": "images",
		".zip":  "zip_files",
		".dmg":  "dmgs_images",
	}

	// create the target folders if these don't exist
	for _, folder := range dataTypesToFolders {
		log.Println("Checking if target folders exist")
		_, err := os.Stat(folder)
		if os.IsNotExist(err) {
			log.Println("Folder does not exists, creating...", folder)
			os.Mkdir(folder, 0755) // Only owner has read, write permissions
		}
	}

	fmt.Println("All target folders exist.")

	// Iterate each file in the current working directory. Move each to the corresponding folder based on its type.

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			extension := filepath.Ext(filename)

			destinationFolder, exists := dataTypesToFolders[extension]

			if !exists {
				log.Println("File type not supported", extension)
				continue
			}
			cwd := filepath.Join(currentWorkingDirectory, filename)
			destination := filepath.Join(currentWorkingDirectory, destinationFolder, filename)
			move(cwd, destination)

		}
	}

	// Requirements
	// navigate to the file path
	// Inside the folder, based on file type, move to a specific set of folders. e.g .png files go to the images folder
	// Print that we are done
}

func move(sourcePath string, targetPath string) {
	error := os.Rename(sourcePath, targetPath)

	if error != nil {
		log.Println("Error moving file", error)
		os.Exit(1)
	}

	log.Println("File moved", sourcePath)
}
