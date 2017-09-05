package models

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// JSON function
func JSON(w http.ResponseWriter, r *http.Request, resp interface{}) {
	final, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(final)
}

// MakeTimestamp function
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
