// Subair: this is  controller not service
package amazon

import (
	"encoding/json"

	"../config"
)

// AmazonItem ...
// Subair:  Each API will have different request and reposnse struct that will come inside DAO
// The struct you initiallise in interface file should not be this.
// Think of a struct which have common to all the apis.
// Add market place struct here and add a initialising fn here that should be called from interface.
type AmazonItem struct {
	ItemID      string  `json:"itemid"`
	ItemName    string  `json:"item"`
	Price       float32 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Rating      string  `json:"rating"`
	MarketPlace string  `json:"platform"`
}

// GetItem returns a specific item if available
func (a *AmazonItem) GetItem(itemname string) (string, error) {
	row := config.DB.QueryRow(`select * from items join marketplace using(itemid) where (marketplace = 'amazon' and name = '` + itemname + `');`)

	var item AmazonItem
	err := row.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.MarketPlace)
	if err != nil {
		return "", err
	}

	itemJSON, _ := json.Marshal(item)
	return string(itemJSON), nil
}

func CreateItem() *AmazonItem {
	return &AmazonItem{MarketPlace: "amazon"}
}
