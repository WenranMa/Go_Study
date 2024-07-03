// server/selectsort.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func add(x, y int) int {
	return x + y
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	var param AddParam
	json.Unmarshal(b, &param)
	ret := add(param.X, param.Y)
	fmt.Println("Add function called: ", param)
	respBytes, _ := json.Marshal(AddResult{Code: 0, Data: ret})
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
