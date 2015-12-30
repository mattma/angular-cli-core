package cmd

import (
	"fmt"
	"os"
)

// Creates the folder directory
func createFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, 0777)
		fmt.Printf("[Log]: Creating directory  \"%s\"\n", folderName)
	}
}
