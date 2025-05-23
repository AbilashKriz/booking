package renders

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AbilashKriz/bookings/pkg/config"
	"github.com/AbilashKriz/bookings/pkg/models"
)

// Creating a variable to pass to the below function which is then used in the main fucntion to return the values of AppConfig variable with a map of templcatecahes
var app *config.AppConfig

//If you need to add the data to all the pages

func AddDeafultData(td *models.TempData) *models.TempData {
	return td
}

//Creating this function to run it in main func to get access to the app wide config data

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderingHtml(w http.ResponseWriter, tmpl string, td *models.TempData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TempCache
	} else {
		tc, _ = CreatingTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("couldnt get the template")
	}

	buf := new(bytes.Buffer)

	AddDeafultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

	// err = t.Execute(w, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func CreatingTemplateCache() (map[string]*template.Template, error) {
	tempCache := map[string]*template.Template{}
	file, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range file {
		filename := filepath.Base(name)

		tmplPointer, err := template.New(filename).ParseFiles(name)
		if err != nil {
			log.Fatal(err)
			return tempCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Fatal(err)
			return tempCache, err
		}

		if len(layouts) > 0 {
			tmplPointer, err = tmplPointer.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tempCache, err
			}
			tempCache[filename] = tmplPointer
		}

	}
	return tempCache, nil
}
