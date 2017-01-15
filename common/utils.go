package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type appErr struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

//errores presentes en la autenticación
type logErr struct {
	Message string `json:"message"`
}

func JsonError(w http.ResponseWriter, handlerErr error, code int) {
	log.Printf("Error: %s\n", handlerErr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if res, err := json.Marshal(&appErr{handlerErr.Error(), code}); err == nil {
		w.Write(res)
	}
}

func JsonOk(w http.ResponseWriter, res []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(res)
}

func JsonStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

//custom error
func (e *logErr) Error() string {
	return e.Message
}

func NewLogErr(msg string) error {
	return &logErr{msg}
}
