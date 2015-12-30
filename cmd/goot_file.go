package cmd

import (
	"fmt"
	"os"
)

type gootFile struct {
	anchor, packName, fileName string
	okMessage                  string
	template                   Parseble
}

type filesList []gootFile

func (fl filesList) Process() error {
	for _, file := range fl {
		err := file.performCreation()
		if err != nil {
			return err
		}
	}
	return nil
}

// Performs creation based on the subcommand passed
func (gf gootFile) performCreation() error {
	if err := gf.createFile(); err != nil {
		return err
	}
	return nil
}

// Creates the based on the construction
// passed on the gootstrap file
func (gf gootFile) createFile() error {

	// Creates the file and defer its closing
	fileCreate, err := os.Create(gf.fileName)
	if err != nil {
		return err
	}
	defer fileCreate.Close()

	// Writes the template into file and
	// then, writes the output to os.Stdout.
	fileCreate.WriteString(gf.template.Parse())
	fmt.Println(gf.okMessage)

	return nil
}
