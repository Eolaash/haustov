package haustov

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// headers - func to show headers
func headers(w http.ResponseWriter, req *http.Request) {
	for name, values := range req.Header {
		// coz we can have multivalues for single header - []string map to parse
		for _, val := range values {
			fmt.Fprintf(w, "%v: %v\n", name, val)
		}
	}
}

// postReply - TASK5 main func
func postReply(w http.ResponseWriter, req *http.Request) {

	// Only POST filter
	if req.Method != "POST" {
		http.Error(w, "Sorry, only POST methods are supported", 404)
		//fmt.Fprintf(w, "Sorry, only POST method supported")
		return
	}

	// Read body (+ error of reading handle - return as response and quit func)
	// Defer for autoclose body afterway
	tBody, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//fmt.Println(req.Header["Content-Type"]) // to get body type by content-type in header
	//w.Header().Set("Content-Type", req.Header["Content-Type"][0]) // return content-type of req to resp
	//fmt.Fprintf(w, tBody)
	w.Write(tBody)
}

// StartHTTPServer - starting HTTP server for TASK5 (in VAR - Port, HOST - localhost by default)
func StartHTTPServer() {

	// SERVER default ADDR
	fServerAddr := ":8882"

	// headers enlist as reply
	http.HandleFunc("/headers", headers)

	// task5 main func - POST echo to reply
	http.HandleFunc("/postreply", postReply)

	// Server START and LISTEN (+ error check)
	fmt.Printf("Starting server <" + fServerAddr + "> HTTP POST...\n")
	if tError := http.ListenAndServe(fServerAddr, nil); tError != nil {
		fmt.Printf("Error occured on starting server:" + tError.Error())
	}
}
