package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		http.Error(writer, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(writer, "<html><body>hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
