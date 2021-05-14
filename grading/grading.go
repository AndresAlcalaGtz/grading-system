package grading

import (
	"fmt"
	"strconv"
)

type Container struct {
	Student string
	Subject string
	Grade   string
}

type Dataset struct {
	Students map[string]map[string]string
	Subjects map[string]map[string]string
}

func (ds *Dataset) SetGradeStudent(cnt Container) {
	_, errorStudent := ds.Students[cnt.Student][cnt.Subject]
	_, errorSubject := ds.Subjects[cnt.Subject][cnt.Student]
	if errorStudent || errorSubject {
		return
	}

	_, registeredStudent := ds.Students[cnt.Student]
	if !registeredStudent {
		ds.Students[cnt.Student] = map[string]string{cnt.Subject: cnt.Grade}
	} else {
		ds.Students[cnt.Student][cnt.Subject] = cnt.Grade
	}

	_, registeredSubject := ds.Subjects[cnt.Subject]
	if !registeredSubject {
		ds.Subjects[cnt.Subject] = map[string]string{cnt.Student: cnt.Grade}
	} else {
		ds.Subjects[cnt.Subject][cnt.Student] = cnt.Grade
	}
}

func (ds *Dataset) GetAverageStudent(stu string) string {
	_, errorStudent := ds.Students[stu]
	if !errorStudent {
		return "the student has not been registered"
	}

	var average float64 = 0
	for _, grd := range ds.Students[stu] {
		number, _ := strconv.ParseFloat(grd, 64)
		average += number
	}
	average /= float64(len(ds.Students[stu]))

	return fmt.Sprint(average)
}

func (ds *Dataset) GetAverageSubject(sub string) string {
	_, errorSubject := ds.Subjects[sub]
	if !errorSubject {
		return "the subject has not been registered"
	}

	var average float64 = 0
	for _, grd := range ds.Subjects[sub] {
		number, _ := strconv.ParseFloat(grd, 64)
		average += number
	}
	average /= float64(len(ds.Subjects[sub]))

	return fmt.Sprint(average)
}

func (ds *Dataset) GetAverageGeneral() string {
	errorStudent := len(ds.Students)
	errorSubject := len(ds.Subjects)
	if errorStudent == 0 && errorSubject == 0 {
		return "there has not been registered neither a student nor a subject"
	}

	var averages []float64
	var average float64

	for _, sub := range ds.Students {
		average = 0
		for _, grd := range sub {
			number, _ := strconv.ParseFloat(grd, 64)
			average += number
		}
		average /= float64(len(sub))
		averages = append(averages, average)
	}

	average = 0
	for _, avg := range averages {
		average += avg
	}
	average /= float64(len(averages))

	return fmt.Sprint(average)
}
