package models

import "github.com/powsianik/thinking-in-code/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct{
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Message   string
	Warning   string
	Error     string
	Post      PostData
	Form *forms.Form
}
