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
func SearchFile(sourceDir string, query string) (string, error) {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if len(entry.Name()) > 0 && entry.Name()[0] == '.' {
			continue
		}

		path, err := filepath.Abs(fmt.Sprintf("%s/%s", sourceDir, entry.Name()))
		if err != nil {
			return "", err
		}

		// Check if the current entry matches the query
		if strings.Compare(entry.Name(), query) == 0 {
			return path, nil
		}

		if entry.IsDir() {
			foundPath, err := SearchFile(path, query)
			if err == nil && foundPath != "" {
				return foundPath, nil
			}
		}
	}

	return "", fmt.Errorf("%sno file found with name containing %s in %s", colors.RED, query, sourceDir)
}
