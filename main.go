package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/unim26/filehunt/colors"
)

// search file
func searchFile(source string, path string, query string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", nil
	}

	for idx, entry := range entries {
		fp, _ := filepath.Abs(fmt.Sprintf("%s/%s", source, entry.Name()))

		//check if current path the file
		if !entry.IsDir() && strings.Compare(entry.Name(), query) == 0 {
			return fmt.Sprintf("%sFound the file%s: %s", colors.GREEN, colors.RESET, fp), nil
		} else if idx+1 == len(entries) {
			return fmt.Sprintf("%sno file found with name containing %s in %s", colors.RED, query, source), nil
		}

		//check if current entry is directory
		if entry.IsDir() {
			searchFile(source, fp, query)
		}
	}
	return "", nil
}

// verify path and search file
func verifyPathAndSearchFile(source string, query string) error {

	//check if source directory path is valid or not
	_, err := os.Stat(source)
	if errors.Is(err, fs.ErrNotExist) {
		return errors.New(fmt.Sprint(colors.RED, err.Error(), "\nProvide the correct source directory"))
	}

	absPath, _ := filepath.Abs(source)

	path, err := searchFile(absPath, absPath, query)
	if err != nil {
		return nil
	}

	fmt.Println(path)

	return nil
}

func main() {
	//get the arguments
	args := os.Args

	//check if there are 3 argument `[filehunt starting_dir search_term]`
	if len(args) != 3 {
		fmt.Printf("%sNot a valid command%s, Correct usages: ", colors.RED, colors.RESET)
		fmt.Println("\n\tfilehunt <source directory path> <search term>")
		return
	}

	source_dir := args[1]
	search_term := args[2]

	err := verifyPathAndSearchFile(source_dir, search_term)
	if err != nil {
		fmt.Println(err)
	}

}
