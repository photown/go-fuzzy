// Package fuzzy contains the main functionality of go-fuzzy.
package fuzzy

import "../defaults"

// Searchable interface defines the methods needed for a SearchFN object.
type Searchable interface {
  
  // SetPattern is used to give you the pattern and bitap options.
  SetPattern(pattern string, options *BitapOptions)

  // Search performs the Bitap algorithm on the given string.
  // Returns match and score in the form of a SearchResult object.
  Search(text string) *SearchResult
}

// SearchResult struct defines a match and score for a performed check.
type SearchResult struct {
  // IsMatch is true if the given text is a match.
  IsMatch bool

  // Score is an integer representing the score of the text.
  Score int
}

// BitapOptions defines the Bitap algorithm-related options.
type BitapOptions struct {

  // Determines approximately where in the text is the pattern expected to be found.
  Location int

  // At what point the match algorithm gives up.
  // A threshold of 0.0 requires a perfect match (of both letters and location), a threshold of 1.0 would match anything.
  Threshold float32

  //Determines how close the match must be to the fuzzy location (specified by location).
  // An exact letter match which is distance characters away from the fuzzy location would score
  // as a complete mismatch. A distance of 0 requires the match be at the exact location specified,
  // a threshold of 1000 would require a perfect match to be within 800 characters of the location to be found
  // using a threshold of 0.8.
  Distance int

  // The maximum length of the pattern. The longer the pattern, the more intensive the search operation will be.
  // Whenever the pattern exceeds the maxPatternLength, an error will be thrown.
  MaxPatternLength int
}

// Fuzzy defines the objects which initiates the fuzzy search.
// It contains all of the options, needed for the search.
// Initialize it using NewFuzzy() to use the default values.
type Fuzzy struct {

  BitapOptions

  // objects is a list for GoFuzzy to match against.
  objects []interface{}

  // The list of properties to use fuzzy search on. It supports nested properties via dot notation.
  Keys []string

  // Name of the identifier property.
  // If set, instead of returning the objects themselves, it will return the specified identifier of the objects.
  Id string

  // Whether comparisons should be case sensitive.
  CaseSensitive bool

  // Whether to sort the result list by score.
  ShouldSort bool

  // The search function to use. The object must implement `Searchable` interface.
  SearchFn Searchable

  // The method used to access an object's properties.
  // The default implementation handles dot notation nesting (i.e. a.b.c).
  GetFn func(object interface{}, path string) string

  // The function that is used for sorting the result list.
  SortFn func(comparator func(object1 interface{}, object2 interface{}) int)
}

// @param {string} pattern The pattern string to fuzzy search on.
// @return A list of all search matches.
// Searches for all the items whose keys (fuzzy) match the pattern.
func (f *Fuzzy) Search(pattern string) []interface{} {
  return nil
}

// @param list
// @return The newly set list
// Sets a new list for GoFuzzy to match against.
func (f *Fuzzy) Set(list []interface{}) []interface{} {
  return nil
}

// NewFuzzy creates a new Fuzzy object with the default values.
func NewFuzzy() *Fuzzy {
  return &Fuzzy{BitapOptions: BitapOptions{Location: 0, Threshold: 0.6, Distance: 100, MaxPatternLength: 32},
    Keys: nil, Id:"", CaseSensitive: false, ShouldSort: true,
    SearchFn: defaults.BitapSearcher, GetFn: defaults.DefaultGet, SortFn: defaults.DefaultComparator}
}