package handler 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"school/model"
	"school/teacher"
	"school/utils"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
)

type TeacherHandler struct {
	teacherUsecase teacher.TeacherUsecase
}

func CreateTeacherHandler(r *mux.Router, teacherUsecase teacher.TeacherUsecase) {
	teacherHandler := TeacherHandler{teacherUsecase}

	r.HandleFunc("/teacher", teacherHandler.getAllTeacher).Methods(http.MethodGet)
	r.HandleFunc("/teacher/{id}", teacherHandler.getTeacherById).Methods(http.MethodGet)
	r.HandleFunc("/teachers/{name}", teacherHandler.getTeacherByName).Methods(http.MethodGet)
	r.HandleFunc("/teacher", teacherHandler.addTeacher).Methods(http.MethodPost)
	r.HandleFunc("/teacher/{id}", teacherHandler.updateTeacher).Methods(http.MethodPut)
	r.HandleFunc("/teacher/{id}", teacherHandler.deleteTeacher).Methods(http.MethodDelete)
}

func (e *TeacherHandler) getAllTeacher(resp http.ResponseWriter, req *http.Request) {
	teacher, err := e.teacherUsecase.ViewAll()
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, teacher)
}

func (e *TeacherHandler) getTeacherById(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "ID Must be a number")
		return
	}

	teacher, err := e.teacherUsecase.ViewById(id)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, teacher)
}

func (e *TeacherHandler) getTeacherByName(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	str := muxVar["name"]

	teacher, err := e.teacherUsecase.ViewByName(str)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, teacher)
}

func (e *TeacherHandler) addTeacher(resp http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("[TeacherHandler.addTeacher] Error when do iotil read body for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	var teacher = model.Teacher{}
	err = json.Unmarshal(reqBody, &teacher)
	if err != nil {
		fmt.Println("[TeacherHandler.addTeacher] Error when do json unmarshal for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}

	newTeacher, err := e.teacherUsecase.Insert(&teacher)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, newTeacher)
}

func (e *TeacherHandler) updateTeacher(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "ID Must be a number")
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("[TeacherHandler.updateTeacher] Error when do iotil read body for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	var teacher = model.Teacher{}
	err = json.Unmarshal(reqBody, &teacher)
	if err != nil {
		fmt.Println("[TeacherHandler.updateTeacher] Error when do json unmarshal for error handling : %v \n", err)
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	newTeacher, err := e.teacherUsecase.Update(id, &teacher)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, newTeacher)
}

func (e *TeacherHandler) deleteTeacher(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strId := muxVar["id"]
	
	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.HandleError(resp, "ID Must be a number")
		return
	}

	err = e.teacherUsecase.Delete(id)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}
	utils.HandleSuccess(resp, "Delete teacher success")
}