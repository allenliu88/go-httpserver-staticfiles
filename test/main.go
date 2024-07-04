package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/allenliu88/go-httpserver-staticfiles/utils"
)

type AnsibleRepo struct {
	AnsibleRepo string
}

// 获取软链接绝对目录
func BuildSoftLinkAbsolutePath(path string) string {
	cmd := exec.Command("readlink", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n, err:\n%v\n", string(out), err)
	}
	if len(out) > 0 {
		fmt.Printf("combined out:\n%s\n", string(out))
	}

	return string(out)
}

// 更改目录权限
func ChmodAbsolutePath(path string, mode string) {
	if !filepath.IsAbs(path) {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fmt.Sprintf("sudo chmod %s %s", mode, path))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n, err:\n%v\n", string(out), err)
	}
	if len(out) > 0 {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
}

func main() {
	abs := BuildSoftLinkAbsolutePath("a.yml")
	fmt.Println("SoftLink Path: " + abs)
	ChmodAbsolutePath(abs, "777")
	fmt.Println("File Path: [" + BuildSoftLinkAbsolutePath("requirements.yml") + "]")

	dir, _ := os.Getwd()
	fmt.Println(dir)

	// bizRepo, _ := utils.GetBizAnsibleRepo(filepath.Join(dir, "requirements2.yml"))
	bizRepo, _ := utils.GetBizAnsibleRepo("~/Downloads/requirements2.yml")
	fmt.Println("Output: " + bizRepo)

	// err := utils.BuildRequirementsFromTemplate(filepath.Join(dir, "requirements.yml"), filepath.Join(dir, "requirements-new.yml"), utils.RequirementValues{AnsibleRepo: bizRepo})
	err := utils.BuildRequirementsFromTemplate("~/Downloads/requirements.yml", filepath.Join(dir, "requirements-new.yml"), utils.RequirementValues{AnsibleRepo: bizRepo})

	fmt.Printf("Error: %v", err)
}
