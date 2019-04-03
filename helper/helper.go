package helper

import (
	"encoding/json"
	"net/http"
)

type Helper struct {}

func (u *Helper) RespondWithError(w http.ResponseWriter, code int, msg string) {
	u.RespondWithJson(w, code, map[string]string{"error": msg})
}

func (u *Helper) RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
