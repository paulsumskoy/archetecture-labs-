package main

import (
	  "encoding/json"
    "fmt"
		"time"
    "net/http"
)

func timeHandler(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["time"] = time.Now().Format(time.RFC3339)
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		} else {
		  w.Write(jsonResp)
		}
}

func main() {
    http.HandleFunc("/time", timeHandler)
    http.ListenAndServe(":8795", nil)
}
