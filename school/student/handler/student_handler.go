package handler

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"school/model"
	"school/student"
	"school/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	studentUsecase student.StudentUsecase
}

func CreateStudentHandler(r *mux.Router, studentUsecase student.StudentUsecase) {
	studentHandler := StudentHandler{studentUsecase}

	r.HandleFunc("/student", studentHandler.getAllStudent).Methods(http.MethodGet)
	r.HandleFunc("/student", studentHandler.addStudent).Methods(http.MethodPost)
	r.HandleFunc("/student/{idStudent}", studentHandler.getStudentById).Methods(http.MethodGet)
	r.HandleFunc("/students/{name}", studentHandler.getStudentByName).Methods(http.MethodGet)
	r.HandleFunc("/student/{idStudent}", studentHandler.updateStudent).Methods(http.MethodPut)
	r.HandleFunc("/student/{idStudent}", studentHandler.deleteStudent).Methods(http.MethodDelete)
}

func (e *StudentHandler) getAllStudent(resp http.ResponseWriter, req *http.Request) {
	student, err := e.studentUsecase.ViewAll()
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, student)
}

func (e *StudentHandler) getStudentById(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["idStudent"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "ID Must be a number")
		return
	}

	student, err := e.studentUsecase.ViewById(id)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, student)
}

func (e *StudentHandler) getStudentByName(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	name := muxVar["name"]

	student, err := e.studentUsecase.ViewByName(name)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, student)
}

func (e *StudentHandler) updateStudent(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["idStudent"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "ID Must be a number")
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("[StudentHandler.updateStudent] Error when do iotil read body for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}

	reqStudent := model.Student{}
	err = json.Unmarshal(reqBody, &reqStudent)
	if err != nil {
		fmt.Println("[StudentHandler.updateStudent] Error when do json unmarshal for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}

	newStudent, err := e.studentUsecase.Update(id, &reqStudent)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, newStudent)
}

func (e *StudentHandler) addStudent(resp http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("[StudentHandler.addStudent] Error when do iotil read body for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	reqStudent := model.Student{}
	err = json.Unmarshal(reqBody, &reqStudent)
	if err != nil {
		fmt.Println("[StudentHandler.addStudent] Error when do json unmarshal for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}

	newStudent, err := e.studentUsecase.Insert(&reqStudent)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, newStudent)
}

func (e *StudentHandler) deleteStudent(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["idStudent"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "ID Must be a number")
		return
	}
	err = e.studentUsecase.Delete(id)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, "Success delete student")
}

