package apiwriter

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	ApiWriter ApiWriterInf
)

const (
	SuccesStatusCode int = 200
	FailStatusCode   int = 500
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type ApiWriterInf interface {
	WriteSuccesaResp(w http.ResponseWriter, data interface{})
	WriteFailResp(w http.ResponseWriter, err error)
}

type ApiWriterImpl struct{}

func NewAPIWriter() {
	ApiWriter = &ApiWriterImpl{}
}

func (aw *ApiWriterImpl) WriteSuccesaResp(w http.ResponseWriter, data interface{}) {
	var resp Response
	resp = Response{
		StatusCode: SuccesStatusCode,
		Data:       data,
		Error:      nil,
	}

	finalResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(SuccesStatusCode)
	w.Write(finalResp)
}

func (aw *ApiWriterImpl) WriteFailResp(w http.ResponseWriter, err error) {
	var resp Response
	resp = Response{
		StatusCode: FailStatusCode,
		Data:       nil,
		Error:      err.Error(),
	}

	finalResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(FailStatusCode)
	w.Write(finalResp)
}
