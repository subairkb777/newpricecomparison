package services

import (
	"encoding/json"

	"./amazon"
	"./flipkart"
)

// Subair: change the name to MarketPlaceClient
type ItemInterface interface {
	// suabir: return a struct instead of string.
	GetItem(string) (string, error)
}

// Subair: no definitions here
// Subair: call initialise fn amazon, flipkart etc. object specific struct here. (no definition) and store the corresponding objects in a hashMap
// key: market place, value: initialised struct
func GetItem(itemname string, marketplace string) (ItemInterface, error) {
	var i ItemInterface
	switch marketplace {
	case "amazon":
		i = amazon.CreateItem()
	case "flipkart":
		i = flipkart.CreateItem()
	}
	// a1 := amazon.CreateItem()
	// f1 := flipkart.CreateItem()
	dataString, err := i.GetItem(itemname)
	if err != nil {
		return nil, err
	}

	// var data *ItemInterface

	err = json.Unmarshal([]byte(dataString), &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
