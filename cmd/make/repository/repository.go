package repository

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// Repository holds the template data
type Repository struct {
	RepositoryName string
	FileName       string
}

// Convert PascalCase to snake_case
func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(re.ReplaceAllString(str, "${1}_${2}"))
}

func createRepository(repositoryName string) {
	fileName := toSnakeCase(repositoryName)

	// Define template file and output directory
	tmplFile := "templates/repository_template.go.tmpl"
	outputDir := "src/repositories"

	// Parse the template
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Prepare template data
	repo := Repository{
		RepositoryName: repositoryName,
		FileName:       fileName,
	}

	// Ensure output directory exists
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	// Output file path
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", fileName))

	// Check if file already exists
	if _, err := os.Stat(outputFile); err == nil {
		log.Fatalf("Error: The file %s already exists.\n", outputFile)
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Execute template and write to file
	err = tmpl.Execute(file, repo)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Printf("Repository %s created successfully at %s\n", repositoryName, outputFile)

}

// RepositoryCmd generates a repository file
var RepositoryCmd = &cobra.Command{
	Use:   "repository [RepositoryName]",
	Short: "Create a new repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createRepository(args[0])
	},
}
