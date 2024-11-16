package utils

import (
	"os"
	"path/filepath"
	"sort"
)

// ReadSQLFiles reads all .sql files in the specified directory, sorts them by filename.
func ReadSQLFiles(migrationsDir string) ([]string, error) {
	var files []string

	// Read files from the directory
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Add only .sql files to the list
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort the files by filename to ensure correct execution order
	sort.Strings(files)

	return files, nil
}
