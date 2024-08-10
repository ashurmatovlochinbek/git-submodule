package submodule

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

var HandlerF1 = func(writer http.ResponseWriter, request *http.Request) {
	uuidValue := uuid.New()
	json, _ := json.Marshal(uuidValue)
	writer.Write(json)
}

//var HandlerF2 = func(writer http.ResponseWriter, request *http.Request) {
//	getValue := "Getting value"
//	json, _ := json.Marshal(getValue)
//	writer.Write(json)
//}
