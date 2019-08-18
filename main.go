package main

// import "github.com/1026957152/myProject"

import (
	"awesomeProject/mymath"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	//"mymath"
	"net/http"
	"os"
)

//go.mod下载下来的内容存放在$GOPATH/pkg/mod/目录下
var client MQTT.Client

func main() {
	resp, err := http.Get("https://google.com")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	mymath.Square(1)
	fmt.Println(len(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
