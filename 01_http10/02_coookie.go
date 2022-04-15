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
	cookie_response(writer, request)
}

func cookie_response(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := request.Header["Cookie"]; ok {
		fmt.Fprintf(writer, "<html><body>VISITED~</body></html>\n")
	} else {
		fmt.Fprintf(writer, "<html><body>FIRST VISIT!!</body></html>\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
