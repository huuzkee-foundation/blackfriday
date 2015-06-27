package main

import (
	"fmt"
	"html/template"
	"os"
	//"strings"
	"github.com/huuzkee-foundation/blackfriday"
)

// Data to put into  template
type Page struct {
	Title string
	Body  string
}

// The markdown
var markdownText string = `##ABC
> 123
# h1 Heading
## h2 Heading
### h3 Heading
#### h4 Heading
##### h5 Heading
###### h6 Heading`

// The template
var templateText string = `
<head>
  <title>{{.Title}}</title>
</head>

<body>
  {{.Body | markDown}}
</body>
`

// Real blackfriday functionality commented out, using strings.ToLower for demo
func markDowner(args ...interface{}) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...))))
	//return template.HTML(strings.ToLower(fmt.Sprintf("%s", args...)))
}

func main() {
	// Create a page
	p := &Page{Title: "A Test Demo", Body: markdownText}

	// Parse the template and add the function to the funcmap
	tmpl := template.Must(template.New("page.html").Funcs(template.FuncMap{"markDown": markDowner}).Parse(templateText))

	// Execute the template
	err := tmpl.ExecuteTemplate(os.Stdout, "page.html", p)
	if err != nil {
		fmt.Println(err)
	}
}
