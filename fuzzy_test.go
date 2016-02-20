package fuzzy

import (
	"fmt"
	"testing"
)

type Laptop struct {
	Manufacturer string
	HasFan       bool
	Processor    *CPU
}

type CPU struct {
	Manufacturer string
	Series       string
	Cores        int
	Frequency    int
}

func generateLaptops() *[]interface{} {
	processor1 := &CPU{Manufacturer: "Intel", Series: "i7", Cores: 4, Frequency: 4}
	laptop1 := Laptop{Manufacturer: "Acer", HasFan: true, Processor: processor1}

	processor2 := &CPU{Manufacturer: "AMD", Series: "Athlon", Cores: 4, Frequency: 4}
	laptop2 := Laptop{Manufacturer: "Lenovo", HasFan: true, Processor: processor2}

	return &[]interface{}{laptop1, laptop2}
}

func TestKeys1(t *testing.T) {
	laptops := generateLaptops()

	fuzzyDemo := NewFuzzy()
	fuzzyDemo.Set(laptops)
	fuzzyDemo.SetKeys([]string{"Manufacturer", "Processor.Manufacturer"})

	results, _ := fuzzyDemo.Search("Inetl")
	if len(results) != 1 || results[0] != (*laptops)[0] {
		t.Error("Expected the Acer laptop with Intel processor.")
	}
}

func TestKeys2(t *testing.T) {
	laptops := generateLaptops()

	fuzzyDemo := NewFuzzy()
	fuzzyDemo.Set(laptops)
	fuzzyDemo.SetKeys([]string{"Manufacturer", "Processor.Manufacturer"})
	fuzzyDemo.SetCaseSensitive(false)
	fuzzyDemo.SetShouldSort(false)

	results, _ := fuzzyDemo.Search("Elnovo")
	if len(results) != 1 || results[0] != (*laptops)[1] {
		t.Error("Expected the Lenovo laptop with AMD processor.")
	}
}

func TestId(t *testing.T) {
	laptops := generateLaptops()

	fuzzyDemo := NewFuzzy()
	fuzzyDemo.Set(laptops)
	fuzzyDemo.SetKeys([]string{"Manufacturer", "Processor.Manufacturer"})
	fuzzyDemo.SetId("Processor.Series")

	results, _ := fuzzyDemo.Search("Inetl")
	if len(results) != 1 || results[0] != "i7" {
		t.Error(fmt.Sprintf("Expected i7, got %s and %d results.", results[0], len(results)))
	}
}

func TestCaseSensitive(t *testing.T) {
	laptops := generateLaptops()

	fuzzyDemo := NewFuzzy()
	fuzzyDemo.Set(laptops)
	fuzzyDemo.SetKeys([]string{"Manufacturer", "Processor.Manufacturer"})
	fuzzyDemo.SetCaseSensitive(true)

	results, _ := fuzzyDemo.Search("eLnovo")
	if len(results) != 1 || results[0] != (*laptops)[1] {
		t.Error("Expected Lenovo laptop.")
	}
}

func TestMaxThreshold(t *testing.T) {
	laptops := generateLaptops()

	fuzzyDemo := NewFuzzy()
	fuzzyDemo.Set(laptops)
	fuzzyDemo.SetKeys([]string{"Manufacturer", "Processor.Manufacturer"})
	fuzzyDemo.Options.SetThreshold(10)

	results, _ := fuzzyDemo.Search("Elnovo")
	if len(results) != len((*laptops)) {
		t.Error(fmt.Sprintf("Expected %d elements, got %d instead.", len((*laptops)), len(results)))
	}
}

func TestMinThreshold(t *testing.T) {
	laptops := generateLaptops()

	fuzzyDemo := NewFuzzy()
	fuzzyDemo.Set(laptops)
	fuzzyDemo.SetKeys([]string{"Manufacturer", "Processor.Manufacturer"})
	fuzzyDemo.SetThreshold(0)

	results, _ := fuzzyDemo.Search("Elnovo")
	if len(results) != 0 {
		t.Error(fmt.Sprintf("Expected 0 elements, got %d instead.", len(results)))
	}
}
