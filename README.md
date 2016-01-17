# go-fuzzy
Go-fuzzy is a Golang fuzzy search implementation that for given keys returns objects with the closest values to an input.

A shameless port of [Fuse](https://github.com/krisk/Fuse).

## Installation
```
go get github.com/antoan-angelov/go-fuzzy
```
## Example

```go
type Laptop struct {
  Manufacturer string
  HasFan bool
  Processor CPU
}

type CPU struct {
  Manufacturer string
  Series string
  Cores int
  Frequency int
}

processor1 := &CPU{Manufacturer: "Intel", Series: "i7", Cores: 4, Frequency: 4}
laptop1 := &Laptop{Manufacturer: "Acer", HasFan: true, Processor: processor}

processor2 := &CPU{Manufacturer: "AMD", Series: "Athlon", Cores: 4, Frequency: 4}
laptop2 := &Laptop{Manufacturer: "Lenovo", HasFan: true, Processor: processor}

laptops := []Laptop { laptop1, laptop2 }

fuzzyDemo := fuzzy.NewFuzzy()
fuzzyDemo.Set(laptops)
fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}

results := fuzzyDemo.Search("Inet") // returns [laptop1]
```

## Options
**Кeys** (_type:_ `[]string`, _default:_ `nil`)  
The list of properties to use fuzzy search on. It supports nested properties via dot notation.

---

**Id** (_type:_ `string`, _default:_ `""`)  
Name of the identifier property. If set, instead of returning the objects themselves, it will return the specified identifier of the objects.

---

**CaseSensitive** (_type:_ `bool`, _default:_ `false`)  
Whether comparisons should be case sensitive.

---

**ShouldSort** (_type:_ `bool`, _default:_ `true`)  
Whether to sort the result list by score.

---

**SearchFn** (_type:_ `Searchable`, _default:_ `defaults.BitapSearcher`)  
The search function to use. The object must implement `Searchable` interface:
```go
type Searchable interface {
    SetPattern(pattern string, options *fuzzy.BitapOptions)
    Search(text string) *fuzzy.SearchResult
}
```

---

**GetFn** (_type:_ `func(object interface{}, path string) interface{}`, _default:_ `defaults.DefaultGet`)  
The method used to access an object's properties. The default implementation handles dot notation nesting (i.e. a.b.c).

---

**SortFn** (_type:_ `func(object1 interface{}, object2 interface{}) int`, _default:_ `defaults.DefaultComparator`)  
The function that is used for sorting the result list.


### Bitap specific options
**Location** (_type:_ `int`, _default:_ `0`)  
Determines approximately where in the text is the pattern expected to be found.

---

**Threshold** (_type:_ `float32`, _default:_ `0.6`)  
At what point the match algorithm gives up. A threshold of 0.0 requires a perfect match (of both letters and location), a threshold of 1.0 would match anything.

---

**Distance** (_type:_ `int`, _default:_ `100`)  
Determines how close the match must be to the fuzzy location (specified by location). An exact letter match which is distance characters away from the fuzzy location would score as a complete mismatch. A distance of 0 requires the match be at the exact location specified, a threshold of 1000 would require a perfect match to be within 800 characters of the location to be found using a threshold of 0.8.

---

**MaxPatternLength** (_type:_ `int`, _default:_ `32`)  
The maximum length of the pattern. The longer the pattern, the more intensive the search operation will be. Whenever the pattern exceeds the maxPatternLength, an error will be thrown.

## Methods

### func Search(pattern string) []interface{}

@param {string} pattern The pattern string to fuzzy search on.
@return A list of all search matches.

Searches for all the items whose keys (fuzzy) match the pattern.

### func Set(list []interface{}) []interface{}

@param list
@return The newly set list

Sets a new list for GoFuzzy to match against.

## License
```
The MIT License (MIT)

Copyright (c) 2016 Antoan Angelov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
