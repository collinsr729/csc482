package main

import (
	// "os"
	"io/ioutil"
	// "log"
	"net/http"
	// "encoding/json"
	"fmt"
	loggly "github/loggly"
	gjson "github/gjson"
	// "reflect"
)

	type Data struct{
		// Response string
		Response struct{
			// Error string
			Error struct{
				Route string
				Message string
			}
		}
	}

func main() {
	//resp, err := http.Get("https://api.github.com/repos/collinsr729/csc482")
	resp, _ := http.Get("http://bus-time.centro.org/bustime/api/v3/getvehicles?key=jE6Q7MaB7MJMmRAwbB4yPXN4y&format=json&rt=OSW10")
	// resp := `
	// 	   {
 //        "bustime-response": {
 //                "error": [
 //                        {
 //                                "rt": "OSW10",
 //                                "msg": "No data found for parameter"
 //                        }
 //                ]
 //        }`
// }
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	
	value := gjson.Get(string(body), "bustime-response.error.0.msg")
	fmt.Println(value.Str)
	client := loggly.New("MyApplication")
	// client.EchoSend("info", "The api returned good stuff")
	if(value.Str != ""){
	client.EchoSend("error","The bustime api returned an error message of "+value.Str)
	}else{
		client.EchoSend("info", "The api returned good values")
	}
	// client.EchoSend("info", "The api returned good values")
	// client.EchoSend("silly", "This is a silly test message")//This works
}


	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.Stdout.Write(body)

	// d := &Data{}
	// err = json.Unmarshal([]byte(body), &d)
	// if(err!=nil){
	// 	log.Fatal(err)
	// }

	// fmt.Println(d)
	// fmt.Println(d.Response.Error.Route)
