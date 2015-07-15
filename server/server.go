package server

import (
	//"github.com/codegangsta/negroni"
	//"github.com/rprakashg/go-o365api-explorer/office365"
	"log"
	"net/http"
	"os"
)

func Start() {
	projdir := os.Getenv("GOPATH") + "/src/github.com/rprakashg/go-o365api-explorer"
	pubdir := projdir + "/public/"
	certdir := projdir + "/server/certs"

	router := NewRouter()
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(pubdir))))
	/*  router.Handle("/get", negroni.New(
		negroni.HandlerFunc(office365.IsAuthorized),
		negroni.Wrap(http.HandlerFunc(getHandler)),
	))*/
	http.Handle("/", router)
	err := http.ListenAndServeTLS(":10443", certdir+"/cert.pem", certdir+"/key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
