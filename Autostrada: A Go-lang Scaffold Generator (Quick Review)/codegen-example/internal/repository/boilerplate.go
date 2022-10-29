package repository

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type TemplateParams struct {
	Entity string
}

func GenerateBoilerplate(name string) error {
	tmpl, err := template.ParseFiles("internal/repository/boilerplate.tmpl")
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("internal/repository/%v_repository.go", name))
	defer f.Close()
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	//capitalize the first letter
	err = tmpl.Execute(f, &TemplateParams{Entity: strings.Title(name)})
	if err != nil {
		return fmt.Errorf("failed to generate template: %v", err)
	}
	return nil
}
