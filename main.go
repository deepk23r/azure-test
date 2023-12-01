package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	opaPolicy "github.com/open-policy-agent/conftest/policy"
)

var (
	//go:embed results
	policyDir embed.FS
)

func removeAll(tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		fmt.Printf("Failed to remove temporary directory!! %s", tempDir)
	}
}

func loadOpaEngine() (*opaPolicy.Engine, error) {
	policyDir, err := createPolicyDir()
	if err != nil {
		return nil, err
	}
	defer removeAll(policyDir)

	fmt.Println("TempDir before opa load data -", policyDir)

	opaEngine, err := opaPolicy.LoadWithData(context.Background(), []string{policyDir}, make([]string, 0), "")
	if err != nil {
		return nil, err
	}

	fmt.Println("TempDir after opa load data -", policyDir)

	return opaEngine, nil
}

func createPolicyDir() (string, error) {
	tempPolicyDir, err := os.MkdirTemp("", "rego-policies_*")
	if err != nil {
		fmt.Println("os.MkdirTemp Error", err)
	}
	fmt.Println("TempDir - ", tempPolicyDir)
	fmt.Println()

	// Check the permissions of the current working directory
	fi, err := os.Stat(tempPolicyDir)
	if err != nil {
		fmt.Println("os.Stat Error", err)
		return "", err
	}
	// Print the permissions of the current working directory
	fmt.Println("Permissions - ", fi.Mode())
	fmt.Println()

	// recursively iterate over the embedded policyDir
	err = fs.WalkDir(policyDir, ".", func(path string, d fs.DirEntry, err error) error {

		// check if it is a rego policy file
		if !fs.ModeAppend.IsDir() && filepath.Ext(path) == ".rego" {
			// read the embedded rego policy file
			data, err := policyDir.ReadFile(path)
			if err != nil {
				fmt.Println("ReadFile error", err)
			}

			// add the rego policy file to the temporary directory
			tempFile := filepath.FromSlash(fmt.Sprintf("%s/%s", tempPolicyDir, filepath.Base(path)))
			if err := os.WriteFile(tempFile, data, 0600); err != nil {
				fmt.Println("WriteFile error", err)
			}
		}
		return err
	})

	if err != nil {
		fmt.Println("Walking folder policyDir error", err)
	}

	err = filepath.Walk(tempPolicyDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Walking folder tempDir error", err)
		}

		fmt.Println()
		fmt.Println("files in Temp folder", tempPolicyDir)
		fmt.Println(path, info.Size())
		return nil
	})

	if err != nil {
		fmt.Println("Error checking path of files in temp folder", err)
	}
	return tempPolicyDir, nil
}

func main() {

	//Get all env variables
	fmt.Println("Get environment variables")
	fmt.Println(os.Environ())
	fmt.Println()

	engine, err := loadOpaEngine()

	if err != nil {
		fmt.Println("Error loading opa engine", err)
	}

	fmt.Println(" loading opa engine", engine.Policies())

}
