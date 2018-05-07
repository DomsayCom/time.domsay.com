package lib

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(id string, r http.ResponseWriter) {

	value := map[string]string{
		"uid": id,
	}

	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {

		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(r, cookie)

	}

}

func ClearSession(r http.ResponseWriter) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(r, cookie)

}

func CookieValid(r *http.Request) (uid string) {

	cookie, err := r.Cookie("session")
	if err == nil {

		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {

			uid = cookieValue["uid"]

		}

	}

	return uid

}

func CheckAuth(w http.ResponseWriter, r *http.Request, c bool, route string) {

	cookieExist := CookieValid(r)

	if c == true && cookieExist != "" {

		AppRedirect(w, r, route, 302)

	}

	if c == false && cookieExist == "" {

		AppRedirect(w, r, route, 302)

	}

}

func AppRedirect(w http.ResponseWriter, r *http.Request, route string, status int) {

	http.Redirect(w, r, "http://localhost:8080"+route, status)

}
