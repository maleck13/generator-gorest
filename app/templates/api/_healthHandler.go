package api

import (
	"net/http"
	"encoding/json"
)

func HealthHandler(rw http.ResponseWriter, req *http.Request) HttpError {
	//fill in health checks here
	return nil
}

func Ping(rw http.ResponseWriter, req *http.Request) HttpError{
	res := make(map[string]string)
	res["ok"] = "200"
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(res); err != nil{
		return NewHttpError(err,http.StatusInternalServerError)
	}
	return nil
}
