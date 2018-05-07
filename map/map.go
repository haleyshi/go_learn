package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["中国"] = "北京"
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"

	for country := range countryCapitalMap {
		fmt.Println(country, "with capital", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["USA"]

	if ok {
		fmt.Println("USA with capital", capital)
	} else {
		fmt.Println("Capital for USA is not defined")
	}

	delete(countryCapitalMap, "France") // delete one k-v pair
	for country := range countryCapitalMap {
		fmt.Println(country, "with capital", countryCapitalMap[country])
	}
}
