package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"school/model"
	"school/subject"
	"school/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type SubjectHandler struct {
	subjectUsecase subject.SubjectUsecase
}

func CreateSubjectHandler(r *mux.Router, subjectUsecase subject.SubjectUsecase) {
	subjectHandler := SubjectHandler{subjectUsecase}

	r.HandleFunc("/subject", subjectHandler.viewAllSubject).Methods(http.MethodGet)
	r.HandleFunc("/subject/{id}", subjectHandler.viewSubjectById).Methods(http.MethodGet)
	r.HandleFunc("/subject", subjectHandler.insertSubject).Methods(http.MethodPost)
	r.HandleFunc("/subject/{id}", subjectHandler.updateSubject).Methods(http.MethodPut)
	r.HandleFunc("/subject/{id}", subjectHandler.deleteSubject).Methods(http.MethodDelete)
}

func (e *SubjectHandler) viewAllSubject(resp http.ResponseWriter, req *http.Request) {
	subject, err := e.subjectUsecase.ViewAll()
	if err != nil {
		utils.HandleError(resp, err.Error())
	}
	utils.HandleSuccess(resp, subject)
}

func (e *SubjectHandler) viewSubjectById(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "Id must be a number")
		return
	}
	subject, err := e.subjectUsecase.ViewByid(id)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, subject)
}

func (e *SubjectHandler) insertSubject(resp http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("[subject_handler.InserSubject] failed read data body ioutil %v", err)
		utils.HandleError(resp, "Opss server somting wrong")
		return
	}

	var newSubject = model.Subject{}
	err = json.Unmarshal(reqBody, &newSubject)
	if err != nil {
		fmt.Println("[subject_handler.InsertSubject] failed to unmarshal data")
		utils.HandleError(resp, "Opss server somting wrong")
		return
	}
	subject, err := e.subjectUsecase.Insert(&newSubject)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, subject)
}

func (e* SubjectHandler) updateSubject(resp http.ResponseWriter, req *http.Request) {
	muxvar := mux.Vars(req)
	strId := muxvar["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "Id must be a number")
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("[subject_handler.UpdateSubject] failed read data body ioutil %v", err)
		utils.HandleError(resp, "Opss server somting wrong")
		return
	}
	var newSubject = model.Subject{}
	err = json.Unmarshal(reqBody, &newSubject)
	if err != nil {
		fmt.Println("[subject_handler.updateSubject] failed to unmarshal data")
		utils.HandleError(resp, "Opss server somting wrong")
		return
	}
	subject, err := e.subjectUsecase.Update(id, &newSubject)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, subject)
}

func (e *SubjectHandler) deleteSubject(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "Id must be a number")
		return
	}
	err = e.subjectUsecase.Delete(id)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, "success delete object")
}