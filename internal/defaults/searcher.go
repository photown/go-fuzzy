package defaults

import "math"
import "github.com/antoan-angelov/go-fuzzy/internal/models"

// The object containing the pattern and search function.
// The object must implement `Searchable` interface.
type Searcher struct {
	Pattern string
	Options *models.Options
}

// Sets the pattern and options for the search.
func (b *Searcher) SetPattern(pattern string, options *models.Options) {
	b.Pattern = pattern
	b.Options = options
}

// The search function to use.
func (b *Searcher) Search(text string) *models.SearchResult {

	m := len(b.Pattern)
	n := len(text)

	d := [][]int{}

	for i := 0; i <= m; i++ {
		var a = make([]int, n+1)
		d = append(d, a)
	}

	for i := 1; i <= m; i++ {
		d[i][0] = i
	}

	for j := 1; j <= n; j++ {
		d[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			substitutionCost := 0
			if b.Pattern[i-1] == text[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}

			var min float64 = float64(d[i-1][j] + 1)
			min = float64(math.Min(min, float64(d[i][j-1]+1)))
			min = float64(math.Min(min, float64(d[i-1][j-1]+substitutionCost)))
			d[i][j] = int(min)
		}
	}

	var score uint = uint(d[m][n])

	if score <= b.Options.Threshold {
		return &models.SearchResult{true, score}
	}

	return &models.SearchResult{false, score}
}
