package main

import (
	"encoding/json"
	"indexer/consumer/events"
	"io"
	"log"
	"os"
)

func main() {
	parsed_json := unmarshal()
	log.Println(parsed_json)
}

// func getObject() ( Person) {
// 	r, err := http.Get("https://swapi.dev/api/people/1")
// 	check(err, "Calling SW people API")
// 	defer r.Body.Close()

// 	data, err := io.ReadAll(r.Body)
// 	check(err, "Read JSON from response")

// 	err = json.Unmarshal(data, &person)
// 	check(err, "Unmarshalling")
// 	return
// }

func unmarshal() (checkpoint events.TransferObject) {
	file, err := os.Open("events/samples/transfer_object.json")
	if err != nil {
		log.Println("Error opening json file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading json data:", err)
	}

	err = json.Unmarshal(data, &checkpoint)
	if err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func check(err error, msg string) {
	if err != nil {
		log.Println("Error encountered:", msg)
	}
}
