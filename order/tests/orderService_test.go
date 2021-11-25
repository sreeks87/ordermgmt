package service_test

import (
	"testing"

	"github.com/sreeks87/ordermgmt/order/domain"
	"github.com/sreeks87/ordermgmt/order/service"
	"github.com/stretchr/testify/assert"
)

func testSVC() domain.Service {
	inmemoryRepo := make(map[string]*domain.Order)
	return service.NewOrderSvc(inmemoryRepo)
}

func TestAddOrderSuccess(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "test1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	oId, _ := s.AddOrder(o1)
	assert.Equal(t, oId, o1.OrderID)
}

func TestAddOrderNoOrderID(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	assert.Equal(t, e.Error(), "order should contain an order ID")
}

func TestGetshipmentSuccess(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	sku, e := s.GetShipment(o1.OrderID)
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, sku[0].ShipmentId, "")

}

func TestGetShipmentWithoutOrderID(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.GetShipment("")
	assert.Equal(t, e.Error(), "no order ID specified")

}

func TestGetShipmentInvalidOrderID(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.GetShipment("order100")
	assert.Equal(t, e.Error(), "order id order100 not found")

}

func TestValidateEmptySKUList(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.Validate([]string{}, "trackin1", "order1")
	assert.Equal(t, e.Error(), "the list of skus to be updated cant be empty")

}

func TestValidateEmptyShipTrackID(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.Validate([]string{"SKU1"}, "", "order1")
	assert.Equal(t, e.Error(), "shipment tracking id cant be nil")

}

func TestValidateEmptyOrderId(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.Validate([]string{"SKU1"}, "tracking1", "")
	assert.Equal(t, e.Error(), "order id cant be nil")

}

func TestValidateInvalidOrderId(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.Validate([]string{"SKU1"}, "tracking1", "order100")
	assert.Equal(t, e.Error(), "order id order100 is invalid")

}

func TestValidateMoreSKUs(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.Validate([]string{"SKU1", "SKU2"}, "tracking1", "order1")
	assert.Equal(t, e.Error(), "SKU SKU2 is either not present or has more quantity than in original order")

}

func TestValidateMoreSKUThanRoomForUpdate(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.Validate([]string{"SKU1", "SKU1"}, "tracking1", "order1")
	assert.Equal(t, e.Error(), "SKU SKU1 is either not present or has more quantity than in original order")

}

func TestValidateSuccess(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	b, _ := s.Validate([]string{"SKU1"}, "tracking1", "order1")
	assert.Equal(t, b, true)

}

func TestShipmentUpdateSuccess(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	sku2 := domain.SKU{
		SKUId: "SKU2",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
			&sku2,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku1.SKUId}, "tracking1", "order1")
	if e != nil {
		t.Fatal(e)
	}
	sku, _ := s.GetShipment(o1.OrderID)
	assert.Equal(t, sku[0].ShipmentId, "tracking1")
}

func TestShipmentUpdateSuccessDifferentTimestamp(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	sku2 := domain.SKU{
		SKUId: "SKU2",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
			&sku2,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku1.SKUId}, "tracking1", "order1")
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku2.SKUId}, "tracking2", "order1")
	if e != nil {
		t.Fatal(e)
	}
	sku, _ := s.GetShipment(o1.OrderID)
	assertion := assert.Equal(t, sku[0].ShipmentId, "tracking1") &&
		assert.Equal(t, sku[1].ShipmentId, "tracking2")
	assert.Equal(t, assertion, true)
}

func TestShipmentUpdateSuccessFollowedByValidAndInvalid(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	sku2 := domain.SKU{
		SKUId: "SKU2",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
			&sku2,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku1.SKUId}, "tracking1", "order1")
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku2.SKUId, sku1.SKUId}, "tracking2", "order1")
	sku, _ := s.GetShipment(o1.OrderID)
	assertion := assert.Equal(t, e.Error(), "SKU SKU1 is either not present or has more quantity than in original order") &&
		assert.Equal(t, sku[0].ShipmentId, "tracking1")
	assert.Equal(t, assertion, true)
}

func TestShipmentUpdateSuccessFollowedByValidAndInvalid2(t *testing.T) {
	s := testSVC()
	sku1 := domain.SKU{
		SKUId: "SKU1",
	}
	sku2 := domain.SKU{
		SKUId: "SKU2",
	}
	o1 := domain.Order{
		OrderID: "order1",
		Skus: []*domain.SKU{
			&sku1,
			&sku2,
		},
	}
	_, e := s.AddOrder(o1)
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku1.SKUId}, "tracking1", "order1")
	if e != nil {
		t.Fatal(e)
	}
	_, e = s.ShipmentUpdate([]string{sku2.SKUId, "SKU3"}, "tracking2", "order1")
	sku, _ := s.GetShipment(o1.OrderID)
	assertion := assert.Equal(t, e.Error(), "SKU SKU3 is either not present or has more quantity than in original order") &&
		assert.Equal(t, sku[0].ShipmentId, "tracking1")
	assert.Equal(t, assertion, true)
}
