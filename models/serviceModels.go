package models

import (
	"context"
	"fmt"
	"time"

	"github.com/EleisonC/Service-Package/configSetup"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type ServiceType struct {
	ID primitive.ObjectID `bson:"_id"`
	Name        string `bson:"name" validate:"required"`
	Description string `bson:"description" validate:"required"`
	Status      bool `bson:"status"`
	DelTrue bool `bson:"delValue, omitempty"`
}

type Service struct {
	ID primitive.ObjectID `bson:"_id"`
	Name        string `bson:"name" validate:"required"`
	Description string `bson:"description" validate:"required"`
	Status      bool `bson:"status"`
	ServiceTypeId string `bson:"serviceTypeId" validate:"required"`
	OwnerID string `bson:"ownerId" validate:"required"`
	DelTrue bool `bson:"delValue, omitempty"`
}

var services_collection *mongo.Collection = configs.GetCollection(configs.DB, "services")
var service_types_collection *mongo.Collection = configs.GetCollection(configs.DB, "service_types")

func CreateService(service Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := services_collection.InsertOne(ctx, service)
	if err != nil {
		return err
	}
	return nil
}

func CreateServiceType(serviceType ServiceType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := service_types_collection.InsertOne(ctx, serviceType)
	if err != nil {
		return err
	}
	return nil
}

func ReadServiceByOwnerID(ownerID string)([]Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var services []Service
	cursor, err := services_collection.Find(ctx, bson.M{"ownerId": ownerID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &services); err != nil {
		return nil, err
	}
	return services, nil
}

func ReadServiceTypes()([]ServiceType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var servicesTypes []ServiceType
	cursor, err := service_types_collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &servicesTypes); err != nil {
		return nil, err
	}
	return servicesTypes, nil
}

func ReadServiceByIDAndOwnerID(serviceId primitive.ObjectID, ownerId string)(Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var service Service
	filter := bson.M{"_id": serviceId, "owner_id": ownerId}
	err := services_collection.FindOne(ctx, filter).Decode(&service)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Service{}, err
		}
		return Service{}, err
	}
	return service, nil
}

func ReadServiceTypeByName(serviceTypeId primitive.ObjectID)(ServiceType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var serviceType ServiceType
	filter := bson.M{"_id": serviceTypeId}
	err := services_collection.FindOne(ctx, filter).Decode(&serviceType)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ServiceType{}, err
		}
		return ServiceType{}, err
	}
	return serviceType, nil
}

func UpdateServiceByIDAndOwnerID(id primitive.ObjectID, ownerId string, updates map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id, "owner_id": ownerId}
	update := bson.M{"$set": updates}
	result, err := services_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("service not found or does not belong to owner")
	}
	return nil
}

func UpdateServiceTypeByID(id primitive.ObjectID, updates map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updates}
	result, err := service_types_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("service not found or does not belong to owner")
	}
	return nil
}

func DeleteServiceByIDAndOwnerID(id primitive.ObjectID, ownerId string, updates map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id, "owner_id": ownerId}
	update := bson.M{"$set": updates}
	result, err := services_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("service not found or does not belong to owner")
	}
	return nil
}

func DeleteServiceTypeByID(id primitive.ObjectID, ownerId string, updates map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updates}
	result, err := service_types_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("service not found or does not belong to owner")
	}
	return nil
}


