package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

/*
CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  location TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER
	)*/

func (e Event) Save() error {
	// later add to DB
	query := `insert into events (name, description, location, dateTime, user_id)
	values(?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		fmt.Println("err is ", err)
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "select * from events where id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
