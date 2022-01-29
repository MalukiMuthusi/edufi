package db

import (
	"context"
	"log"

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
