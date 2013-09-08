package main

import (
	"github.com/gorilla/sessions"
	"io"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func pageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "get_name_session")
	name, ok := session.Values["name"].(string)
	session.Values["name"] = "Dean"
	session.Save(r, w)
	if ok {
		io.WriteString(w, "hello, "+string(name))
	}
}

func main() {
	http.HandleFunc("/", pageHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
