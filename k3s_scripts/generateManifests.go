package k3sscripts

import (
	"fmt"
	"os"
	"path/filepath"
)

func (ctx *Context) GenerateManifests() K3sError {
	tempDir := filepath.Join(os.TempDir(), "tufin-assignment")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return NewDirCreationError(filepath.Join(os.TempDir(), "tufin-assignment"), fmt.Sprintf("failed to create temp directory: %s", err.Error()))
	}

	manifestFiles := []string{
		"manifests/custom-local-path.yaml",
		"manifests/mysql-config-map.yaml",
		"manifests/wordpress-config-map.yaml",
		"manifests/mysql-pvc.yaml",
		"manifests/wordpress-pvc.yaml",
		"manifests/mysql-deployment.yaml",
		"manifests/wordpress-deployment.yaml",
	}

	for _, filePath := range manifestFiles {
		if err := ctx.renderAndSaveManifest(filePath, tempDir); err != nil {
			return err
		}
	}

	fmt.Println("Manifests successfully generated and saved to the OS temp directory.")
	return nil
}

func (ctx *Context) renderAndSaveManifest(filePath string, tempDir string) K3sError {
	templateData, err := ctx.manifests.ReadFile(filePath)
	if err != nil {
		return NewFileReadError(filePath, fmt.Sprintf("failed to read manifest file %s: %s", filePath, err.Error()))
	}

	outputFilePath := filepath.Join(tempDir, filepath.Base(filePath))
	if err := os.WriteFile(outputFilePath, templateData, 0644); err != nil {
		return NewFileWriteError(outputFilePath, fmt.Sprintf("failed to write rendered manifest to %s: %s", outputFilePath, err.Error()))
	}

	ctx.tempFiles = append(ctx.tempFiles, outputFilePath)
	fmt.Printf("Rendered manifest saved: %s\n", outputFilePath)
	return nil
}

func (ctx *Context) CleanupTempFiles() {
	for _, filePath := range ctx.tempFiles {
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("Failed to remove temp file %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Removed temp file: %s\n", filePath)
		}
	}
	ctx.tempFiles = []string{}
}
