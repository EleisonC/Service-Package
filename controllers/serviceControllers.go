package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/EleisonC/Service-Package/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	ownerID := vars["ownerID"]
	serviceTypeID := vars["serviceTypeID"]
	service.OwnerID = ownerID
	service.ServiceTypeId = serviceTypeID
	service.Status = false


	if validateErr := validate.Struct(&service); validateErr != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := models.CreateService(service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func CreateServiceTypeHandler(w http.ResponseWriter, r *http.Request) {
	var serviceType models.ServiceType
	if err := json.NewDecoder(r.Body).Decode(&serviceType); err != nil {
		http.Error(w, "invalid request body1", http.StatusBadRequest)
		return
	}

	serviceType.Status = false
	serviceType.ID = primitive.NewObjectID()
	serviceType.DelTrue = false

	if validateErr := validate.Struct(&serviceType); validateErr != nil {
		http.Error(w, "invalid request body2", http.StatusBadRequest)
		return
	}

	err := models.CreateServiceType(serviceType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ownerID := vars["ownerID"]
	services, err := models.ReadServiceByOwnerID(ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}

func GetAllServiceTypesHandler(w http.ResponseWriter, r *http.Request) {
	servicesTypes, err := models.ReadServiceTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(servicesTypes)
}

func UpdateServiceTypeHandler(w http.ResponseWriter, r *http.Request) {
	var serviceType models.ServiceType
	if err := json.NewDecoder(r.Body).Decode(&serviceType); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	serviceTypeId := vars["serviceTypeId"]
	serviceType.Status = false
	serviceTypeID, err := primitive.ObjectIDFromHex(serviceTypeId)
	if err != nil { 
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if validateErr := validate.Struct(&serviceType); validateErr != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	update := map[string]interface{}{
		"name": serviceType.Name,
		"description": serviceType.Description,
		"status": serviceType.Status,
	}

	err = models.UpdateServiceTypeByID(serviceTypeID, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	serviceId := vars["serviceId"]
	ownerID := vars["ownerID"]
	serviceTypeID := vars["serviceTypeId"]
	service.Status = false
	service.ServiceTypeId = serviceTypeID
	serviceID, err := primitive.ObjectIDFromHex(serviceId)
	if err != nil { 
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if validateErr := validate.Struct(&service); validateErr != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	update := map[string]interface{}{
		"name": service.Name,
		"description": service.Description,
		"serviceTypeId": service.ServiceTypeId,
		"status": service.Status,
		"ownerId": ownerID,
	}

	err = models.UpdateServiceByIDAndOwnerID(serviceID, ownerID, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteServiceByIDAndOwnerIdHandler(w http.ResponseWriter, r * http.Request) {

	vars := mux.Vars(r)
	serviceId := vars["serviceId"]
	ownerID := vars["ownerID"]
	serviceID, err := primitive.ObjectIDFromHex(serviceId)
	if err != nil { 
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	update := map[string]interface{}{
		"delValue": true,
	}

	err = models.DeleteServiceByIDAndOwnerID(serviceID, ownerID, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteServiceTypeByIDHandler(w http.ResponseWriter, r * http.Request) {

	vars := mux.Vars(r)
	serviceTypeId := vars["serviceTypeId"]
	ownerID := vars["ownerID"]
	serviceTypeID, err := primitive.ObjectIDFromHex(serviceTypeId)
	if err != nil { 
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	update := map[string]interface{}{
		"delValue": true,
	}

	err = models.DeleteServiceTypeByID(serviceTypeID, ownerID, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

