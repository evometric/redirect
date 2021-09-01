package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	httpSrv := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host, _, Err := net.SplitHostPort(r.Host)
			if Err != nil {
				host = r.Host // not in host:port syntax
			}
			u := r.URL
			u.Host = host
			u.Scheme = "https"
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}),
	}
	log.Println(httpSrv.ListenAndServe())
}
