package config

import "html/template"

type Config struct {
	TemplateCache map[string]*template.Template
	Port          string
}
