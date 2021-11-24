package main

import (
	"fmt"

	"github.com/sreeks87/ordermgmt/order/domain"
	"github.com/sreeks87/ordermgmt/order/domain/service"
)

func main() {

	inmemoryRepo := make(map[string]*domain.Order)

	svc := service.NewOrderSvc(inmemoryRepo)
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	sku2 := domain.SKU{
		SKUId: "SKU2",
	}
	sku3 := domain.SKU{
		SKUId: "SKU1",
	}
	skus1 := []*domain.SKU{
		&sku1,
		&sku3,
		&sku2,
	}
	o1 := domain.Order{
		OrderID: "ord1",
		Skus:    skus1,
	}
	var id string
	var e error
	if id, e = svc.AddOrder(o1); e != nil {
		fmt.Println(e)
	}
	fmt.Println("Order placed with ID :", id)
	var oid string

	// 	Shipment with tracking tracking-1 ships SKU1, SKU2
	// Shipment with tracking tracking-2 ships SKU1
	// Updates for a Shipment can arrive in multiple parts. For example:

	// # update 1 at timestamp-1
	if oid, e = svc.ShipmentUpdate([]string{sku1.SKUId}, "Tracking1", o1.OrderID); e != nil {
		fmt.Println(e)
	}
	fmt.Println("updated shipment for ", oid)
	if oid, e = svc.ShipmentUpdate([]string{sku2.SKUId}, "Tracking1", o1.OrderID); e != nil {
		fmt.Println(e)
	}
	fmt.Println("updated shipment for ", oid)

	if oid, e = svc.ShipmentUpdate([]string{"SKU3", sku1.SKUId}, "Tracking1", o1.OrderID); e != nil {
		fmt.Println(e)
	}
	fmt.Println("updated shipment for ", oid)

	var ship []*domain.SKU
	if ship, e = svc.GetShipment(o1.OrderID); e != nil {
		fmt.Println(e)
	}
	fmt.Println(o1.OrderID, " Shipment details ")
	for _, v := range ship {
		fmt.Println(v.SKUId, " | ", v.ShipmentId)
	}

}
