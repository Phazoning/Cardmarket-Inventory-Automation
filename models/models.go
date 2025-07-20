package models

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type State string

const (
	Mint          State = "MI"
	NearMint      State = "NM"
	Excellent     State = "EX"
	Good          State = "GD"
	Played        State = "PL"
	HeavilyPlayed State = "HP"
	Poor          State = "PO"
	VeryPoor      State = "VP"
)

type Status string

const (
	OnSale Status = "On sale"
	Album  Status = "Album"
	Deck   Status = "Deck"
)

type Card struct {
	Id         int
	Name       string
	Collection string
	State      State
	Value      float32
	Amount     int
	Status     Status
}

func (s State) Ok() (err error) {
	valid := map[State]bool{
		Mint:          true,
		NearMint:      true,
		Excellent:     true,
		Good:          true,
		Played:        true,
		HeavilyPlayed: true,
		Poor:          true,
		VeryPoor:      true,
	}

	if !valid[s] {
		err = fmt.Errorf("unable to parse card status: wrong status - %s", string(s))
	}
	return
}

func (s Status) Ok() (err error) {
	valid := map[Status]bool{
		OnSale: true,
		Album:  true,
		Deck:   true,
	}

	if !valid[s] {
		err = fmt.Errorf("unable to parse card status type: unknown status - %s", string(s))
	}

	return
}

func (c *Card) IsValid() (err error) {
	err = c.State.Ok()

	if err != nil {
		return
	}

	err = c.Status.Ok()

	return
}

func LoadFromCSVRow(header string, row string) (card Card, err error) {
	var splitChar string

	if strings.Contains(header, ";") {
		splitChar = ";"
	} else if strings.Contains(header, ",") {
		splitChar = ","
	} else if strings.Contains(header, ".") {
		splitChar = "."
	} else {
		err = fmt.Errorf("unable to parse csv split character, unknown character used")

		return
	}

	headerParams := strings.Split(header, splitChar)
	rowParams := strings.Split(row, splitChar)

	if len(headerParams) != len(rowParams) {
		err = fmt.Errorf("unable to parse csv row to struct, different number of fields")

		return
	}

	rStruct := reflect.ValueOf(&card).Elem()

	for i, e := range headerParams {
		eAsStructField := e

		for i, r := range eAsStructField {
			eAsStructField = string(unicode.ToUpper(r)) + eAsStructField[i+len(string(r)):]
			break
		}

		field := rStruct.FieldByName(eAsStructField)

		if !field.IsValid() {
			err = fmt.Errorf("unable to parse csv row to struct, unknown header field %s", e)
			return
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(rowParams[i])
		case reflect.Int:
			conv, err := strconv.Atoi(rowParams[i])

			if err != nil {
				return card, err
			}
			field.SetInt(int64(conv))
		case reflect.Float32:
			conv, err := strconv.ParseFloat(rowParams[i], field.Type().Bits())

			if err != nil {
				return card, err
			}

			field.SetFloat(conv)
		default:
			fType := field.Type()

			if fType == reflect.TypeOf(State("")) {
				field.Set(reflect.ValueOf(State(e)))
			} else if fType == reflect.TypeOf(Status("")) {
				field.Set(reflect.ValueOf(Status(e)))
			}
		}

	}

	err = card.IsValid()

	return
}
