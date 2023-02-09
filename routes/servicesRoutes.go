package routes

import (
	"github.com/EleisonC/Service-Package/controllers"
	"github.com/gorilla/mux"
)


var RegisterServiceRoutes = func(router *mux.Router) {
	router.HandleFunc("/{ownerID}/{serviceTypeID}/createservice/", controllers.CreateServiceHandler).Methods("POST")
	router.HandleFunc("/createservicetype/", controllers.CreateServiceTypeHandler).Methods("POST")
	router.HandleFunc("/{ownerID}/getallservices/", controllers.GetAllServicesHandler).Methods("GET")
	router.HandleFunc("/getallservicestypes/", controllers.GetAllServiceTypesHandler).Methods("GET")
	router.HandleFunc("/{serviceType}/updateservicestype/", controllers.UpdateServiceTypeHandler).Methods("PUT")
	router.HandleFunc("/{ownerID}/{serviceTypeId}/{serviceId}/updateservices/", controllers.UpdateServiceHandler).Methods("PUT")
	router.HandleFunc("/{ownerID}/{serviceId}/deleteservices/", controllers.UpdateServiceHandler).Methods("DELETE")
	router.HandleFunc("/{ownerID}/{serviceId}/deleteservices/", controllers.UpdateServiceHandler).Methods("DELETE")
}
