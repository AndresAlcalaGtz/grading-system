package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../grading"
)

const (
	FIELD = "Content-Type"
	TYPE  = "text/html"
)

var dataset grading.Dataset

func main() {
	dataset.Students = make(map[string]map[string]string)
	dataset.Subjects = make(map[string]map[string]string)

	http.HandleFunc("/", index)
	http.HandleFunc("/grade-student", gradeStudent)
	http.HandleFunc("/average-student", averageStudent)
	http.HandleFunc("/average-subject", averageSubject)
	http.HandleFunc("/average-general", averageGeneral)
	http.HandleFunc("/send", send)

	fmt.Println("the server is running")
	log.Fatal(http.ListenAndServe("localhost:80", nil))
}

func loadHTML(path string) string {
	html, _ := ioutil.ReadFile(path)
	return string(html)
}

func index(rsp http.ResponseWriter, _ *http.Request) {
	rsp.Header().Set(FIELD, TYPE)
	fmt.Fprintf(rsp, loadHTML("index.html"))
}

func gradeStudent(rsp http.ResponseWriter, _ *http.Request) {
	rsp.Header().Set(FIELD, TYPE)
	fmt.Fprintf(rsp, loadHTML("pages/grade-student.html"))
}

func averageStudent(rsp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}

	stu := req.FormValue("student")

	rsp.Header().Set(FIELD, TYPE)
	fmt.Fprintf(rsp, loadHTML("pages/average-student.html"), dataset.GetAverageStudent(stu))
}

func averageSubject(rsp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}

	sub := req.FormValue("subject")

	rsp.Header().Set(FIELD, TYPE)
	fmt.Fprintf(rsp, loadHTML("pages/average-subject.html"), dataset.GetAverageSubject(sub))
}

func averageGeneral(rsp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}

	rsp.Header().Set(FIELD, TYPE)
	fmt.Fprintf(rsp, loadHTML("pages/average-general.html"), dataset.GetAverageGeneral())
}

func send(rsp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}

	if req.Method == "POST" {
		cnt := grading.Container{
			Student: req.FormValue("student"),
			Subject: req.FormValue("subject"),
			Grade:   req.FormValue("grade"),
		}
		dataset.SetGradeStudent(cnt)

		rsp.Header().Set(FIELD, TYPE)
		fmt.Fprintf(rsp, loadHTML("pages/send.html"))
	}
}
