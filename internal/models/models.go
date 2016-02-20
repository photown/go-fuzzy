// Package fuzzy contains the main functionality of go-fuzzy.
package models

// Searchable interface defines the methods needed for a SearchFN object.
type Searchable interface {

	// SetPattern is used to give you the pattern and Levenshtein options.
	SetPattern(pattern string, options *Options)

	// Search performs the Levenshtein algorithm on the given string.
	// Returns match and score in the form of a SearchResult object.
	Search(text string) *SearchResult
}

// SearchResult struct defines a match and score for a performed check.
type SearchResult struct {
	// IsMatch is true if the given text is a match.
	IsMatch bool

	// Score is an integer representing the score of the text.
	Score uint
}

type FuncSorter struct {
	FLen  int
	FSwap func(int, int)
	FLess func(int, int) bool
}

func (f *FuncSorter) Len() int           { return f.FLen }
func (f *FuncSorter) Less(i, j int) bool { return f.FLess(i, j) }
func (f *FuncSorter) Swap(i, j int)      { f.FSwap(i, j) }

type ResultWrapper struct {
	Result *SearchResult
	Item   interface{}
}

// LevenshteinOptions defines the Levenshtein algorithm-related options.
type Options struct {
	// At what point the match algorithm gives up.
	// A threshold of 0 requires a perfect match (of both letters and location). The higher the number, the less the required similarity between words.
	Threshold uint
}

func (o *Options) SetThreshold(threshold uint) {
	o.Threshold = threshold
}

type InvalidKeyError struct{}

func (e *InvalidKeyError) Error() string {
	return "Provided key is either nil or does not point to a string value."
}
