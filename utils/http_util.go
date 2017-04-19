package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestBody(rw http.ResponseWriter, r *http.Request, model interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(model); err != nil {
		ResponseWithJson(rw, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()
}

func ResponseWithJson(rw http.ResponseWriter, httpcode int, respBody interface{}) {
	response, _ := json.Marshal(respBody)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpcode)
	rw.Write(response)
}
