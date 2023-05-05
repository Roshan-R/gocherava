package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database instance
var db *sql.DB

type Workflow struct {
	ID          string `json:"id"`
	User        string `json:"user"`
	Data        string `json:"data"`
	Selector    string `json:"selector"`
	Cron        string `json:"cron"`
	LastUpdated int64  `json:"lastupdated"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

type Workflows struct {
	Workflows []Workflow `json:"workflows"`
}

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		// panic("Error loading .env file")
	}

	dbname := os.Getenv("PGDATABASE")
	host := os.Getenv("PGHOST")
	password := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")

	fmt.Println(dbname)

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func getWorkflowByID(id string) (workflow *Workflow, error error) {

	query := `SELECT * FROM Workflows where id = $1`
	row, err := db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()
	w := Workflow{}
	row.Next()

	if err := row.Scan(&w.ID, &w.User, &w.Data, &w.Selector, &w.Cron, &w.LastUpdated, &w.URL, &w.Name, &w.Email); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &w, nil
}

func getWorkflowsByUserID(uid string) (workflows *Workflows, error error) {

	query := `SELECT * FROM Workflows where "user" = $1`
	rows, err := db.Query(query, uid)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	result := new(Workflows)

	for rows.Next() {
		w := Workflow{}
		if err := rows.Scan(&w.ID, &w.User, &w.Data, &w.Selector, &w.Cron, &w.LastUpdated, &w.URL, &w.Name, &w.Email); err != nil {
			fmt.Println(err)
			return nil, err
		}
		result.Workflows = append(result.Workflows, w)
	}
	return result, nil
}

func getAllWorkflows() (workflows *Workflows, error error) {
	query := `SELECT * FROM Workflows `
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	result := new(Workflows)

	for rows.Next() {
		w := Workflow{}
		if err := rows.Scan(&w.ID, &w.User, &w.Data, &w.Selector, &w.Cron, &w.LastUpdated, &w.URL, &w.Name, &w.Email); err != nil {
			fmt.Println(err)
			return nil, err
		}
		result.Workflows = append(result.Workflows, w)
	}
	return result, nil
}

func updateWorkflow(w Workflow, new_data string) error {
	query := `UPDATE Workflows set data = $1, lastupdated = $2 where id = $3`
	_, err := db.Exec(query, new_data, fmt.Sprint(time.Now().UnixMilli()), w.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func createNewWorkflow(w Workflow) error {

	query := `INSERT INTO workflows (id, "user", data, selector, cron, lastupdated, url, name, email) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := db.Exec(query, w.ID, w.User, w.Data, w.Selector, w.Cron, w.LastUpdated, w.URL, w.Name, w.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
