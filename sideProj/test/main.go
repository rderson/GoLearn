package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Get_A_Grip26"
	dbname   = "testGo"
)

func main() {
	fmt.Println(time.Now().Format("02-01-2006"))

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

	// sqlStatement := `INSERT INTO users (age, first_name, last_name, email)
	// VALUES ($1, $2, $3, $4) RETURNING id`
	// id := 0
	// err = db.QueryRow(sqlStatement, 25, "Mikail", "Torshkhoev", "mikail.torshkhoev@gmail.com").Scan(&id)
	// if err != nil {	
	// 	panic(err)
	// }

	// fmt.Printf("The newly added record's id is: %d\n", id)

	// sqlStatement = `UPDATE users 
	// SET age = $2 
	// WHERE id = $1;`
	// res, err := db.Exec(sqlStatement, id, 13)
	// if err != nil {
	// 	panic(err)
	// }

	// rowsAffected, err := res.RowsAffected()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Number of rows affected: %d\n", rowsAffected)

	// fmt.Printf("Record with id %d updated successfully!\n", id)

	// sqlStatement = `
	// DELETE FROM users
	// WHERE id = $1;`
	// _, err = db.Exec(sqlStatement, id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Record deleted successfully!")

	sqlStatement := `SELECT id, email FROM users WHERE id=$1;`
	var email string
	var id int

	row := db.QueryRow(sqlStatement, 3)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}

	type User struct {
		ID        int
		Age 	 int
		FirstName string
		LastName  string
		Email     string
	}

	sqlStatement = `SELECT * FROM users WHERE id = $1;`
	var user User
	row = db.QueryRow(sqlStatement, 9)
	err = row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}	

	fmt.Println("Listing users...")

	sqlStatement = `SELECT * FROM users LIMIT $1;`
	rows, err := db.Query(sqlStatement, 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

// Код использованный для вставки записей в базу данных SkillMacth

// testCandidates := []struct {
//     Name         string
//     Email        string
//     PasswordHash string
//     ContactInfo  string
// }{
//     {"Ivan Petrov", "ivan.petrov@example.com", "hash1", "+7 999 123-45-67"},
//     {"Anna Smirnova", "anna.smirnova@example.com", "hash2", "+7 999 234-56-78"},
//     {"John Doe", "john.doe@example.com", "hash3", "+1 555-123-4567"},
//     {"Maria Ivanova", "maria.ivanova@example.com", "hash4", "+7 999 345-67-89"},
//     {"Alex Johnson", "alex.johnson@example.com", "hash5", "+1 555-234-5678"},
// }

// sqlStatement := `INSERT INTO candidates (name, email, password_hash, contact_info, created_at) VALUES ($1, $2, $3, $4, NOW())`
// for _, cand := range testCandidates {
//     _, err = db.Exec(sqlStatement, cand.Name, cand.Email, cand.PasswordHash, cand.ContactInfo)
//     if err != nil {
//         fmt.Println("Error inserting candidate:", err)
//     }
// }

// testEmployers := []struct {
//     CompanyName string
//     ContactInfo string
//     Description string
// }{
//     {"Google", "+1 650-253-0000", "Innovative technology company specializing in search and cloud."},
//     {"Yandex", "+7 495 739-7000", "Russian multinational IT company, search and AI leader."},
//     {"Microsoft", "+1 425-882-8080", "Leading global provider of software, services, and solutions."},
// }

// sqlStatement := `INSERT INTO employers (id, company_name, contact_info, description, created_at) VALUES ($1, $2, $3, $4, NOW())`
// id := 3
// for _, emp := range testEmployers {
//     _, err = db.Exec(sqlStatement, id, emp.CompanyName, emp.ContactInfo, emp.Description)
//     if err != nil {
//         fmt.Println("Error inserting employer:", err)
//     }
// 	id += 1
// }

// sqlStatement := `INSERT INTO employers (id, company_name, contact_info, description, created_at) VALUES ($1, $2, $3, $4, NOW())`
// _, err = db.Exec(sqlStatement, 2, "SAP", "+49 (0)6227 / 7-47474", "As a global leader in enterprise applications and business AI, SAP stands at the nexus of business and technology.")
// if err != nil {
// 	panic(err)
// }