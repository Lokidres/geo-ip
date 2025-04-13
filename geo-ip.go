package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GeoIPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run geo-ip.go <ip-address>")
		os.Exit(1)
	}

	ip := os.Args[1]
	url := "http://ipinfo.io/" + ip + "/json"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var ipInfo GeoIPInfo
	json.Unmarshal(body, &ipInfo)

	fmt.Println("=== GEO-IP LOCATION FINDER ===")
	fmt.Println("IP:", ipInfo.IP)
	fmt.Println("City:", ipInfo.City)
	fmt.Println("Region:", ipInfo.Region)
	fmt.Println("Country:", ipInfo.Country)
	fmt.Println("Location:", ipInfo.Loc)
	fmt.Println("Timezone:", ipInfo.Timezone)
	fmt.Println("============================")
}
