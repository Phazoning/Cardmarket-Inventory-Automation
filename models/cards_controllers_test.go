package models_test

import (
	"cardmarket_inventory_automation/db"
	"cardmarket_inventory_automation/models"
	"testing"
)

func TestInsertCard(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to db: %s\n", err.Error())
		return
	}

	defer emptyCards(database)

	defer database.Close()

	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

	err = newCard.InsertCard(database)

	if err != nil {
		t.Fatalf("Unable to insert record, %s", err.Error())
		return
	}
}

func TestInsertCardWrongStatus(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to db: %s\n", err.Error())
		return
	}

	defer database.Close()
	defer emptyCards(database)

	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Azbuka"}

	err = newCard.InsertCard(database)

	if err == nil {
		t.Fatal("Wrong status type not recognized")
		return
	}

}

func TestInsertCardWrongState(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to db: %s\n", err.Error())
		return
	}

	defer database.Close()
	defer emptyCards(database)

	newCard := models.Card{Id: 0, Name: "Testcard", State: "AA", Value: 0.1, Amount: 1, Status: "Album"}

	err = newCard.InsertCard(database)

	if err == nil {
		t.Fatal("Wrong state type not recognized")
		return
	}

}

func TestChangeCardStatus(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to database, %s\n", err.Error())
		return
	}

	defer database.Close()
	defer emptyCards(database)

	status := string(models.Deck)

	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

	err = newCard.InsertCard(database)

	if err != nil {
		t.Fatalf("Unable to create card: %s\n", err.Error())
		return
	}

	err = newCard.ChangeStatus(database, status)

	if err != nil {
		t.Fatalf("Unable to perform status change, %s\n", err.Error())
		return
	}
}

func TestChangeCardStatusWrong(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to database, %s\n", err.Error())
		return
	}

	defer emptyCards(database)

	status := "Azbuka"

	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

	err = newCard.InsertCard(database)

	if err != nil {
		t.Fatalf("Unable to create card: %s\n", err.Error())
		return
	}

	err = newCard.ChangeStatus(database, status)

	if err == nil {
		t.Fatal("Status changed to wrong type\n")
		return
	}
}

func TestChangeCardValue(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to database, %s\n", err.Error())
		return
	}

	defer database.Close()
	defer emptyCards(database)

	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

	err = newCard.InsertCard(database)

	if err != nil {
		t.Fatalf("Unable to create card: %s\n", err.Error())
		return
	}

	err = newCard.ChangeValue(database, 0.2)

	if err != nil {
		t.Fatalf("Unable to change card value: %s", err.Error())
		return
	}
}

func TestGetAllCards(t *testing.T) {
	database, err := db.FetchDB()

	if err != nil {
		t.Fatalf("Unable to connect to db: %s\n", err.Error())
		return
	}
	defer database.Close()
	defer emptyCards(database)

	card1 := models.Card{Id: 0, Name: "Testcard1", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}
	card2 := models.Card{Id: 1, Name: "Testcard2", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

	err = card1.InsertCard(database)

	if err != nil {
		t.Fatalf("Unable to insert first card, %s\n", err.Error())
		return
	}

	err = card2.InsertCard(database)

	if err != nil {
		t.Fatalf("Unable to insert second card, %s\n", err.Error())
		return
	}

	_, err = models.GetAllCards(database)

	if err != nil {
		t.Fatalf("Unable to getch all cards: %s\n", err.Error())
	}
}
