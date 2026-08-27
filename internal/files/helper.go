package files

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/unim26/filehunt/internal/colors"
)

// verify if path is valid or not
func VerifyPath(path string) error {

	_, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return errors.New(fmt.Sprint(colors.RED, err.Error(), "\nProvide the correct source directory"))
	}

	return nil
}

// search file and return path if search file name is found
func SearchFile(sourceDir string, query string) (*[]string, error) {
	paths := []string{}

	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return &paths, err
	}

	for _, entry := range entries {
		if len(entry.Name()) > 0 && entry.Name()[0] == '.' {
			continue
		}

		path, err := filepath.Abs(fmt.Sprintf("%s/%s", sourceDir, entry.Name()))
		if err != nil {
			return &paths, err
		}

		// fmt.Printf("idx: %d, len: %d\n", idx, len(entries))
		// fmt.Printf("current path:%s\n", path)
		// Check if the current entry matches the query
		if strings.Compare(entry.Name(), query) == 0 {
			paths = append(paths, path)
			// fmt.Printf("%sFound the file%s: %s\n", colors.GREEN, colors.RESET, path)
			// if idx == len(entries) {
			// 	return nil
			// }
		}

		if entry.IsDir() {
			ps, err := SearchFile(path, query)
			if err == nil {
				paths = append(paths, *ps...)
			}
		}
	}

	return &paths, nil
}
