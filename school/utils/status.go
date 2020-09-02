package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school/model"
)

func HandleSuccess(resp http.ResponseWriter, data interface{}) {
	returnData := model.Respons{
		Success: true,
		Message: "SUCCESS",
		Data:    data,
	}

	jsonData, err := json.Marshal(returnData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[HandleSucess.utils] Error when do json Marshalling for error handling : %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

func HandleError(resp http.ResponseWriter, message string) {
	data := model.Respons{
		Success: false,
		Message: message,
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(resp).Encode(data)

}
