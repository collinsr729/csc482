package main

import (
	"os"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gorilla/mux"
)

type TableItem struct {
	busID     string
	Latitiude string
	Longitude string
}

type Table struct {
	Items []TableItem
}
type TableInfo struct {
	tableName   string
	recordCount int
}

func getall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ALL PRINTING")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	// Show the session
	fmt.Printf("SESS: %s", sess.Config.Region)
	svc := dynamodb.New(sess)
	w.Header().Set("Content-Type", "application/json")
	// Create the Expression to fill the input struct with.
	// Get all movies in that year; we'll pull out those with a higher rating later
	// filt := expression.Name("Longitude").LessThan(expression.Value(0))

	// Or we could get by ratings and pull out those with the right year later
	//    filt := expression.Name("info.rating").GreaterThan(expression.Value(min_rating))

	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("BusID"), expression.Name("Latitiude"), expression.Name("Longitude"))

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("CentroBus"),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	fmt.Println("RESULT:", result, "\nERR:", err)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	items := []TableItem{}
	// table :=  Table{items}

	for _, i := range result.Items {
		item := TableItem{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(item.Latitiude)

		items = append(items, item)

	}

	json.NewEncoder(w).Encode(items)

}
func getstatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STATUS PRINTING")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := dynamodb.New(sess)
	w.Header().Set("Content-Type", "application/json")
	proj := expression.NamesList(expression.Name("BusID"), expression.Name("Latitiude"), expression.Name("Longitude"))

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("CentroBus"),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	fmt.Println("RESULT:", result, "\nERR:", err)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	numItems := 0

	for _, i := range result.Items {
		fmt.Println(i)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		numItems++
	}
	inf2 := map[string]interface{}{"Name of Table": "CentroBus", "Records": numItems}

	json.NewEncoder(w).Encode(inf2)

}

func main() {
	///mux router///
	r := mux.NewRouter()
	fmt.Println("Starting")
	r.HandleFunc("/rcollin3/all", getall).Methods("GET")

	r.HandleFunc("/rcollin3/status", getstatus).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
