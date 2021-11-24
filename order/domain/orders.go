package domain

// orders:[{
// 		orderid :"123"
// skus:[{
// skuid:"1234"
// shipmentid:"sdsd"
// },{
// skuid:"1234"
// shipmentid:"sdsd"
// }]
// }
// }]

type Order struct {
	OrderID string
	Skus    []*SKU
}

type SKU struct {
	SKUId      string
	ShipmentId string
}

type Service interface {
	AddOrder(Order) (string, error)
	ShipmentUpdate([]string, string, string) (string, error)
	GetShipment(string) ([]*SKU, error)
	Validate([]string, string, string) (bool, error)
}
