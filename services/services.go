package services

import (
	"encoding/json"

	"./amazon"
	"./flipkart"
)

type ItemInterface interface {
	GetItem(string) (string, error)
}

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
