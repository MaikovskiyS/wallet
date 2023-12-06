package apperrors

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrResponse(w http.ResponseWriter, er *AppErr) {
	if er == nil {
		log.Println("nil err")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	resp := make(map[string]string, 1)
	w.WriteHeader(er.Code())
	resp["error"] = er.Error()

	rBytes, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(rBytes)

}
