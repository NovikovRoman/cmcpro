package types

import (
	"encoding/json"
	"strconv"
)

type Platform struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

func (p *Platform) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var name string
	err = json.Unmarshal(*objMap["name"], &name)
	if err != nil {
		return err
	}
	var symbol string
	err = json.Unmarshal(*objMap["symbol"], &symbol)
	if err != nil {
		return err
	}
	var slug string
	err = json.Unmarshal(*objMap["slug"], &slug)
	if err != nil {
		return err
	}
	var token_address string
	err = json.Unmarshal(*objMap["token_address"], &token_address)
	if err != nil {
		return err
	}

	var idInt int
	err = json.Unmarshal(*objMap["id"], &idInt)
	if err != nil {
		var idString string
		err = json.Unmarshal(*objMap["id"], &idString)
		if err != nil {
			return err
		}
		aI, err := strconv.Atoi(idString)
		if err != nil {
			return err
		}
		p.ID = aI

	} else {
		p.ID = idInt
	}

	p.Name = name
	p.Symbol = symbol
	p.Slug = slug
	p.TokenAddress = token_address
	return nil
}
