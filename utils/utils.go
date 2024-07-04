package utils

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Requirements struct {
	Src     string `yaml:"src"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type RequirementValues struct {
	AnsibleRepo string
}

func GetBizAnsibleRepo(bizRequirementsFilePath string) (string, error) {
	bizRequirementsFile, err := os.ReadFile(bizRequirementsFilePath)
	if err != nil {
		return "", err
	}

	bizRequirements := []Requirements{}
	err = yaml.Unmarshal(bizRequirementsFile, &bizRequirements)
	if err != nil {
		return "", err
	}

	for _, item := range bizRequirements {
		// 返回第一个依赖项的仓库地址
		segs := strings.Split(item.Src, "/")
		return strings.Join(segs[0:len(segs)-1], "/"), nil
	}

	return "", errors.New("not found")
}

// 获取绝对目录
func BuildAbsolutePath(path string) string {
	cmd := exec.Command("dirname", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	return string(out)
}

func BuildRequirementsFromTemplate(requirementsTemplateFilePath string,
	requirementsFilePathDest string, requirRequirementValues RequirementValues) error {

	if path.IsAbs(requirementsTemplateFilePath) {
		fmt.Println("Absolute")
	} else {
		fmt.Println("Not Absolute")
	}

	fmt.Println(BuildAbsolutePath(requirementsTemplateFilePath))

	requirementsTemplateFilePath, err := filepath.Abs(requirementsTemplateFilePath)

	// 创建模板
	tpl, err := template.New(path.Base(requirementsTemplateFilePath)).ParseFiles(requirementsTemplateFilePath)
	if err != nil {
		return err
	}

	// 创建目标文件
	dest, err := os.Create(requirementsFilePathDest)
	if err != nil {
		return err
	}

	// 依据模板生成需求文件
	err = tpl.Execute(dest, requirRequirementValues)
	if err != nil {
		return err
	}

	return nil
}
