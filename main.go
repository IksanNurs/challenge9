package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	var data1 Data
	var windStatus string
	var waterStatus string
	var resp *http.Response
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		data := &Data{
			Water: water,
			Wind:  wind,
		}

		requestBody, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		url := "https://jsonplaceholder.typicode.com/posts"
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal(body, &data1)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != http.StatusCreated {
			fmt.Println("POST request failed")
		} else {
			b, err := json.MarshalIndent(data1, "", " ")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s\n", b)

			if data1.Water < 5 {
				waterStatus = "aman"
			} else if data1.Water >= 6 && data1.Water <= 8 {
				waterStatus = "siaga"
			} else if data1.Water > 8 {
				waterStatus = "bahaya"
			}

			if data1.Wind < 6 {
				windStatus = "aman"
			} else if data1.Wind >= 7 && data1.Wind <= 15 {
				windStatus = "siaga"
			} else if data1.Wind > 15 {
				windStatus = "bahaya"
			}
			fmt.Println("status water:", waterStatus)
			fmt.Println("status wind:", windStatus)

		}

		fmt.Println()

	}
	defer resp.Body.Close()
}
