package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/allenliu88/go-httpserver-staticfiles/utils"
)

type AnsibleRepo struct {
	AnsibleRepo string
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	bizRepo, _ := utils.GetBizAnsibleRepo(filepath.Join(dir, "requirements2.yml"))
	fmt.Println("Output: " + bizRepo)

	err := utils.BuildRequirementsFromTemplate(filepath.Join(dir, "requirements.yml"), filepath.Join(dir, "requirements-new.yml"),
		utils.RequirementValues{AnsibleRepo: bizRepo})

	fmt.Printf("Error: %v", err)
}
