package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/rustingoff/internal/entities"
	"github.com/rustingoff/pkg/repository"
	"github.com/rustingoff/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func getAllProductsMongo(w http.ResponseWriter, r *http.Request, urlPattern string) {
	w.Header().Add("content-type", "application/json")
	var products []entities.ProductMGDB

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client = repository.NewClientMongoDB(ctx)
	collection := client.Database("crud_system").Collection("products")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product entities.ProductMGDB
		cursor.Decode(&product)
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func getProductByIdMongo(w http.ResponseWriter, r *http.Request, urlPattern string) {
	w.Header().Add("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(r.URL.Path, urlPattern))
	var product entities.ProductMGDB

	json.NewDecoder(r.Body).Decode(&product)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client = repository.NewClientMongoDB(ctx)
	collection := client.Database("crud_system").Collection("products")
	err := collection.FindOne(ctx, entities.ProductMGDB{ID: id}).Decode(&product)
	if err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func addProductMongo(w http.ResponseWriter, r *http.Request, urlPattern string) {
	w.Header().Add("content-type", "application/json")

	var product entities.ProductMGDB

	json.NewDecoder(r.Body).Decode(&product)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client = repository.NewClientMongoDB(ctx)
	collection := client.Database("crud_system").Collection("products")
	result, _ := collection.InsertOne(ctx, product)

	json.NewEncoder(w).Encode(result)
}

func updateProductMongo(w http.ResponseWriter, r *http.Request, urlPattern string) {
	w.Header().Add("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(r.URL.Path, urlPattern))
	var product entities.ProductMGDB

	json.NewDecoder(r.Body).Decode(&product)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client = repository.NewClientMongoDB(ctx)
	collection := client.Database("crud_system").Collection("products")
	result, err := collection.UpdateOne(ctx, entities.ProductMGDB{ID: id}, product)
	if err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func deleteProductMongo(w http.ResponseWriter, r *http.Request, urlPattern string) {
	w.Header().Add("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(r.URL.Path, urlPattern))
	var product entities.ProductMGDB

	json.NewDecoder(r.Body).Decode(&product)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client = repository.NewClientMongoDB(ctx)
	collection := client.Database("crud_system").Collection("products")
	result, err := collection.DeleteOne(ctx, entities.ProductMGDB{ID: id})
	if err != nil {
		utils.ServerResponse(w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
