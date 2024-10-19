package main

// CategoryBoost ...
type CategoryBoost struct {
	CatID int
	Boost float64
}

// TopQueryCategories ...
type TopQueryCategories struct {
	Categories []CategoryBoost
}

// Coverage constants
const (
	Target = 27
)

func main() {
	var topCategories TopQueryCategories
	boostedTypes := make(map[string]struct{})

	storeIDs := []int{
		27,
	}

	searchQuery := "organic strawberry"

	targetQueryToCategories := TopQueryCategories{
		Categories: []CategoryBoost{
			{
				CatID: 1000,
				Boost: 10.0,
			},
		},
	}

	for _, storeID := range storeIDs {
		if _, targetBoosted := boostedTypes["target"]; !targetBoosted && storeID == Target {
			boostedTypes["target"] = struct{}{}
			topCategories.Categories = append(topCategories.Categories, targetQueryToCategories[searchQuery].Categories...)
		}
	}
}
