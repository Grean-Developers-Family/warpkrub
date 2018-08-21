package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to WarpKrub.")
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "alpha 0.0.1")
}

func Warps(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("secrets.json")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully Open secrets.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var warps Warp

	json.Unmarshal(byteValue, &warps)

	//add instragram url
	for i := 0; i < len(warps.Warp); i++ {
		if warps.Warp[i].Ig != "" {
			warps.Warp[i].Ig = strings.Replace(warps.Warp[i].Ig, "@", "", 1)
			warps.Warp[i].Ig = "https://www.instagram.com/" + warps.Warp[i].Ig
		}
	}

	if err := json.NewEncoder(w).Encode(warps); err != nil {
		panic(err)
	} else {
		log.Println("let's get Warp!!")
	}
}

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	// os.Setenv("DB", "admin:password@11.57.69.69")
// 	// os.Setenv("APIKEY", "P99zE8YiLm")
// 	todos := Todos{
// 		Todo{Version: "1.0.0",
// 			Hostname: getHostname(),
// 			Time:     time.Now(),
// 			// Db:       os.Getenv("DB"),
// 			// Apikey:   os.Getenv("APIKEY"),
// 			Db:     "admin:password@11.57.69.69",
// 			Apikey: "P99zE8YiLm",
// 		},
// 	}
// 	//add commanet
// 	if err := json.NewEncoder(w).Encode(todos); err != nil {
// 		panic(err)
// 	} else {
// 		log.Println("Request path /version  httpCode:200")
// 	}
// }
