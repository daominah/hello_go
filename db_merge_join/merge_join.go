package main

import (
	"encoding/json"
	"log"
)

type Student struct {
	Name      string
	CollegeId int64
	College   *College
}

type College struct {
	Id   int64
	Name string
}

// Join inner joins 2 tables
func MergeJoin(students []Student, colleges []College) []Student {
	rows := make([]Student, 0)
	for true {
		if len(students) < 1 || len(colleges) < 1 {
			break
		}
		lastStudent := students[len(students)-1]
		lastCollege := colleges[len(colleges)-1]
		if lastStudent.CollegeId > lastCollege.Id {
			students = students[:len(students)-1]
			continue
		}
		if lastStudent.CollegeId < lastCollege.Id {
			colleges = colleges[:len(colleges)-1]
			continue
		}
		rows = append(rows, Student{
			Name:      lastStudent.Name,
			CollegeId: lastStudent.CollegeId,
			College:   &lastCollege,
		})
		students = students[:len(students)-1]
		// do not pop colleges because students can have a same collegeId
	}
	return rows
}

func main() {
	// sorted by CollegeId
	students := []Student{
		{Name: "Hai", CollegeId: 3},
		{Name: "Tung", CollegeId: 3},
		{Name: "Lan", CollegeId: 4},
		{Name: "Huong", CollegeId: 6},
	}
	// sorted by CollegeId
	colleges := []College{
		{Id: 1, Name: "DH khoa hoc tu nhien"},
		{Id: 2, Name: "DH khoa hoc xa hoi"},
		{Id: 3, Name: "DH Back Khoa Ha Noi"},
		{Id: 4, Name: "DH kinh te quoc dan"},
		{Id: 5, Name: "HV ngan hang"},
	}
	joinedStudents := MergeJoin(students, colleges)
	beauty, _ := json.MarshalIndent(joinedStudents, "", "	")
	log.Printf("%s\n", beauty)
}
