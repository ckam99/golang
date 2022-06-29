package main

import "fmt"

func main() {

	dic := map[string]int{
		"php":  19,
		"go":   98,
		"java": 89,
	}
	dic["c#"] = 75

	// Other way to define map
	var capitalCountry map[string]string = make(map[string]string)

	capitalCountry["Ivory Cost"] = "Abidjan"
	capitalCountry["Cameroun"] = "Yaounde"
	capitalCountry["Mali"] = "Bamako"

	// Read map value
	fmt.Println(dic["php"])
	fmt.Println(capitalCountry["Ivory Cost"])

	bamako, ok := capitalCountry["Mali"]
	if ok {
		fmt.Printf("Capital of Mali is %v\n", bamako)
	} else {
		fmt.Printf("Capital of Mali not found\n")
	}

	//
	for k, v := range dic {
		fmt.Printf("%v : %v\n", k, v)
	}

	for country := range capitalCountry {
		fmt.Printf("Capital of %v is %v \n", country, capitalCountry[country])
	}
}
