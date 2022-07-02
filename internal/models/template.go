package models

type TemplateData struct {
	Data map[string]Bag
}

type Bag interface {
	string | int | any
}
