package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"school/model"
	studentHandler "school/student/handler"
	studentRepo "school/student/repository"
	studentUsecase "school/student/usecase"
	teacherHandler "school/teacher/handler"
	teacherRepo "school/teacher/repository"
	teacherUsecase "school/teacher/usecase"
	subjectRepo "school/subject/repository"
	subjectUsecase "school/subject/usecase"
	subjectHandler "school/subject/handler"
)

func main() {
	port := "8089"
	conStr := "root:@tcp(127.0.0.1:3306)/schoolme?parseTime=true"

	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + " : " + err.Error())
	}
	defer db.Close()

	db.Debug().AutoMigrate(
		&model.Student{},
		&model.Teacher{},
		&model.Subject{},
	)

	router := mux.NewRouter().StrictSlash(true)

	studentRepo := studentRepo.CreateStudentRepo(db)
	studentUsecase := studentUsecase.CreateStudentUsecase(studentRepo)
	studentHandler.CreateStudentHandler(router, studentUsecase)

	teacherRepo := teacherRepo.CreateTeacherRepo(db)
	teacherUsecase := teacherUsecase.CreateTeacherUsecase(teacherRepo)
	teacherHandler.CreateTeacherHandler(router, teacherUsecase)

	subjectRepo := subjectRepo.CreateSubjectRepo(db)
	subjectUsecase := subjectUsecase.CreateSubjectUsecase(subjectRepo)
	subjectHandler.CreateSubjectHandler(router, subjectUsecase)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
