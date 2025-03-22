// Define the package
package main

import (
	"fmt"
	"os"
)

type Permissions int

const READ_WRITE_OWNER_ONLY Permissions = 0755

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file path provided")
		return
	}

	// Accept a file path to a folder as an argument  For now, it has to be an exact path
	var filepath string = os.Args[1]

	fmt.Println("Changing directory to", filepath)
	err := os.Chdir(filepath)
	// Return an error if the error does not come from a filepath known to exist.
	if err != nil {
		fmt.Println("Input is not a valid filepath")
		return
	}

	dataTypesToFolders := map[string]string{
		"png": "images",
		"pdf": "pdfs",
		"jpg": "images",
	}

	// create the target folders if these don't exist
	for _, folder := range dataTypesToFolders {
		fmt.Println("Checking if target folders exist")
		_, err := os.Stat(folder)
		if os.IsNotExist(err) {
			fmt.Println("Folder does not exists, creating...", folder)
			os.Mkdir(folder, 0755) // Only owner has read, write permissions
		}
	}

	fmt.Println("All target folders exist.")

	// Iterate each file in the current working directory. Move each to the corresponding folder based on its type.

	// Get the current files

	// Get the file type extension

	// Move to target folder

	// Requirements
	// navigate to the file path
	// Inside the folder, based on file type, move to a specific set of folders. e.g .png files go to the images folder
	// Print that we are done
}
