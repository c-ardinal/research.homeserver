package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antonholmquist/jason"
	"github.com/julienschmidt/httprouter"
)

var LocalJSON *jason.Object
var Devices map[string]*jason.Object = map[string]*jason.Object{}

// デバイス一覧
func getDevices(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	format, _ := LocalJSON.Object()
	formated, _ := json.MarshalIndent(format, "", "\t")
	fmt.Fprintf(w, "%s", formated)
}

// デバイス情報取得
func getDevice(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	devices, _ := LocalJSON.GetObjectArray("devices")
	for _, device := range devices {
		deviceName, _ := device.GetString("name")
		if deviceName == p.ByName("name") {
			formated, _ := json.MarshalIndent(device, "", "\t")
			fmt.Fprintf(w, "%s", formated)
			return
		}
	}
	fmt.Fprintf(w, "{\"error\": \"NotFound\"}")
}

func addDevice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "WIP")
}

func fixDevice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "WIP")
}

func deleteDevice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "WIP")
}

func doControl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	j, err := jason.NewObjectFromReader(r.Body)
	if err != nil {
		fmt.Fprintf(w, "{\"Faild\": \"%s\"", err.Error())
		return
	}
	jo, err := j.Object()
	if err != nil {
		fmt.Fprintf(w, "{\"Faild\": \"%s\"", err.Error())
		return
	}
	fmt.Println(jo)

	requests, _ := j.GetObjectArray("operations")
	for _, request := range requests {
		target, err := request.GetString("target")
		if err != nil {
			fmt.Fprintf(w, "{\"Faild\": \"%s\"", err.Error())
			return
		}
		operation, err := request.GetString("operation")
		if err != nil {
			fmt.Fprintf(w, "{\"Faild\": \"%s\"", err.Error())
			return
		}

		api, _ := Devices[target].GetObject("api")
		ops, _ := api.GetObjectArray("operations")
		for _, op := range ops {
			opName, _ := op.GetString("name")
			if opName == operation {
				path, _ := op.GetString("path")
				endpoint, _ := api.GetString("endpoint")
				url := endpoint + path
				fmt.Println("SendRequest: " + url)
				//sendRequest(url, "POST", "")
			}
		}
	}
	fmt.Fprintf(w, "{\"Accept\": \"%s\"", jo)
}

func doScan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "WIP")
}

func sendRequest(url string, method string, body string) {
	var req *http.Request

	fmt.Println()
	if body == "" {
		req, _ = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func initServer() {
	LocalJSON, _ = jason.NewObjectFromReader(os.Stdin)
	devices, _ := LocalJSON.GetObjectArray("devices")
	for _, device := range devices {
		deviceName, _ := device.GetString("name")
		Devices[deviceName] = device
	}
	fmt.Println(LocalJSON.GetObject())

	router := httprouter.New()
	// デバイス一覧
	router.GET("/devices", getDevices)

	// デバイス情報操作
	router.GET("/device/:name", getDevice)
	router.POST("/device/:name", addDevice)
	router.PUT("/device/:name", fixDevice)
	router.DELETE("/device/:name", deleteDevice)

	// デバイス操作
	router.POST("/control", doControl)
	router.POST("/scan", doScan)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	initServer()
}
