package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
)

// CreateModule saves a new module to the database
func (P *Postgres) CreateModule(ctx context.Context, input model.NewModule) (*string, error) {
	sqlStmt := "INSERT INTO modules(name,synopsis) VALUES( $1,$2) RETURNING id"

	rows := Db.QueryRowContext(ctx, sqlStmt, input.Name, input.Synopsis)

	var i int64 = 0

	err := rows.Scan(&i)
	if err != nil {
		log.Printf("CreateModule: %v", err)
		return nil, err
	}

	id, err := P.createID("modules", i)
	if err != nil {
		log.Printf("CreateModule: %v", err)
		return nil, err
	}

	return id, nil
}

// createID creates a schema ID from a table id
func (P *Postgres) createID(table string, id int64) (*string, error) {
	s := fmt.Sprintf("%s-%d", table, id)

	return &s, nil
}

// getID retrieves a table id from a schema ID
func (P *Postgres) getID(id string, table string) (*int64, error) {
	s := strings.Split(id, "-")
	if len(s) != 2 {
		log.Printf("getID: id not found")
		return nil, fmt.Errorf("id not found")
	}
	itemID := s[1]
	i, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		log.Printf("getID: %v", err)
		return nil, err
	}
	return &i, nil
}
