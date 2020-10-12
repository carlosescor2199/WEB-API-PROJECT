package Contact

import (
	"../../Database"
	contactModel "../../Model/contact"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := Database.GetConnectionMongo()
	collection := client.Database("GO-REST-API").Collection("contacts")

	var contacts []contactModel.Contact

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)

	for res.Next(ctx) {
		var contact contactModel.Contact

		err := res.Decode(&contact)
		if err != nil {
			log.Fatal(err)
		}

		contacts = append(contacts, contact)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}

	_ = json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := Database.GetConnectionMongo()
	collection := client.Database("GO-REST-API").Collection("contacts")

	var contact contactModel.Contact
	// we get params with mux.
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&contact)

	if err != nil {
		log.Fatal(err)
	}

	_ = json.NewEncoder(w).Encode(contact)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := Database.GetConnectionMongo()
	collection := client.Database("GO-REST-API").Collection("contacts")

	var contact contactModel.Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)

	result, err := collection.InsertOne(ctx, contact)

	if err != nil {
		log.Fatal(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateContact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := Database.GetConnectionMongo()
	collection := client.Database("GO-REST-API").Collection("contacts")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var contact contactModel.Contact

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&contact)

	update := bson.D{
		{"$set", bson.D{
			{"firstname", contact.FirstName},
			{"lastname", contact.LastName},
			{"phoneNumber", contact.PhoneNumber},
			{"email", contact.Email},
			},
		},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&contact)

	if err != nil {
		log.Fatal(err)
	}

	contact.ID = id

	_ = json.NewEncoder(w).Encode(contact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	client, ctx := Database.GetConnectionMongo()
	collection := client.Database("GO-REST-API").Collection("contacts")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	_ = json.NewEncoder(w).Encode(deleteResult)
}