package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/k0kubun/pp"
)

func digest_handler(writer http.ResponseWriter, request *http.Request) {
	pp.Printf("URL: %s\n", request.URL.String())
	pp.Printf("Query: %s\n", request.URL.Query())
	pp.Printf("Proto: %s\n", request.Proto)
	pp.Printf("Method: %s\n", request.Method)
	pp.Printf("Header: %s\n", request.Header)

	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Printf("--body--\n%s\n", string(body))

	if _, ok := request.Header["Authorization"]; !ok {
		writer.Header().Add("WWW-Authenticate",
			`Digest
				realm="Secret Zone",
				nonce="TgLc25U2BQA=f410a2789473e18e6587be703c2e67fe2b04afd",
				algorithm=MD5,
				qop="auth"`)
		writer.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Fprintf(writer, "<html><body>secret page</body></html>\n")
	}
}

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
	http.HandleFunc("/digest", digest_handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
