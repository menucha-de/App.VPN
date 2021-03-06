/*
 * App vpn
 *
 * API version: 1.0.0
 * Contact: info@menucha.de
 */

package main

import (
	"flag"
	"fmt"
	"html"
	"net/http"

	"github.com/menucha-de/App.VPN/vpn"
	"github.com/menucha-de/logging"
	"github.com/menucha-de/utils"
)

var log *logging.Logger = logging.GetLogger("vpn")

func main() {
	var port = flag.Int("p", 8080, "port")
	var host = flag.String("h", "127.0.0.1", "host")
	flag.Parse()

	vpn.AddRoutes(logging.LogRoutes)
	vpn.AddRoutes(vpn.Routes)
	router := vpn.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFound)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", *host, *port),
		Handler: router,
	}

	log.Info("Server Started on port ", *port)

	log.Fatal(srv.ListenAndServe())
}

func notFound(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == "GET") {
		w.WriteHeader(404)
	}
	file := "./www" + html.EscapeString(r.URL.Path)
	if file == "./www/" {
		file = "./www/index.html"
	}
	if utils.FileExists(file) {
		http.ServeFile(w, r, file)
	} else {
		w.WriteHeader(404)
	}
}
