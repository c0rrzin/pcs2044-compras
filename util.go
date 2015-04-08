package main

import (
	"fmt"
	"net/http"
)

func Success(w http.ResponseWriter, msg string) {
	json := JSONMessage(msg)
	fmt.Fprintf(w, json)
}

func JSONMessage(msg string) string {
	return fmt.Sprintf(`{ message: "%s" }`, msg)
}
