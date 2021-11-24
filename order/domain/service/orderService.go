package service

import (
	"fmt"

	"github.com/sreeks87/ordermgmt/order/domain"
)

type orderSvc struct {
	repo map[string]*domain.Order
}

func NewOrderSvc(r map[string]*domain.Order) domain.Service {
	return &orderSvc{
		repo: r,
	}
}

func (o *orderSvc) AddOrder(order domain.Order) (string, error) {
	o.repo[order.OrderID] = &order
	return order.OrderID, nil
}

func (o *orderSvc) ShipmentUpdate(skus []string, shipid string, orderid string) (string, error) {
	if b, e := o.Validate(skus, shipid, orderid); !b {
		return "", e
	}
	// imagine we have duplicate skuid in the order, then we dont need o update all the suid with the shipping id,
	// we just update the first one with the shippingid and break
	for _, s := range skus {
		for _, skuInOrder := range o.repo[orderid].Skus {
			if s == skuInOrder.SKUId && skuInOrder.ShipmentId == "" {
				skuInOrder.ShipmentId = shipid
				break
			}
		}
	}
	return orderid, nil
}

func (o *orderSvc) GetShipment(orderid string) ([]*domain.SKU, error) {
	if val, ok := o.repo[orderid]; ok {
		return val.Skus, nil
	}
	return nil, fmt.Errorf("order id %s not found", orderid)
}

func (o *orderSvc) Validate(skus []string, shipid string, orderid string) (bool, error) {
	if _, ok := o.repo[orderid]; ok {
		order := o.repo[orderid]
		var skuMapFromOrder = make(map[string]int, len(o.repo[orderid].Skus))
		// get a map of all skus and quantity in an order where shipment id is not updated yet
		for _, sku := range order.Skus {
			if sku.ShipmentId == "" {
				skuMapFromOrder[sku.SKUId] += 1
			}
		}

		for _, s := range skus {
			// 3.There is no such SKU 'SKU3' in the original order, so this update is invalid.
			// all elements of skus should be present in skuSlicefromOrder
			// if skus contain more than what is in skuSlicefromOrder
			// then it is an invalid request
			if _, ok := skuMapFromOrder[s]; !ok {
				return false, fmt.Errorf("SKU %s is either not present or has more quantity than in original order", s)
			}
			// 1. if the shipment information is for more SKUs than there is room for to update
			// keep decrementing the quantiy as we meet each sku in the skus
			// when count ==0, delete the element from map, next occurence of the same sku means
			// we have more SKU than there is room for update.
			// 2. Order-1 got SKU1 shipped with tracking-2
			// There are no SKU1s left that need shipment update, the count of sku1 willbe zero
			// therefore it will be treated as an invalid request
			skuMapFromOrder[s] -= 1
			if skuMapFromOrder[s] == 0 {
				delete(skuMapFromOrder, s)
			}
		}
		return true, nil
	}
	return false, fmt.Errorf("order id %s is invalid", orderid)

}
