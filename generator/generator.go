package generator

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

const (
	mainTemplate       = "templates/main.gotmpl"
	controllerTemplate = "templates/controller.gotmpl"
	serviceTemplate    = "templates/service.gotmpl"
	middlewareTemplate = "templates/middleware.gotmpl"
)

type Component string

const (
	CONTROLLER Component = "controller"
	SERVICE    Component = "service"
	MIDDLEWARE Component = "middleware"
)

func CreateProject(projectName, module string) error {
	if err := os.Mkdir(projectName, 0755); err != nil {
		return fmt.Errorf("error creating project directory: %w", err)
	}

	dirs := []string{
		filepath.Join(projectName, "middlewares"),
		filepath.Join(projectName, "infrastructure"),
		filepath.Join(projectName, "modules", "foo", "domain"),
		filepath.Join(projectName, "modules", "foo", "services"),
		filepath.Join(projectName, "modules", "foo", "controllers"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	mainData := ProjectData{
		ProjectName: projectName,
		Module:      module,
		FileName:    "main.go",
	}

	fooController := ControllerData{
		Package:  "controllers",
		Name:     "Foo",
		FileName: "foo_controller.go",
	}

	fooService := ServiceData{
		Package:   "services",
		Name:      "Foo",
		FileName:  "foo_service.go",
		LowerName: "defaultFoo",
	}

	fooMiddleware := MiddlewareData{
		Package:  "middlewares",
		Name:     "Foo",
		FileName: "foo_middleware.go",
	}

	if err := processTemplate(mainTemplate, filepath.Join(projectName, mainData.FileName), mainData); err != nil {
		return err
	}

	if err := processTemplate(middlewareTemplate, filepath.Join(projectName, "middlewares", fooMiddleware.FileName), fooMiddleware); err != nil {
		return err
	}

	if err := processTemplate(serviceTemplate, filepath.Join(projectName, "modules", "foo", "services", fooService.FileName), fooService); err != nil {
		return err
	}

	if err := processTemplate(controllerTemplate, filepath.Join(projectName, "modules", "foo", "controllers", fooController.FileName), fooController); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = projectName

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing 'go mod init %s' in project's folder: %w, %s", module, err, string(output))
	}

	return nil
}

func CreateController(name, location string) error {
	targetDir := location
	if targetDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		targetDir = wd
	}

	fileName := extractFileName(name, CONTROLLER)
	pack := extractPackageName(targetDir)

	fmt.Printf("package: %s | file: %s\n", pack, fileName)

	controller := ControllerData{
		FileName: fileName,
		Package:  pack,
		Name:     name,
	}

	if err := processTemplate(controllerTemplate, filepath.Join(location, fileName), controller); err != nil {
		return err
	}

	return nil
}

func CreateService(name, location string) error {
	targetDir := location
	if targetDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		targetDir = wd
	}

	fileName := extractFileName(name, SERVICE)
	pack := extractPackageName(targetDir)

	fmt.Printf("package: %s | file: %s\n", pack, fileName)

	service := ServiceData{
		FileName:  fileName,
		Package:   pack,
		Name:      name,
		LowerName: strings.Replace(name, string(name[0]), strings.ToLower(string(name[0])), 1),
	}

	if err := processTemplate(serviceTemplate, filepath.Join(location, fileName), service); err != nil {
		return err
	}

	return nil
}

func CreateMiddleware(name, location string) error {
	targetDir := location
	if targetDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		targetDir = wd
	}

	fileName := extractFileName(name, CONTROLLER)
	pack := extractPackageName(targetDir)

	fmt.Printf("package: %s | file: %s\n", pack, fileName)

	middleware := MiddlewareData{
		FileName: fileName,
		Package:  pack,
		Name:     name,
	}

	if err := processTemplate(middlewareTemplate, filepath.Join(location, fileName), middleware); err != nil {
		return err
	}

	return nil
}

func extractFileName(name string, component Component) string {
	var suffix string
	switch component {
	case CONTROLLER:
		suffix = "controller"
	case SERVICE:
		suffix = "service"
	case MIDDLEWARE:
		suffix = "middleware"
	default:
		panic(errors.New("invalid component type"))
	}

	var builder strings.Builder
	for i, char := range name {
		if unicode.IsUpper(char) {
			if i > 0 {
				builder.WriteRune('_')
			}

			builder.WriteRune(unicode.ToLower(char))
		} else {
			builder.WriteRune(char)
		}
	}

	return fmt.Sprintf("%s_%s.go", builder.String(), suffix)
}

func extractPackageName(dir string) string {
	folders := strings.Split(dir, "/")
	pack := string(folders[len(folders)-1])
	if pack == "" {
		pack = folders[len(folders)-2]
	}

	return pack
}

func processTemplate(templatePath, destPath string, data interface{}) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
