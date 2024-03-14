package library

type Page struct {
	Title        string
	TemplateName string
	Data         interface{}
}

func NewPage(title string, templateName string, data interface{}) *Page {
	return &Page{
		Title:        title,
		TemplateName: templateName,
		Data:         data,
	}
}
