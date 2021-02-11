package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"web-golang-restapi/helpers"
)

//Customer - information about Customer Model
type Customer struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

//OrderItem - information about OrderItem Model
type OrderItem struct {
	ProductID int     `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
}

//OrderRequestModel - information about OrderRequestModel json request
type OrderRequestModel struct {
	Customer   Customer    `json:"customer"`
	OrderItems []OrderItem `json:"orderItems"`
}

func (orderRequestModel *OrderRequestModel) Save() error {
	ctx := context.Background()
	tx, err := helpers.DB.BeginTx(ctx, nil)

	if err != nil {
		fmt.Println(err)
	}

	//save customer
	custInserStmt, _ := tx.PrepareContext(ctx, "INSERT INTO customer (name,address,gender,phoneNumber) VALUES(?,?,?,?)")
	defer custInserStmt.Close()

	_, err = custInserStmt.ExecContext(ctx, orderRequestModel.Customer.Name, orderRequestModel.Customer.Address,
		orderRequestModel.Customer.Gender, orderRequestModel.Customer.PhoneNumber)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	//save order
	orderInsertStmt, _ := tx.PrepareContext(ctx, "INSERT INTO orders (total, order_date) VALUES(?,?)")
	defer orderInsertStmt.Close()

	var orderResult sql.Result
	orderDate := time.Now()
	total := getTotal(orderRequestModel.OrderItems)
	orderResult, err = orderInsertStmt.ExecContext(ctx, total, orderDate)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	//save orderitem
	id, _ := orderResult.LastInsertId()
	query, vals := formatOrderItemQuery(orderRequestModel.OrderItems, id)
	orderItemInsertStmt, _ := tx.PrepareContext(ctx, query)
	defer orderItemInsertStmt.Close()

	_, err = orderItemInsertStmt.ExecContext(ctx, vals...)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getTotal(OrderItems []OrderItem) float64 {
	total := 0.0
	for _, orderItem := range OrderItems {
		subtotal := orderItem.Price * float64(orderItem.Quantity)
		total += subtotal
	}
	return total
}

func formatOrderItemQuery(OrderItems []OrderItem, orderID int64) (string, []interface{}) {
	sqlStr := "INSERT INTO orderitem(product_id, quantity, sub_total,order_id,price) VALUES"
	vals := []interface{}{}

	for _, orderItem := range OrderItems {
		sqlStr += "(?,?,?,?,?),"
		subtotal := orderItem.Price * float64(orderItem.Quantity)
		vals = append(vals, orderItem.ProductID, orderItem.Quantity, subtotal, orderID, orderItem.Price)
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	fmt.Println(sqlStr)
	fmt.Println(vals...)

	return sqlStr, vals
}
