package models_test

import (
	"cardmarket_inventory_automation/models"
	"testing"
)

// newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

func TestCardIsValid(t *testing.T) {
	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Album"}

	err := newCard.IsValid()

	if err != nil {
		t.Fatalf("Failed to validate: %s\n", err.Error())
		return
	}
}

func TestCardIsValidWrongStatus(t *testing.T) {
	newCard := models.Card{Id: 0, Name: "Testcard", State: "NM", Value: 0.1, Amount: 1, Status: "Azbuka"}

	err := newCard.IsValid()

	if err == nil {
		t.Fatal("Card validated in spite of wrong status")
		return
	}
}

func TestCardIsValidWrongState(t *testing.T) {
	newCard := models.Card{Id: 0, Name: "Testcard", State: "NI", Value: 0.1, Amount: 1, Status: "Album"}

	err := newCard.IsValid()

	if err == nil {
		t.Fatal("Card validated in spite of wrong state")
		return
	}
}

func TestCardLoadsFromCSV(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "0;testCard;AAA;MI;0.1;1;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err != nil {
		t.Fatalf("Unable to convert to CSV: %s\n", err.Error())
	}
}

func TestCardLoadsFromCSVMissingFields(t *testing.T) {
	header := "name;collection;state;value;amount;status"
	bodyRow := "testCard;AAA;MI;0.1;1;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err != nil {
		t.Fatalf("Unable to convert to CSV: %s\n", err.Error())
	}
}

func TestCardLoadsFromCSVWrongId(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "aaaa;testCard;AAA;MI;0.1;1;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err == nil {
		t.Fatal("Converted to card in spite of wrong id field")
	}
}

func TestCardLoadsFromCSVWrongState(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "0;testCard;AAA;MN;0.1;1;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err == nil {
		t.Fatal("Converted to card in spite of wrong state field")
	}
}

func TestCardLoadsFromCSVWrongValue(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "0;testCard;AAA;MI;aaaa;1;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err == nil {
		t.Fatal("Converted to card in spite of wrong value field")
	}
}

func TestCardLoadsFromCSVWrongAmount(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "0;testCard;AAA;MI;0.1;aaa;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err == nil {
		t.Fatal("Converted to card in spite of wrong amount field")
	}
}

func TestCardLoadsFromCSVWrongStatus(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "0;testCard;AAA;MI;0.1;1;Azbuka"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err == nil {
		t.Fatal("Converted to card in spite of wrong status field")
	}
}

func TestCardLoadsFromCSVHeaderRowMismatch(t *testing.T) {
	header := "id;name;collection;state;value;amount;status"
	bodyRow := "testCard;AAA;MI;0.1;1;Album"

	_, err := models.LoadFromCSVRow(header, bodyRow)

	if err == nil {
		t.Fatal("Converted to card in spite of mismatched header and row sizes")
	}
}
