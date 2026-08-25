package main

import (
	"fmt"
	"os"

	"github.com/unim26/filehunt/internal/colors"
	"github.com/unim26/filehunt/internal/files"
)

func main() {
	//get the arguments
	args := os.Args

	//check if there are 3 argument `[filehunt starting_dir searchTerm]`
	if len(args) != 3 {
		fmt.Printf("%sNot a valid command%s, Correct usages: ", colors.RED, colors.RESET)
		fmt.Println("\n\tfilehunt <source directory path> <search term>")
		return
	}

	sourceDir := args[1]
	searchTerm := args[2]

	err := files.VerifyPath(sourceDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	path, err := files.SearchFile(sourceDir, searchTerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	if path != "" {
		fmt.Printf("%sFound the file%s: %s\n", colors.GREEN, colors.RESET, path)
	}

}
