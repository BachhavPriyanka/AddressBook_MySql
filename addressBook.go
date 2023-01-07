package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Contact struct {
	id          int
	FirstName   string
	LastName    string
	Address     string
	City        string
	State       string
	PhoneNumber string
	Email       string
}

func main() {
	var err error
	//"UserName:Password@tcp(portNumber)/databaseName"
	db, err = sql.Open("mysql", "root:root@tcp(localhost:6603)/addressBook")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")
	fmt.Println("**************** Wel-Come To Address Book ****************")
	Operation()
}

func Operation() {
	var option int
	fmt.Println("\n-------------- Select which operation do you wanna perform --------------\n1.Add Contact\n2.Update Contact \n3.Display Contacts\n4.Details of Persons belongs to same city and state\n5.Count of Persons belongs to same city and state\n6.Exit")
	fmt.Scanln(&option)
	switch option {
	case 1:
		addContact()
	case 2:
		var updateID int
		fmt.Println("Enter a ID which you want to Update :")
		fmt.Scanln(&updateID)
		updateContact(updateID)
	case 3:
		fmt.Println(DatabaseReader())
	case 4:
		fmt.Println(findCityState())
	case 5:
		findCityStateCount()
	case 6:
		return
	}
	Operation()
}

func updateContact(updateID int) {
	var person Contact
	var option int
	fmt.Println("Which Field do you want to Update \n1.FirstName\n.LastName\n3.Address\n4.City\n5.State\n6.Phone Number\n7.Email ID")
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Println("Enter First Name: ")
		fmt.Scanln(&person.FirstName)
		result, err := db.Exec("UPDATE Contact SET FirstName = ? where id = ?", person.FirstName, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)

	case 2:
		fmt.Println("Enter Last Name: ")
		fmt.Scanln(&person.LastName)
		result, err := db.Exec("Update Contact SET LastName = ? where id = ?", person.LastName, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)
	case 3:
		fmt.Println("Enter Address: ")
		fmt.Scanln(&person.Address)
		result, err := db.Exec("Update Contact SET Address = ? where id = ?", person.Address, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)
	case 4:
		fmt.Println("Enter City: ")
		fmt.Scanln(&person.City)
		result, err := db.Exec("Update Contact SET City = ? where id = ?", person.City, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)
	case 5:
		fmt.Println("Enter State: ")
		fmt.Scanln(&person.State)
		result, err := db.Exec("Update Contact SET State = ? where id = ?", person.State, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)
	case 6:
		fmt.Println("Enter Phone Number: ")
		fmt.Scanln(&person.PhoneNumber)
		result, err := db.Exec("Update Contact SET PhoneNumber = ? where id = ?", person.PhoneNumber, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)
	case 7:
		fmt.Println("Enter Email-Id: ")
		fmt.Scanln(&person.Email)
		result, err := db.Exec("Update Contact SET Email = ? where id = ?", person.Email, updateID)
		if err != nil {
			panic(err.Error())
		}
		id, _ := result.RowsAffected()
		fmt.Println(id)
	}

}

func findCityStateCount() {
	var cityStateName string

	fmt.Println("To Find the Count of Persons From Same City and State\nEnter City or State:")
	fmt.Scanln(&cityStateName)
	rows, err := db.Query("SELECT Count(*) FROM Contact WHERE City = ? OR State = ?", cityStateName, cityStateName)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var citystate string
		if err := rows.Scan(&citystate); err != nil {
			panic(err.Error())
		}
		fmt.Printf("Person Present : %s\n", citystate)
	}
}
func findCityState() ([]Contact, error) {
	var contacts []Contact
	var newcityStateName string
	fmt.Println("To Find Persons From Same City and state \nEnter City or State :")
	fmt.Scanln(&newcityStateName)
	rows, err := db.Query("SELECT * From Contact WHERE City = ? OR State = ?", newcityStateName, newcityStateName)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var person Contact
		if err := rows.Scan(&person.id, &person.FirstName, &person.LastName, &person.Address, &person.City, &person.State, &person.PhoneNumber, &person.Email); err != nil {
			return nil,
				fmt.Errorf("error in query : %v", err)
		}
		contacts = append(contacts, person)
	}
	return contacts, nil
}
func DatabaseReader() ([]Contact, error) {
	var contacts []Contact

	rows, err := db.Query("SELECT * FROM Contact;")
	if err != nil {
		return nil, fmt.Errorf("error in query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var person Contact

		if err := rows.Scan(&person.id, &person.FirstName, &person.LastName, &person.Address, &person.City, &person.State, &person.PhoneNumber, &person.Email); err != nil {
			return nil,
				fmt.Errorf("error in query : %v", err)
		}

		contacts = append(contacts, person)

	}
	return contacts, nil
}

func addContact() (int64, error) {
	var person Contact

	fmt.Println("Enter First Name: ")
	fmt.Scanln(&person.FirstName)

	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&person.LastName)

	fmt.Println("Enter Address: ")
	fmt.Scanln(&person.Address)

	fmt.Println("Enter City: ")
	fmt.Scanln(&person.City)

	fmt.Println("Enter State: ")
	fmt.Scanln(&person.State)

	fmt.Println("Enter Phone Number: ")
	fmt.Scanln(&person.PhoneNumber)

	fmt.Println("Enter Email-Id: ")
	fmt.Scanln(&person.Email)

	result, err := db.Exec("INSERT INTO Contact(FirstName, LastName, Address, City, State, PhoneNumber, Email) VALUES (?, ?, ?, ?, ?, ?, ?)", person.FirstName, person.LastName, person.Address, person.City, person.State, person.PhoneNumber, person.Email)
	if err != nil {
		return 0, fmt.Errorf("add person: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add person: %v", err)
	}
	return id, nil

}
