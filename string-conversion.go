package main

import (
	"fmt"
	"strconv"
)

func GetProductIDs(hits []map[string]interface{}) []int {

	productIDs := []int{}
	for _, elem := range hits {
		intID := GetProductID(elem)
		if intID != 0 {
			productIDs = append(productIDs, intID)
		}
	}
	return productIDs
}

func GetProductID(hit map[string]interface{}) int {

	id, _ := hit["product_id"]
	stringID, _ := id.(string)
	if stringID == "" {
		stringID = fmt.Sprintf("%.0f", id)
	}
	intID, _ := strconv.Atoi(stringID)
	return intID
}

func main() {
	purchaseHistoryHits := []map[string]interface{}{
		{
			"product_id":              1430580,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-03-29T23:22:18.037Z",
			"last_order_created_date": "2023-03-27T22:06:25.000Z",
			"version":                 19,
			"purchase_counter_total":  17,
		},
		{
			"product_id":              1375922,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-03-29T23:22:18.037Z",
			"last_order_created_date": "2023-03-27T22:06:25.000Z",
			"version":                 8,
			"purchase_counter_total":  8,
		},
		{
			"product_id":              1509381,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-02-16T20:00:51.633Z",
			"last_order_created_date": "2021-04-27T18:36:39.000Z",
			"version":                 3,
			"purchase_counter_total":  8,
		},
		{
			"product_id":              1462245,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-03-29T23:22:18.037Z",
			"last_order_created_date": "2023-03-27T22:06:25.000Z",
			"version":                 9,
			"purchase_counter_total":  7,
		},
		{
			"product_id":              1442165,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-03-29T23:22:18.037Z",
			"last_order_created_date": "2023-03-27T22:06:25.000Z",
			"version":                 8,
			"purchase_counter_total":  7,
		},
		{
			"product_id":              1375933,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-03-29T23:22:18.037Z",
			"last_order_created_date": "2023-03-27T22:06:25.000Z",
			"version":                 9,
			"purchase_counter_total":  7,
		},
		{
			"product_id":              1558343,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2020-08-29T00:06:50.627Z",
			"last_order_created_date": "2020-02-25T20:14:57.000Z",
			"version":                 9,
			"purchase_counter_total":  7,
		},
		{
			"product_id":              1365405,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2023-03-29T23:22:18.037Z",
			"last_order_created_date": "2023-03-27T22:06:25.000Z",
			"version":                 8,
			"purchase_counter_total":  5,
		},
		{
			"product_id":              3542924,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2020-08-29T00:07:26.182Z",
			"last_order_created_date": "2020-02-24T02:00:09.000Z",
			"version":                 7,
			"purchase_counter_total":  5,
		},
		{
			"product_id":              1430579,
			"store_id":                27,
			"user_id":                 1694927,
			"hidden":                  false,
			"es_last_updated_date":    "2020-08-29T00:06:57.300Z",
			"last_order_created_date": "2020-02-21T22:11:35.000Z",
			"version":                 7,
			"purchase_counter_total":  5,
		},
	}

	fmt.Println(GetProductIDs(purchaseHistoryHits))
}
