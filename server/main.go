package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func dumpJsonRequestHandlerFunc(w http.ResponseWriter, req *http.Request){
  //Validate request
  if req.Method != "POST" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }


  if req.Header.Get("Content-Type") != "application/json" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  //To allocate slice for request body
  length, err := strconv.Atoi(req.Header.Get("Content-Length"))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Println("27error")
    return
  }

  //Read body data to parse json
  body := make([]byte, length)
  length, err = req.Body.Read(body)
  if err != nil && err != io.EOF {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  //parse json
  var jsonBody map[string]interface{}
  err = json.Unmarshal(body[:length], &jsonBody)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  fmt.Printf("%v\n", jsonBody)

  w.WriteHeader(http.StatusOK)  
}

func main(){
  http.HandleFunc("/json", dumpJsonRequestHandlerFunc)
  http.ListenAndServe(":8080", nil)
}
