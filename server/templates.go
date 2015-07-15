package server

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

//map to store collection of templates to be rendered
var templates = make(map[string]*template.Template)

//load the templates
func init() {
	rootPath := os.Getenv("GOPATH") + "/src/github.com/rprakashg/go-o365api-explorer"
	loadTemplate("layout", rootPath+"/views/layout.html")
	loadTemplate("home", rootPath+"/views/home.html")
}

// loadTemplate reads the specified template file for use.
func loadTemplate(name string, path string) {
	// Read the html template file.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a template using the content from file
	t, err := template.New(name).Parse(string(data))
	if err != nil {
		log.Fatalln(err)
	}

	if _, exists := templates[name]; exists {
		log.Fatalf("Template %s already loaded", name)
	}

	// Store the template for use.
	templates[name] = t
}

// executeTemplate executes the specified template with the specified variables.
func executeTemplate(name string, vars map[string]interface{}) []byte {
	markup := new(bytes.Buffer)
	if err := templates[name].Execute(markup, vars); err != nil {
		log.Println(err)
		return []byte("Error Processing Template")
	}

	return markup.Bytes()
}
