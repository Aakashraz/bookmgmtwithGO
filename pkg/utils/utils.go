package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		// Unmarshal the body and copy to x
		if err := json.Unmarshal(body, x); err != nil {
			log.Printf("error unmarshalling JSON: %v", err)
			return
		}
	}
}
