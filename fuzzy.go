// Package fuzzy contains the main functionality of go-fuzzy.
package fuzzy

import "strings"
import "sort"
import "github.com/antoan-angelov/fuzzy/internal/defaults"
import "github.com/antoan-angelov/fuzzy/internal/models"

type Results []models.ResultWrapper

// Fuzzy defines the objects which initiates the fuzzy search.
// It contains all of the options, needed for the search.
// Initialize it using NewFuzzy() to use the default values.
type Fuzzy struct {
	models.Options

	// objects is a list for GoFuzzy to match against.
	objects *[]interface{}

	// The list of properties to use fuzzy search on. It supports nested properties via dot notation.
	keys []string

	// Name of the identifier property.
	// If set, instead of returning the objects themselves, it will return the specified identifier of the objects.
	id string

	// Whether comparisons should be case sensitive.
	caseSensitive bool

	// Whether to sort the result list by score.
	shouldSort bool

	// The search function to use. The object must implement `Searchable` interface.
	searchFn models.Searchable

	// The method used to access an object's properties.
	// The default implementation handles dot notation nesting (i.e. a.b.c).
	getFn func(object interface{}, path string) (interface{}, error)

	// The function that is used for sorting the result list.
	sortFn func(object1, object2 models.ResultWrapper) bool
}

// @param {string} pattern The pattern string to fuzzy search on.
// @return A list of all search matches.
// Searches for all the items whose keys (fuzzy) match the pattern.
func (f *Fuzzy) Search(pattern string) ([]interface{}, error) {

	if !f.caseSensitive {
		pattern = strings.ToLower(pattern)
	}

	f.searchFn.SetPattern(pattern, &f.Options)
	result, error := f.retrieveSearchResults()

	if error != nil {
		return nil, error
	}

	if f.shouldSort {
		swap := func(i, j int) { result[i], result[j] = result[j], result[i] }
		less := func(i, j int) bool { return f.sortFn(result[i], result[j]) }
		sort.Sort(&models.FuncSorter{len(result), swap, less})
	}

	return f.constructResult(result)
}

func (f *Fuzzy) retrieveSearchResults() (Results, error) {
	result := make(Results, 0, len(*f.objects))
	for _, element := range *f.objects {
		var maxResult *models.SearchResult

		if f.keys != nil {
			bestMatch, error := f.getBestMatchForKeys(element)
			if error != nil {
				return nil, error
			}

			maxResult = bestMatch
		} else {
			if value, ok := element.(string); ok {
				if !f.caseSensitive {
					value = strings.ToLower(value)
				}

				searchResult := f.searchFn.Search(value)
				if searchResult != nil && searchResult.IsMatch {
					maxResult = searchResult
				}
			} else {
				return nil, &models.InvalidKeyError{}
			}
		}

		if maxResult != nil {
			result = append(result, models.ResultWrapper{maxResult, element})
		}
	}

	return result, nil
}

func (f *Fuzzy) getBestMatchForKeys(element interface{}) (*models.SearchResult, error) {
	var maxResult *models.SearchResult

	for _, key := range f.keys {
		val, _ := f.getFn(element, key)

		if value, ok := val.(string); ok {
			if !f.caseSensitive {
				value = strings.ToLower(value)
			}

			searchResult := f.searchFn.Search(value)
			if searchResult != nil && searchResult.IsMatch {
				if maxResult == nil || maxResult.Score < searchResult.Score {
					maxResult = searchResult
				}
			}
		} else {
			return nil, &models.InvalidKeyError{}
		}
	}

	return maxResult, nil
}

func (f *Fuzzy) constructResult(result Results) ([]interface{}, error) {
	finalResult := make([]interface{}, 0, len(result))

	for _, element := range result {
		if f.id != "" {
			resultValue, error := f.getFn(element.Item, f.id)
			if error == nil {
				finalResult = append(finalResult, resultValue)
			} else {
				return nil, &models.InvalidKeyError{}
			}
		} else {
			finalResult = append(finalResult, element.Item)
		}
	}

	return finalResult, nil
}

// @param list
// @return The newly set list
// Sets a new list for GoFuzzy to match against.
func (f *Fuzzy) Set(list *[]interface{}) *[]interface{} {
	f.objects = list
	return f.objects
}

func (f *Fuzzy) SetKeys(keys []string) {
	f.keys = keys
}

func (f *Fuzzy) SetId(id string) {
	f.id = id
}

func (f *Fuzzy) SetShouldSort(b bool) {
	f.shouldSort = b
}

func (f *Fuzzy) SetCaseSensitive(b bool) {
	f.caseSensitive = b
}

func (f *Fuzzy) SetSearchFn(fn models.Searchable) {
	f.searchFn = fn
}

func (f *Fuzzy) SetGetFn(fn func(object interface{}, path string) (interface{}, error)) {
	f.getFn = fn
}

func (f *Fuzzy) SetSortFn(fn func(object1, object2 models.ResultWrapper) bool) {
	f.sortFn = fn
}

// NewFuzzy creates a new Fuzzy object with the default values.
func NewFuzzy() *Fuzzy {
	return &Fuzzy{Options: models.Options{Threshold: 4},
		keys: nil, id: "", caseSensitive: false, shouldSort: true,
		searchFn: &defaults.Searcher{}, getFn: defaults.DefaultGet, sortFn: defaults.DefaultComparator}
}
