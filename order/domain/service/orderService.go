package service

import (
	"errors"

	"github.com/sreeks87/ordermgmt/order/domain"
)

type orderSvc struct {
	repo map[string]domain.Order
}

func NewOrderSvc(r map[string]domain.Order) domain.Service {
	return &orderSvc{
		repo: r,
	}
}

func (o *orderSvc) AddOrder(order domain.Order) (string, error) {
	o.repo[order.OrderID] = order
	return order.OrderID, nil
}

func (o *orderSvc) ShipmentUpdate(skus []string, shipid string, orderid string) (string, error) {
	if _, ok := o.repo[orderid]; ok {
		order := o.repo[orderid]
		// if the len of skus and order.skus does not match
		if len(skus) != len(order.Skus) {
			return "", errors.New("sku length does not match")
		}
		// do not update the map yet, the validation may fail at the end of the process too
		for _, sku := range order.Skus {
			for _, s := range skus {
				if sku.SKUId == s {
					// valid for update
					// update the first occurence of the skuid
					// example the repo contains sku1 sku2 and skus contains [sku1,sku1] for
					// shipid = 1234 and orderid=1
					// then update the first occurence of sku1 with shipid, that has a shipid==nil
					// since it doesnt matter what sku we update the shipid with this will guarentee that
					// sku1 will get updated atleast once.
					// break once the first skuid is updated with trackingid
					if sku.ShipmentId == "" {
						sku.ShipmentId = shipid
						break
					}
				}
			}
		}

	}
	return "", nil
}

func (o *orderSvc) GetShipment(shipid string) ([]domain.SKU, error) {
	return []domain.SKU{}, nil
}
