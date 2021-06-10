package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	FirstName  string
	LastName   string
	Age        string
	BloodGroup string
}

func main() {
	// Acquire the file handle
	csvLines := readCSV("data.csv")

	// Connect with the DB
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Traverse the read file line by line
	for _, line := range csvLines {
		person := Person{
			FirstName:  line[0],
			LastName:   line[1],
			Age:        line[2],
			BloodGroup: line[3],
		}
		fmt.Println(person.FirstName + " " + person.LastName + " " + person.Age + " " + person.BloodGroup)

		query := "INSERT INTO Persons (first_name, last_name, age, blood_group) VALUES ('" + person.FirstName + "', '" + person.LastName + "', '" + person.Age + "', '" + person.BloodGroup + "')"

		// Insert the read line
		insert, err := db.Query(query)

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
	}

}

// Function to read a CSV file given file path
func readCSV(path string) [][]string {
	csvFile, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	// Read the file
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err.Error())
	}

	return csvLines

}
