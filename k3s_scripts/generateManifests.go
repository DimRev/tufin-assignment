package k3sscripts

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Values struct {
	Mysql struct {
		Image        string `yaml:"image"`
		RootPassword string `yaml:"rootPassword"`
		Database     string `yaml:"database"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
	} `yaml:"mysql"`

	Wordpress struct {
		Image      string `yaml:"image"`
		DbHost     string `yaml:"dbHost"`
		DbPort     string `yaml:"dbPort"`
		DbUser     string `yaml:"dbUser"`
		DbPassword string `yaml:"dbPassword"`
		DbName     string `yaml:"dbName"`
	} `yaml:"wordpress"`
}

func (ctx *Context) GenerateManifests() error {
	values, err := ctx.loadDefaultValues()
	if err != nil {
		return err
	}

	err = ctx.renderAndSaveManifests(values)
	if err != nil {
		return err
	}

	fmt.Println("Manifests successfully generated and saved to the OS temp directory.")
	return nil
}

func (ctx *Context) loadDefaultValues() (Values, error) {
	data, err := ctx.manifests.ReadFile("manifests/default-values.yaml")
	if err != nil {
		return Values{}, fmt.Errorf("failed to read default values: %w", err)
	}

	var v Values
	if err := yaml.Unmarshal(data, &v); err != nil {
		return Values{}, fmt.Errorf("failed to parse default values: %w", err)
	}

	return v, nil
}

func (ctx *Context) renderAndSaveManifests(values Values) error {
	tempDir := filepath.Join(os.TempDir(), "tufin-assignment")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create tmp directory: %w", err)
	}

	manifestFiles := []string{
		"manifests/mysql-deployment.yaml",
		"manifests/wordpress-deployment.yaml",
	}

	for _, filePath := range manifestFiles {
		templateData, err := ctx.manifests.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read manifest file %s: %w", filePath, err)
		}

		renderedData, err := renderTemplate(string(templateData), values)
		if err != nil {
			return fmt.Errorf("failed to render template for %s: %w", filePath, err)
		}

		outputFilePath := filepath.Join(tempDir, filepath.Base(filePath))
		if err := os.WriteFile(outputFilePath, []byte(renderedData), 0644); err != nil {
			return fmt.Errorf("failed to write rendered manifest to %s: %w", outputFilePath, err)
		}

		ctx.tempFiles = append(ctx.tempFiles, outputFilePath)
		fmt.Printf("Rendered manifest saved: %s\n", outputFilePath)
	}

	return nil
}

func renderTemplate(templateData string, values Values) (string, error) {
	tmpl, err := template.New("manifest").Parse(templateData)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, values); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return rendered.String(), nil
}

func (ctx *Context) CleanupTempFiles() {
	for _, filePath := range ctx.tempFiles {
		err := os.Remove(filePath)
		if err != nil {
			fmt.Printf("Failed to remove temp file %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Removed temp file: %s\n", filePath)
		}
	}
	ctx.tempFiles = []string{}
}
