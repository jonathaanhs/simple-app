package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Students struct {
	StudentID         int
	Name              string
	MidTermScore      float64
	SemesterTestScore float64
	AttendanceScore   float64
}

type FinalData struct {
	StudentID      int
	Name           string
	FinalScore     float64
	Grade          string
	TotalStudent   int
	PassingStudent int
	FailedStudent  int
}

func main() {
	var numberOfStudent int
	var result []FinalData

	fmt.Println("Input the number of students : ")

	fmt.Scanln(&numberOfStudent)

	studentData := make([]Students, numberOfStudent)

	totalPassingStudent := 0
	totalFailedStudent := 0
	for i, v := range studentData {
		var tmp FinalData

		fmt.Printf("==== Input student No %d ====\n", i+1)
		fmt.Println("Student ID : ")
		fmt.Scanln(&v.StudentID)
		fmt.Println("Name : ")
		fmt.Scanln(&v.Name)
		fmt.Println("Mid Term Score : ")
		fmt.Scanln(&v.MidTermScore)
		fmt.Println("Semester Term Score : ")
		fmt.Scanln(&v.SemesterTestScore)
		fmt.Println("Attendance Score : ")
		fmt.Scanln(&v.AttendanceScore)

		tmp.FinalScore = (0.2 * v.AttendanceScore) + (0.4 * v.MidTermScore) + (0.4 * v.SemesterTestScore)

		if tmp.FinalScore >= 85 && tmp.FinalScore <= 100 {
			tmp.Grade = "A"
			totalPassingStudent++
		} else if tmp.FinalScore >= 76 && tmp.FinalScore <= 84 {
			tmp.Grade = "B"
			totalPassingStudent++
		} else if tmp.FinalScore >= 61 && tmp.FinalScore <= 75 {
			tmp.Grade = "C"
			totalPassingStudent++
		} else if tmp.FinalScore >= 46 && tmp.FinalScore <= 60 {
			tmp.Grade = "D"
			totalFailedStudent++
		} else if tmp.FinalScore >= 0 && tmp.FinalScore <= 45 {
			tmp.Grade = "E"
			totalFailedStudent++
		}

		tmp.TotalStudent = numberOfStudent
		tmp.Name = v.Name
		tmp.StudentID = v.StudentID

		result = append(result, tmp)
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)

	fmt.Fprintln(w, "==========================================================")
	fmt.Fprintln(w, "No.\tStudentID\tName\tFinalScore\tGrade\t")
	fmt.Fprintln(w, "==========================================================")
	for i, v := range result {
		fmt.Fprintln(w, i+1, "\t", v.StudentID, "\t", v.Name, "\t", v.FinalScore, "\t", v.Grade, "\t")
	}
	fmt.Fprintln(w, "==========================================================")
	w.Flush()
	fmt.Println("Number of Students : ", numberOfStudent)
	fmt.Println("Number of Passing Students : ", totalPassingStudent)
	fmt.Println("Number of Failed Students : ", totalFailedStudent)
}
