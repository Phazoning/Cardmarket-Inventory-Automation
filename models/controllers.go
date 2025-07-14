package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (c *Card) GetById(db *pgxpool.Pool, id int) (err error) {
	query := "SELECT id, name, collection, state, value, amount, status FROM cards WHERE id=$1 LIMIT 1"

	row := db.QueryRow(context.Background(), query, id)

	err = row.Scan(&c.Id, &c.Name, &c.Collection, &c.State, &c.Value, &c.Amount, &c.Status)

	return
}

func (c *Card) ChangeStatus(db *pgxpool.Pool, newStatus string) (err error) {
	oldStatus := c.Status

	c.Status = Status(newStatus)

	if err = c.IsValid(); err != nil {
		c.Status = oldStatus
		return
	} else {
		query := "UPDATE cards SET status=$1 WHERE id=$2"

		_, err = db.Exec(context.Background(), query, newStatus, c.Id)
	}

	return
}
