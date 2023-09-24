package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	/*employee, err := getInfo(1001)
	if err != nil {
		// 处理异常
		fmt.Println(err)
		return
	} else {
		fmt.Println(employee)
	}*/

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey! I'm a log message!")

}

func getInfo(id int) (*Employee, error) {
	employee, err := apiCallEmployee(id)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{ID: id, LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
