package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Espresso", Price: 2.5, Description: "Strong and bold coffee shot"},
		{ID: 2, Name: "Cappuccino", Price: 3.0, Description: "Espresso with steamed milk and foam"},
		{ID: 3, Name: "Iced Latte", Price: 3.5, Description: "Chilled coffee with milk and ice"},
	}
	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Now().Add(-2 * time.Hour), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Now().Add(-1 * time.Hour), Amount: 1},
		{DrinkID: 3, CreatedAt: time.Now(), Amount: 3},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}

}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for i := 0; i < len(db.orders); i++ {
		current_order := db.orders[i]

		key := uint64(current_order.DrinkID)
		value := uint64(current_order.Amount)
		totalledOrders[key] += value
	}
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)
	// todo
	// add order to db.orders slice

}
