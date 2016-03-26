package api

import (
	"encoding/json"
	"<%=basePackage %>/<%=baseName %>/config"
	"net/http"
)

//Example route handler
func IndexHandler(rw http.ResponseWriter, req *http.Request) HttpError {
	encoder := json.NewEncoder(rw)
	data := make(map[string]string)
	data["example"] = config.Conf.GetExample()
	if err := encoder.Encode(data); err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}
