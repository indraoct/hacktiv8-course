package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var (
		statusWind  string
		statusWater string
		keadaan     struct {
			Water int `json:"water"`
			Wind  int `json:"wind"`
		}
	)

	//ticker every 15 second
	ticker := time.NewTicker(1 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			water := randRange(1, 20)
			wind := randRange(1, 20)

			select {
			case <-done:
				return
			case <-ticker.C:

				//water
				switch {
				case water <= 8 && water >= 6:
					statusWater = "siaga"
				case water < 5:
					statusWater = "aman"
				default:
					statusWater = "bahaya"
				}

				//wind
				switch {
				case wind <= 15 && water >= 7:
					statusWind = "siaga"
				case water < 6:
					statusWind = "siaga"
				default:
					statusWind = "siaga"
				}

				keadaan.Wind = wind
				keadaan.Water = water
				b, _ := json.Marshal(keadaan)

				var prettyJSON bytes.Buffer
				_ = json.Indent(&prettyJSON, b, "", "\t")

				fmt.Println(string(prettyJSON.Bytes()))

				fmt.Println("status water : ", statusWater)
				fmt.Println("status wind : ", statusWind)

				fmt.Println("=================================")
			}
		}
	}()

	time.Sleep(3600 * time.Second) //program will running until 1 hour
	ticker.Stop()
	done <- true
	fmt.Println("Program stopped")
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
