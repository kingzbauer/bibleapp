package main

import "testing"


func TestExpandVerses(t *testing.T){
	tests := map[string][]string{
		"150.3" : []string{"150", "3"},
		"2.4" : []string{"2", "4"},
		"19.0" : []string{"19", "0"},
		"23399.34399234" : []string{"23399", "34399234"},
		"44.020" : []string{"44", "020"},
		"1.001" : []string{"1", "001"},
		"1.010" : []string{"1", "010"},
	}
	for k, result := range tests{
		expC, expV := ExpandVerse(k)
		if expC != result[0] || expV != result[1] {
			t.Error("ExpandVerse(", k, "); expected",
			"Chapter:", result[0], "Verse:", result[1],
			"But instead got Chapter:", expC,
			"Verse:", expV)
		}
	}
}

func TestFormatVerse(t *testing.T){
	tests := map[string]string{
		"1" : "100",
		"01" : "010",
		"001" : "001",
		"13" : "130",
		"013" : "013",
		"3" : "300",
		"21" : "210",
	}
	
	for key, value := range tests{
		if value != FormatVerse(key){
			t.Errorf("Expected %s but got %s", value, FormatVerse(key))
		}
	}
}
