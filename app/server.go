package main


import (
//	"os"
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var conftoml ConfToml

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
   	fmt.Fprintf(w, conftoml.Keys.Access_key)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Serve()
}
