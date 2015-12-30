package cmd

import (
	"fmt"
	"path/filepath"
)

// Command logic
func New(appName string) {
	fmt.Printf("[Log]: Installing app  \"%s\"\n", appName)

	// Create a new app with default structure
	err := createNewApp(appName)

	if err != nil {
		fmt.Printf("[Error]: %s\n", err)
		return
	}

	fmt.Printf("[OK]: App created! Try `cd %s`.\n", appName)
}

// Creates the package with files in it
func createNewApp(appName string) error {
	sep := string(filepath.Separator)

	// Creates the project's folder
	createFolder(appName)

	// Init files
	// currentYear, currentMonth, currentDay := time.Now().Date()
	// currentDate := fmt.Sprintf("%d-%d-%d", currentYear, currentMonth, currentDay)
	// user, _ := user.Current()

	files := filesList{
		{
			anchor:    "gitignore",
			fileName:  fmt.Sprintf("%s%s.gitignore", appName, sep),
			template:  GitIgnoreFile{},
			okMessage: "[Log]: Creating \".gitignore\" file",
		},
		{
			anchor:    "main",
			fileName:  fmt.Sprintf("%s%s%s.go", appName, sep, appName),
			template:  MainFile{appName},
			okMessage: fmt.Sprintf("[Log]: Creating \"%s.go\" file", appName),
		},
	}

	err := files.Process()
	if err != nil {
		return err
	}

	return nil
}
