package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Definisikan struktur data Student
type Student struct {
	ID          string `json:"id"`
	StudentCode string `json:"student_code"`
	Name        string `json:"student_name"`
	Address     string `json:"student_address"`
	Occupation  string `json:"student_occupation"`
	Reason      string `json:"joining_reason"`
}

// Definisikan struktur data Students yang berisi slice dari Student
type Students struct {
	Participants []Student `json:"participants"`
}

func main() {
	// Periksa jumlah argumen yang diberikan pada program
	if len(os.Args) != 2 {
		fmt.Println("Harap masukkan kode mahasiswa setelah .go")
		return
	}

	// Buka file JSON yang berisi data mahasiswa
	jsonFile, err := os.Open("participants.json")
	if err != nil {
		fmt.Println("Error ketika membuka file:", err)
		return
	}
	defer jsonFile.Close()

	// Baca isi file JSON menjadi byte array
	byteValue, _ := io.ReadAll(jsonFile)

	// Dekode data JSON ke dalam variabel students
	var students Students
	err = json.Unmarshal(byteValue, &students)

	// Periksa kesalahan saat dekode JSON
	if err != nil {
		fmt.Println("Error saat mendekode JSON:", err)
		return
	}

	// Ambil kode mahasiswa dari argumen yang diberikan
	codeMahasiswa := os.Args[1]

	// Cari mahasiswa berdasarkan kode mahasiswa dan cetak informasinya
	searchByCode(students.Participants, codeMahasiswa)
}

// Fungsi untuk mencari mahasiswa berdasarkan kode mahasiswa
func searchByCode(students []Student, code string) {
	for _, student := range students {
		if student.StudentCode == code {
			fmt.Printf("\n")
			fmt.Printf("ID Mahasiswa    : %s\n", student.ID)
			fmt.Printf("Kode Mahasiswa  : %s\n", student.StudentCode)
			fmt.Printf("Nama            : %s\n", student.Name)
			fmt.Printf("Alamat          : %s\n", student.Address)
			fmt.Printf("Pekerjaan       : %s\n", student.Occupation)
			fmt.Printf("Alasan Bergabung: %s\n", student.Reason)
			fmt.Printf("\n")
			return
		}
	}
	fmt.Println("Mahasiswa dengan kode", code, "tidak ditemukan.")
}
