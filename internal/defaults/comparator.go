package defaults

import "github.com/antoan-angelov/go-fuzzy/internal/models"

// The function that is used for sorting the result list.
// Should return true if object1 is LESS than object2, false otherwise.
func DefaultComparator(object1, object2 models.ResultWrapper) bool {
	return object1.Result.Score < object2.Result.Score
}
