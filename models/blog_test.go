package models

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"blog/database"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"go.mongodb.org/mongo-driver/bson"
// )


// func TestCustomerCRUD(t *testing.T) {

// 	testData := `{
// 		"firstname": "Tade",
// 		"lastname": "Ugo",
// 		"phone": "08162454567",
// 		"email": "adamu@gmail.com",
// 		"address":"11, sisi close Opebi"
// 	}`

// 	collection, err := database.GetMongoDbCollection("testdb", "customer_test_collection")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var customer Customer 

// 	json.Unmarshal([]byte(testData), &customer)

// 	response, err := collection.InsertOne(context.Background(), customer)
	
// 	//Test that customer was created successfully
// 	require.NoError(t, err)

// 	filter := bson.M{"_id": response.InsertedID}

// 	var results []bson.M
// 	cur, err := collection.Find(context.Background(), filter)
// 	defer cur.Close(context.Background())

// 	cur.All(context.Background(), &results)

// 	var cusResponse Customer

// 	// convert result to Customer struct
// 	bsonBytes, _ := bson.Marshal(results[0])
// 	bson.Unmarshal(bsonBytes, &cusResponse)

// 	//Test that retrieved document equates to input document
// 	require.Equal(t, "Tade", cusResponse.FirstName)
// 	require.Equal(t, "Ugo", cusResponse.LastName)
// 	require.Equal(t, "08162454567", cusResponse.Phone)
// 	require.Equal(t, "adamu@gmail.com", cusResponse.Email)
// 	require.Equal(t, "11, sisi close Opebi", cusResponse.Address)

	
// 	//update retrived document
// 	cusResponse.FirstName = "Adedamola"

// 	update := bson.M{
// 		"$set": cusResponse,
// 	}

// 	collection.UpdateOne(context.Background(), bson.M{"_id": response.InsertedID}, update)

// 	updateCur, err := collection.Find(context.Background(), filter)
// 	defer updateCur.Close(context.Background())

// 	updateCur.All(context.Background(), &results)

// 	var updateCustomer Customer

// 	// convert result to Customer struct
// 	upBsonBytes, _ := bson.Marshal(results[0])
// 	bson.Unmarshal(upBsonBytes, &updateCustomer)

// 	require.Equal(t, "Adedamola", updateCustomer.FirstName)

// 	//Delete Test collection
// 	if err = collection.Drop(context.Background()); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("customer_test_collection deleted successfully")

// }