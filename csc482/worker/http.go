package main

import (
	"os"
	"io/ioutil"
	"net/http"
	"fmt"
	loggly "github.com/jamespearly/loggly"
	gjson "github.com/tidwall/gjson"
	"time"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct{
	BusID string `json:"busID"`
	Latitiude string `json:"Latitiude"`
	Longitude string `json: "Longitude"`
}





func main() {
	// Create DynamoDB client
	sess, err := session.NewSession(&aws.Config{
	Region: aws.String("us-east-1")},
	)
	svc := dynamodb.New(sess)

	if err != nil {
	    fmt.Println("Error creating session:")
	    fmt.Println(err.Error())
	    os.Exit(1)
	}



	counter := 0
	//while(true) below
	for counter<40{
		counter += 1
		resp, _ := http.Get("http://bus-time.centro.org/bustime/api/v3/getvehicles?key=jE6Q7MaB7MJMmRAwbB4yPXN4y&format=json&rt=OSW10")
		
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		
		value := gjson.Get(string(body), "bustime-response.error.0.msg")
		lat := gjson.Get(string(body), "bustime-response.vehicle.0.lat")
		timeStamp := gjson.Get(string(body), "bustime-response.vehicle.0.tmstmp")
		lon := gjson.Get(string(body), "bustime-response.vehicle.0.lon")
		vehicleID := gjson.Get(string(body), "bustime-response.vehicle.0.vid")
		// fmt.Println(value.Str)
		// fmt.Println(body)
		client := loggly.New("MyApplication")

		item := Item{
			BusID: timeStamp.Str, 
			Latitiude: lat.Str,
			Longitude: lon.Str,
		}

		if(value.Str != ""){
			client.EchoSend("error","The bustime api returned an error message of "+value.Str)
		}else{
			client.EchoSend("info", "The api returned Bus "+vehicleID.Str+" is at:"+lat.Str+","+lon.Str)
			av, err := dynamodbattribute.MarshalMap(item)

			if err != nil {
				fmt.Println("Got error marshalling map:")
				fmt.Println(err.Error())
				os.Exit(1)
			}
			    //fmt.Println(av)
			_, err = svc.PutItem(&dynamodb.PutItemInput{
			    TableName: aws.String("CentroBus"),
			    Item:      av,
			})
			if err != nil {
			    panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
			}
		}
		
		//wait a minute between all requests
		time.Sleep(time.Second*60)
		
		
		    

		// "hacking" the program to make it run continueously
		counter -= 1
	}
}
