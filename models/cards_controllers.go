package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (c *Card) ChangeValue(db *pgxpool.Pool, newValue float32) (err error) {
	query := "UPDATE cards SET value=$1 WHERE Id=$2"

	_, err = db.Exec(context.Background(), query, newValue, c.Id)

	if err != nil {
		return
	}

	c.Value = newValue

	return
}

func (c *Card) InsertCard(db *pgxpool.Pool) (err error) {
	if err = c.IsValid(); err != nil {
		return
	}

	query := "INSERT INTO cards (id, name, collection, state, value, amount, status) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err = db.Exec(context.Background(), query, c.Id, c.Name, c.Collection, c.State, c.Value, c.Amount, c.Status)

	return
}

func GetAllCards(db *pgxpool.Pool) (cards []Card, err error) {

	query := "SELECT id, name, collection, state, value, amount, status FROM cards"

	rows, err := db.Query(context.Background(), query)

	if err != nil {
		return
	}

	for rows.Next() {
		var card Card

		err = rows.Scan(&card.Id, &card.Name, &card.Collection, &card.State, &card.Value, &card.Amount, &card.Status)

		if err != nil {
			return
		}

		cards = append(cards, card)
	}

	return
}
