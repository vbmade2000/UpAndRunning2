package routes

import (
	"github.com/MarvinMenzerath/UpAndRunning2/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/MarvinMenzerath/UpAndRunning2/Godeps/_workspace/src/github.com/op/go-logging"
	"github.com/MarvinMenzerath/UpAndRunning2/lib"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := SiteData{lib.GetConfiguration().Dynamic.Title}
	t, err := template.ParseFiles("views/index.html", "views/partials/styles.html", "views/partials/footer.html", "views/partials/scripts.html")

	if t != nil {
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, data)
	} else {
		logging.MustGetLogger("logger").Error("Error while parsing Template: ", err)
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
	}
}
