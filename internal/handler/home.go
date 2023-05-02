package handler

import (
	"html/template"
	"net/http"

	g "github.com/bahodurnazarov/CatFacts/internal/getFacts"
	gt "github.com/bas24/googletranslatefree"
	d "github.com/bahodurnazarov/CatFacts/pkg/db"
	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {

	factEN := g.GetFacts()
	factRU, _ := gt.Translate(factEN, "en", "ru")

	tmpl, err := template.ParseFiles("../../pkg/assets/html/home.html")
	if err != nil {
		lg.Errl.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
	}

	db := d.ConnDB()
	if factEN != "" {
		_, err := db.Exec("INSERT into facts VALUES ($1, $2)", factEN, factRU)
		if err != nil {
			lg.Errl.Fatalf("111An error occured while executing query: %v", err)
		}
	}

	err = tmpl.Execute(c.Response().Writer, factEN)
	if err != nil {
		lg.Errl.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":   "HOME",
		"factEN": factEN,
		"factRU": factRU,
	})
}
