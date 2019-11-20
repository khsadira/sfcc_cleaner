package auth

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (app *App) getToken(loginURL string) AuthenticityToken {
	client := app.Client

	response, err := client.Get(loginURL)

	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	token, _ := document.Find("input[name='LoginForm_RegistrationDomain']").Attr("value")

	authenticityToken := AuthenticityToken{
		Token: token,
	}

	return authenticityToken
}

func (app *App) login(username string, password string, baseURL string, path string) string {
	client := app.Client

	loginURL := baseURL + path
	authenticityToken := app.getToken(loginURL)

	data := url.Values{
		"LoginForm_RegistrationDomain": {authenticityToken.Token},
		"LoginForm_Login":              {username},
		"LoginForm_Password":           {password},
	}

	response, err := client.PostForm(loginURL, data)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(resp)
}

func scrapeAuth(user string, pass string, baseURL string, path string) bool {
	jar, _ := cookiejar.New(nil)

	app := App{
		Client: &http.Client{Jar: jar},
	}

	resp := app.login(user, pass, baseURL, path)
	if strings.Contains(resp, "Invalid login or password") {
		return false
	}
	return true
}
