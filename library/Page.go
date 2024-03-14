package library

type Page struct {
	Title string
	Data  interface{}
}

func NewPage(title string, data interface{}) *Page {
	return &Page{
		Title: title,
		Data:  data,
	}
}

type Pages map[string]*Page

var pages Pages = make(map[string]*Page)

func GetPages() Pages {
	if len(pages) == 0 {
		pages["index"] = NewPage("Index", "World")
	}

	return pages
}
