package rest

import (
	"encoding/json"
	"net/http"
	"ordersystem/model"
	"ordersystem/repository"

	"github.com/go-chi/render"
)

// GetMenu 			godoc
// @tags 			Menu
// @Description 	Returns the menu of all drinks
// @Produce  		json
// @Success 		200 {array} model.Drink
// @Router 			/api/menu [get]
func GetMenu(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Menu := db.GetDrinks()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, Menu)
		// todo
		// get slice from db
		// render.Status(r, http.StatusOK)
		// render.JSON(w, r, <your-slice>)
	}
}

// todo create GetOrders /api/order/all
func GetOrders(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Orders := db.GetOrders()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, Orders)
	}
}

// todo create GetOrdersTotal /api/order/total
func GetOrdersTotal(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Totalorders := db.GetTotalledOrders()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, Totalorders)
	}
}

// PostOrder 		godoc
// @tags 			Order
// @Description 	Adds an order to the db
// @Accept 			json
// @Param 			b body model.Order true "Order"
// @Produce  		json
// @Success 		200
// @Failure     	400
// @Router 			/api/order [post]
func PostOrder(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		order := model.Order{}
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
		} else {
			db.AddOrder(&order)
			render.Status(r, http.StatusOK)
			render.JSON(w, r, "ok")
		}
		// todo
		// declare empty order struct
		// err := json.NewDecoder(r.Body).Decode(&<your-order-struct>)
		// handle error and render Status 400
		// add to db

	}
}
