package main

import (
	"fmt"
	"time"
)

type ProductSetting struct {
	StoreID               int
	ProductID             int       `db:"product_id"`
	MetroID               int       `db:"metro_id"`
	StoreLocationID       int       `db:"store_location_id"`
	SettingID             int       `db:"setting_id"` // dsom-publisher requires setting key for indexing individual storelocation document.
	UpdatedAt             time.Time `db:"updated_at"`
	AvailabilityUpdatedAt time.Time `db:"availability_updated_at"`
	AvailabilityUpdated   bool      `db:"availability_updated"`
	PrismLastUpdated      time.Time `json:"prism_last_updated"` // obtained from prism - not from DB
}

func main() {
	valuesToCheck := make([]ProductSetting, -1)

	fmt.Println(fmt.Sprintf("valuesToCheck=%v, is nil=%t", valuesToCheck, valuesToCheck == nil))
}
