package fuzzy

import (
      "testing"
      "fmt"
)

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

func generateLaptops() []Laptop {
  processor1 := &CPU{Manufacturer: "Intel", Series: "i7", Cores: 4, Frequency: 4}
  laptop1 := &Laptop{Manufacturer: "Acer", HasFan: true, Processor: processor}

  processor2 := &CPU{Manufacturer: "AMD", Series: "Athlon", Cores: 4, Frequency: 4}
  laptop2 := &Laptop{Manufacturer: "Lenovo", HasFan: true, Processor: processor}

  return []Laptop { laptop1, laptop2 }
}

func TestKeys1(t *testing.T) {
  laptops := generateLaptops()

  fuzzyDemo := fuzzy.NewFuzzy()
  fuzzyDemo.Set(laptops)
  fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}
  
  results := fuzzyDemo.Search("Inet")
  if len(result) != 1 || results[0] != laptops[0] {
    t.Error("Expected the Acer laptop with Intel processor.")
  }
}

func TestKeys2(t *testing.T) {
  laptops := generateLaptops()

  fuzzyDemo := fuzzy.NewFuzzy()
  fuzzyDemo.Set(laptops)
  fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}
  
  results := fuzzyDemo.Search("Elnovo")
  if len(result) != 1 || results[0] != laptops[0] {
    t.Error("Expected the Lenovo laptop with AMD processor.")
  }
}

func TestId(t *testing.T) {
  laptops := generateLaptops()

  fuzzyDemo := fuzzy.NewFuzzy()
  fuzzyDemo.Set(laptops)
  fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}
  fuzzyDemo.Id = "Prosessor.Series"
  
  results := fuzzyDemo.Search("Inet")
  if len(result) != 1 || result[0] != "i7" {
    t.Error(fmt.Sprintf("Expected i7, got %s.", result[0]))
  }
}

func TestCaseSensitive(t *testing.T) {
  laptops := generateLaptops()

  fuzzyDemo := fuzzy.NewFuzzy()
  fuzzyDemo.Set(laptops)
  fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}
  fuzzyDemo.CaseSensitive = true
  
  results := fuzzyDemo.Search("eLnovo")
  if len(result) != 1 || result[0] != laptops[1] {
    t.Error("Expected Lenovo laptop.")
  }
}

func TestMaxThreshold(t *testing.T) {
  laptops := generateLaptops()

  fuzzyDemo := fuzzy.NewFuzzy()
  fuzzyDemo.Set(laptops)
  fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}
  fuzzyDemo.Threshold = 1.0
  
  results := fuzzyDemo.Search("Elnovo")
  if len(result) != len(laptops) {
    t.Error(fmt.Sprintf("Expected %d elements, got %d instead.", len(laptops), len(result)))
  }
}

func TestMinThreshold(t *testing.T) {
  laptops := generateLaptops()

  fuzzyDemo := fuzzy.NewFuzzy()
  fuzzyDemo.Set(laptops)
  fuzzyDemo.Кeys = []string{"Manufacturer", "Processor.Manufacturer"}
  fuzzyDemo.Threshold = 0.0
  
  results := fuzzyDemo.Search("Elnovo")
  if len(result) != 0 {
    t.Error(fmt.Sprintf("Expected 0 elements, got %d instead.", len(result)))
  }
}