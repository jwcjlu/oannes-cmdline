package tpl

import (
	"github.com/iancoleman/strcase"
	"text/template"
)

var FuncMap = template.FuncMap{
	"camelcase": strcase.ToCamel,
	"snakecase": strcase.ToSnake,
}
