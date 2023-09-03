package search

import (
	"encoding/json"
	"fmt"
	"os"
)

type Participant struct {
	ID         string `json:"id"`
	Code       string `json:"student_code"`
	Name       string `json:"student_name"`
	Address    string `json:"student_address"`
	Occupation string `json:"student_occupation"`
	Reason     string `json:"joining_reason"`
}

type ParticipantsData struct {
	Participants []Participant `json:"participants"`
}

func FindStudentByCode(code string) (*Participant, error) {
	// Open the JSON file within the project directory
	file, err := os.Open("students.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var studentsData ParticipantsData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&studentsData)

	if err != nil {
		return nil, err
	}

	// Search student with the student code
	for _, student := range studentsData.Participants {
		if student.Code == code {
			return &student, nil
		}
	}

	return nil, fmt.Errorf("Student not found")
}
