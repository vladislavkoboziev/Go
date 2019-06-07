package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

type sql_manipulation interface {
	getAll()
	insertcsv()
	update()
	delete()
}

var people []Person

type Person struct {
	Id               int
	FirstName        string
	LastName         string
	Email            string
	Gender           string
	DateRegistration time.Time
	Loan             float64
}

func parsing_csv() {
	csvFille, _ := os.Open("MOCK_DATA.csv")
	reader := csv.NewReader(csvFille)
	for i := 0; i < 101; i++ {
		line, _ := reader.Read()
		id, _ := strconv.Atoi(line[0])
		loan, _ := strconv.ParseFloat(line[6], 64)
		dataPeople, _ := time.Parse("1/2/2006", line[5])
		people = append(people, Person{
			Id:               id,
			FirstName:        line[1],
			LastName:         line[2],
			Email:            line[3],
			Gender:           line[4],
			DateRegistration: dataPeople,
			Loan:             loan,
		})
	}

}
func connections() *sql.DB {
	connectbd, err := sql.Open("postgres", "postgres://postgres:7154016@localhost/?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = connectbd.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return connectbd
}
func (person Person) insertcsv() {
	connect := connections()
	defer connect.Close()
	rw, _ := connect.Exec("insert into table_name (id,first_name,last_name,email,gender,date_registration,loan) values ($1,$2,$3,$4,$5,$6,$7)",

		person.Id, person.FirstName, person.LastName, person.Email, person.Gender, time.Time(person.DateRegistration), person.Loan)
	fmt.Println(rw)
} ////2
func (person Person) update() {
	connect := connections()
	defer connect.Close()
	rs, err := connect.Exec("update table_name set first_name = $2  where id = $1", "Kyrachkin", 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs.RowsAffected())
} ////3
func (person Person) getAll() {
	connect := connections()
	defer connect.Close()
	rows, err := connect.Query("select * from table_name")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		pp := Person{}
		err := rows.Scan(&pp.Id, &pp.FirstName, &pp.LastName, &pp.Email, &pp.Gender, &pp.DateRegistration, &pp.Loan)
		if err != nil {
			fmt.Println(err)
			continue
		}
		people = append(people, pp)
	}
	for _, p := range people {
		fmt.Println(p.Id, p.FirstName, p.LastName, p.Email, p.Gender, p.DateRegistration, p.Loan)
	}
} ////4
func (person Person) delete() {
	connect := connections()
	defer connect.Close()
	rs, err := connect.Exec("delete from debtors where id = $1", person.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs.RowsAffected())

}
func main() {
	//parsing_csv()
	//for _, value:= range people{value.insertcsv()}
	//people := Person{}
	//people.getAll()
}
