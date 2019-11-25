package utils


import (
	"github.com/khsadira/cleaner/auth"
	"html/template"
	"log"
	"net/http"
)

func GenerateForm(w http.ResponseWriter, send interface{}, path string) {
	key := auth.GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	w.Header().Set("Authorization", key)

	t := template.New("main")
	t = template.Must(t.ParseFiles(path))
	err := t.ExecuteTemplate(w, "main", send)

	if err != nil {
		log.Println(err)
	}
}