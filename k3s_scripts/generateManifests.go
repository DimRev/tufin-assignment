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

func GenerateManifests() error {
	values, err := loadDefaultValues()
	if err != nil {
		return err
	}

	err = renderAndSaveManifests(values)
	if err != nil {
		return err
	}

	fmt.Println("Manifests successfully generated and saved to ./tmp.")
	return nil
}

func loadDefaultValues() (Values, error) {
	data, err := os.ReadFile("manifests/default-values.yaml")
	if err != nil {
		return Values{}, YAMLFilesNotFound{FileName: "default-values.yaml", Message: err.Error()}
	}

	var v Values
	if err := yaml.Unmarshal(data, &v); err != nil {
		return Values{}, YAMLUnmarshalError{Message: err.Error()}
	}

	return v, nil
}

func renderAndSaveManifests(values Values) error {
	if err := os.MkdirAll("./tmp", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create tmp directory: %w", err)
	}

	manifestFiles := []string{
		"manifests/mysql-deployment.yaml",
		"manifests/wordpress-deployment.yaml",
	}

	for _, filePath := range manifestFiles {
		fmt.Printf("Rendering %s manifest...\n", filePath)
	}

	for _, filePath := range manifestFiles {
		templateData, err := os.ReadFile(filePath)
		if err != nil {
			return YAMLFilesNotFound{FileName: filePath, Message: err.Error()}
		}

		renderedData, err := renderTemplate(string(templateData), values)
		if err != nil {
			return ValuesReplacementError{Message: err.Error()}
		}

		var manifest map[string]interface{}
		if err := yaml.Unmarshal([]byte(renderedData), &manifest); err != nil {
			return YAMLUnmarshalError{Message: err.Error()}
		}

		outputFilePath := filepath.Join("./tmp", filepath.Base(filePath))
		if err := os.WriteFile(outputFilePath, []byte(renderedData), 0644); err != nil {
			return YAMLMarshalError{Message: err.Error()}
		}

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
