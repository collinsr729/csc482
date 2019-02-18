package main

import (
	// "os"
	"io/ioutil"
	// "log"
	"net/http"
	// "encoding/json"
	// "fmt"
	loggly "github/loggly"
	gjson "github/gjson"
	"time"
)


func main() {
	counter := 0
	for counter<20{
		counter += 1
	resp, _ := http.Get("http://bus-time.centro.org/bustime/api/v3/getvehicles?key=jE6Q7MaB7MJMmRAwbB4yPXN4y&format=json&rt=OSW10")
	
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	
	value := gjson.Get(string(body), "bustime-response.error.0.msg")
	lat := gjson.Get(string(body), "bustime-response.vehicle.0.lat")
	lon := gjson.Get(string(body), "bustime-response.vehicle.0.lon")
	vehicleID := gjson.Get(string(body), "bustime-response.vehicle.0.vid")
	// fmt.Println(value.Str)
	// fmt.Println(body)
	client := loggly.New("MyApplication")
	if(value.Str != ""){
	client.EchoSend("error","The bustime api returned an error message of "+value.Str)
	}else{
		client.EchoSend("info", "The api returned Bus "+vehicleID.Str+" is at:"+lat.Str+","+lon.Str)
	}
	time.Sleep(time.Second*60)
	}
}