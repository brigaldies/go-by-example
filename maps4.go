package main

import (
	"fmt"
	"sync"
)

type CategoryInfo struct {
	CategoryID    int    `json:"id"`
	ParentID      int    `json:"parent_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Level         int    `json:"level,omitempty"`
	ChildrenCount int    `json:"children_count,omitempty"`
	CategoryType  string `json:"category_type,omitempty"`
	Delete        bool   `json:"deleted,omitempty"`
}

var CategoryInfoMapping sync.Map

func main() {
	populateCatInfoMapping()

	countElements := 0
	CategoryInfoMapping.Range(func(key, value any) bool {
		countElements++
		return true
	})

	if countElements == 0 {
		fmt.Printf("empty categoryInfoMapping")
	}

	if len(CategoryInfoMapping) == 0 {
		fmt.Printf("empty categoryInfoMapping (with function len)")
	}

}

func populateCatInfoMapping() {
	catInfo1234 := CategoryInfo{CategoryID: 1234, Name: "Breakfast", Level: 1, ChildrenCount: 3, CategoryType: "Cryptex"}
	catInfo3898 := CategoryInfo{CategoryID: 3898, Name: "Breakfast", Level: 1, ChildrenCount: 3, CategoryType: "Solution Center"}
	catInfo4069 := CategoryInfo{CategoryID: 4069, Name: "Yogurt", Level: 2, ChildrenCount: 1, CategoryType: "Solution Center"}
	catInfo4065 := CategoryInfo{CategoryID: 4065, Name: "Cereal", Level: 2, ChildrenCount: 1, CategoryType: "Solution Center"}
	catInfo4064 := CategoryInfo{CategoryID: 4064, Name: "Morning Beverages", Level: 2, ChildrenCount: 0, CategoryType: "Solution Center"}
	catInfo4463 := CategoryInfo{CategoryID: 4463, Name: "Cereal", Level: 3, ChildrenCount: 0, CategoryType: "Solution Center"}
	catInfo4480 := CategoryInfo{CategoryID: 4480, Name: "Greek & High Protein Yogurt", Level: 3, ChildrenCount: 0, CategoryType: "Solution Center"}
	catInfo3431 := CategoryInfo{CategoryID: 3431, Name: "Dairy", Level: 1, ChildrenCount: 9, CategoryType: "Cryptex"}
	catInfo3434 := CategoryInfo{CategoryID: 3434, Name: "Cheese", Level: 2, ChildrenCount: 1, CategoryType: "Cryptex"}
	catInfo3780 := CategoryInfo{CategoryID: 3780, Name: "String Cheese & Cheese Snacks", Level: 3, ChildrenCount: 0, CategoryType: "Cryptex"}
	catInfo9947 := CategoryInfo{CategoryID: 9947, Name: "Entertainment Gifts for Dad", Level: 1, ChildrenCount: 4, CategoryType: "Merchandising"}
	catInfo9361 := CategoryInfo{CategoryID: 9361, Name: "String Cheese & Cheese Snacks", Level: 3, ChildrenCount: 0, CategoryType: "Cryptex"}
	catInfo2689 := CategoryInfo{CategoryID: 2689, Name: "Kitchen & Table Linens", Level: 2, ChildrenCount: 7, CategoryType: "Cryptex"}
	catInfo3888 := CategoryInfo{CategoryID: 3888, Name: "Household Essentials", Level: 1, ChildrenCount: 13, CategoryType: "Cryptex"}

	CategoryInfoMapping.Store(1234, catInfo1234)
	CategoryInfoMapping.Store(3898, catInfo3898)
	CategoryInfoMapping.Store(4069, catInfo4069)
	CategoryInfoMapping.Store(4065, catInfo4065)
	CategoryInfoMapping.Store(4064, catInfo4064)
	CategoryInfoMapping.Store(4463, catInfo4463)
	CategoryInfoMapping.Store(4480, catInfo4480)
	CategoryInfoMapping.Store(3431, catInfo3431)
	CategoryInfoMapping.Store(3434, catInfo3434)
	CategoryInfoMapping.Store(3780, catInfo3780)
	CategoryInfoMapping.Store(9947, catInfo9947)
	CategoryInfoMapping.Store(9361, catInfo9361)
	CategoryInfoMapping.Store(2689, catInfo2689)
	CategoryInfoMapping.Store(3888, catInfo3888)
}
