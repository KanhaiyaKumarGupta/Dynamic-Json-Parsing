package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonData := []byte(`{
	"key1": "value1",
    "key2": 123,
    "key3": [1, 2, 3],
    "key4": ["nestedKey","nestedValue"]
	}`)

	var result interface{}

	err := json.Unmarshal(jsonData, &result)

	if err != nil {
		log.Fatal(err)
		return
	}

	data, ok := result.(map[string]interface{})

	if !ok {
		fmt.Println("Invalid Json Structure")
		return
	}

	key1, keyexists := data["key4"].([]interface{})

	if keyexists {
		for _, value := range key1 {
			fmt.Println(value)
		}
	}
}
