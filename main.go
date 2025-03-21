package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Schedule struct {
	Time  string `json:"time"` // groepeert ze in schedule struct
	Scene string `json:"scene"`
}

func main() {
	data, _ := os.ReadFile("config.json")
	var schedules []Schedule
	json.Unmarshal(data, &schedules)

	for {
		currentTime := time.Now().Format("15:04")
		for _, s := range schedules {
			if s.Time == currentTime {
				setLightScene(s.Scene)
			}
		}
		time.Sleep(time.Minute)
	}
}

func setLightScene(scene string) {
	url := "http:// + bridgeIP + /api/ + appid + /lights"
	jsonData := fmt.Sprintf(`{"scene": "%s"}`, scene)
	_, err := http.Post(url, "application/json", strings.NewReader(jsonData))
	if err != nil {
		fmt.Println("Fout bij het instellen van de scène:", err)
	} else {
		fmt.Println("Scène ingesteld:", scene)
	}
}
