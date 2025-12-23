package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Get_A_Grip26"
	dbname   = "skillMatch"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}
	fmt.Println("Successfully connected to the database!")

	type Employer struct {
		ID          int
		CompanyName string
		ContactInfo string
		Description string
		CreatedAt   time.Time
	}

	fmt.Println("\nEmployers: ")
	sqlStatement := `SELECT * FROM employers LIMIT $1;`
	rows, err := db.Query(sqlStatement, 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var emp Employer
		err = rows.Scan(&emp.ID, &emp.CompanyName, &emp.ContactInfo, &emp.Description, &emp.CreatedAt)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %s %q %s\n", emp.ID, emp.CompanyName, emp.ContactInfo, emp.Description, emp.CreatedAt.Format("02-01-2006 15:04:05"))
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// sqlStatement = `UPDATE candidates SET password_hash = $2 WHERE id = $1`
	// _, err = db.Exec(sqlStatement, 1, "Get_A_Grip26")
	// if err != nil {
	// 	panic(err)
	// }

	type Candidates struct {
		ID				int
		Name 			string
		Email 			string
		PasswordHash 	string
		ContactInfo 	string
		CreatedAt   	time.Time
	}

	fmt.Println("\nCandidates: ")
	sqlStatement = `SELECT * FROM candidates ORDER BY id LIMIT $1;`
	rows, err = db.Query(sqlStatement, 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var cand Candidates
		err = rows.Scan(&cand.ID, &cand.Name, &cand.Email, &cand.PasswordHash, &cand.ContactInfo, &cand.CreatedAt)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %s #%v# %s %s\n", cand.ID, cand.Name, cand.Email, cand.PasswordHash, cand.ContactInfo, cand.CreatedAt.Format("02-01-2006 15:04:05"))
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
