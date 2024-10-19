package main

import "fmt"

type ErrorDetails struct {
	Type         string                   `json:"type"`
	Reason       string                   `json:"reason"`
	ResourceType string                   `json:"resource.type,omitempty"`
	ResourceId   string                   `json:"resource.id,omitempty"`
	Index        string                   `json:"index,omitempty"`
	Phase        string                   `json:"phase,omitempty"`
	Grouped      bool                     `json:"grouped,omitempty"`
	CausedBy     map[string]interface{}   `json:"caused_by,omitempty"`
	RootCause    []*ErrorDetails          `json:"root_cause,omitempty"`
	FailedShards []map[string]interface{} `json:"failed_shards,omitempty"`
}

type BulkResponseItem struct {
	Index         string        `json:"_index,omitempty"`
	Type          string        `json:"_type,omitempty"`
	Id            string        `json:"_id,omitempty"`
	Version       int64         `json:"_version,omitempty"`
	Result        string        `json:"result,omitempty"`
	SeqNo         int64         `json:"_seq_no,omitempty"`
	PrimaryTerm   int64         `json:"_primary_term,omitempty"`
	Status        int           `json:"status,omitempty"`
	ForcedRefresh bool          `json:"forced_refresh,omitempty"`
	Error         *ErrorDetails `json:"error,omitempty"`
}

type BulkResponse struct {
	Took   int                            `json:"took,omitempty"`
	Errors bool                           `json:"errors,omitempty"`
	Items  []map[string]*BulkResponseItem `json:"items,omitempty"`
}

func main() {

	responseItem1 := BulkResponseItem{
		Index:  "product",
		Id:     "1000",
		Result: "updated", // Successful values are created, deleted, and updated.
		Status: 200,
	}

	responseItem2 := BulkResponseItem{
		Index:  "product",
		Id:     "1001",
		Result: "created", // Successful values are created, deleted, and updated.
		Status: 200,
	}

	responseItem3 := BulkResponseItem{
		Index:  "product",
		Id:     "1000",
		Result: "", // Successful values are created, deleted, and updated.
		Status: 404,
	}

	items := make([]map[string]*BulkResponseItem, 3)
	items[0] = make(map[string]*BulkResponseItem)
	items[0]["updated"] = &responseItem1

	items[1] = make(map[string]*BulkResponseItem)
	items[1]["created"] = &responseItem2

	items[2] = make(map[string]*BulkResponseItem)
	items[2]["deleted"] = &responseItem3

	response := BulkResponse{
		Took:   1,
		Errors: false,
		Items:  items,
	}

	fmt.Println(fmt.Sprintf("BulkResponse: %v\n", response))

	for _, item := range response.Items {

		for action, result := range item {
			fmt.Println(fmt.Sprintf("Item: action=%v, result=%v", action, result))
		}
	}

	// Empty array of errors
	var errors []*BulkResponseItem
	fmt.Println(fmt.Sprintf("Errors: %v", errors))
	fmt.Println(fmt.Sprintf("Errors count: %d", len(errors)))

}
