package defaults

import "../fuzzy"

// The object containing the pattern and search function.
// The object must implement `Searchable` interface.
type BitapSearcher struct {
  Pattern string
  Options []string
}

// Sets the pattern and options for the search.
func (b *BitapSearcher) SetPattern(pattern string, options *fuzzy.BitapOptions) {

}

// The search function to use.
func (b *BitapSearcher) Search(text string) *fuzzy.SearchResult {
  return nil
}