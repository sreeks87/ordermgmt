### Problem Statement

An online order contains line items. Each line item is identified by a product SKU and its quantity.

E.g. Order-1 has products SKU1, SKU1, SKU2 (repeating SKUs means same products in more than one quantity in an order). The line items would be:

SKU1, quantity 2

SKU2, quantity 1

Products inside an order can be shipped in more than one shipment. A shipment is uniquely identified with a tracking number.

E.g. Order-1 with the above SKUs could be shipped in two shipments:

Shipment with tracking tracking-1 ships SKU1, SKU2

Shipment with tracking tracking-2 ships SKU1

Updates for a Shipment can arrive in multiple parts. For example:

# update 1 at timestamp-1

Order-1 got SKU1 shipped with tracking-1
# update 2 at timestamp-2

Order-1 got SKU1 shipped with tracking-2
# update 3 at timestamp-3

Order-1 got SKU2 shipped with tracking-1

One shipment update contains all or part of SKUs for one tracking code. (i.e assume one tracking per shipment update)

After the above three updates, all SKUs in Order-1 have shipment information. These will be the shipments for Order-1:

Shipment with tracking-1 shipped SKU1, SKU2

Shipment with tracking-2 shipped SKU1

Shipment information can be repeated in a later update. For example, consider this 
alternative version of update 3:

# update 3, alternative 1 at timestamp-3

Order-1 got SKU1 shipped with tracking-1

Order-1 got SKU2 shipped with tracking-1

After update 2 only one quantity of SKU2 needs an update, whereas the update 
provides for SKU1 also. 

When such a shipment update comes with shipment information for more SKUs than there is room for to update, such updates should not be applied.

Therefore following alternatives cannot be accepted either:

# update 3 alternative2 at timestamp-3

Order-1 got SKU1 shipped with tracking-2

There are no SKU1s left that need shipment update

Order-1 got SKU2 shipped with tracking-2

This part is valid, but since the update itself has an invalid part, the update needs to be rejected

# update 3 alternative 3 at timestamp-3

Order-1 got SKU3 shipped with tracking-1
There is no such SKU 'SKU3' in the original order, so this update is invalid.
Programming Task

Write a service that allows the following operations:

add an order with line items

accept shipment updates for an order

return current known shipments for an order

cover solution with unit tests

* no interaction with network or database is needed. A solution that keeps everything in memory is sufficient.


# To run locally

This module is built in `go 1.16`

Install go

Download dependencies with `go mod download`

Modify the main.go to cover any use case from the readme file above.

Run `go run .`

## Tests

    === RUN   TestAddOrderSuccess
    --- PASS: TestAddOrderSuccess (0.00s)
    === RUN   TestAddOrderNoOrderID
    --- PASS: TestAddOrderNoOrderID (0.00s)
    === RUN   TestGetshipmentSuccess
    --- PASS: TestGetshipmentSuccess (0.00s)
    === RUN   TestGetShipmentWithoutOrderID
    --- PASS: TestGetShipmentWithoutOrderID (0.00s)
    === RUN   TestGetShipmentInvalidOrderID
    --- PASS: TestGetShipmentInvalidOrderID (0.00s)
    === RUN   TestValidateEmptySKUList
    --- PASS: TestValidateEmptySKUList (0.00s)
    === RUN   TestValidateEmptyShipTrackID
    --- PASS: TestValidateEmptyShipTrackID (0.00s)
    === RUN   TestValidateEmptyOrderId
    --- PASS: TestValidateEmptyOrderId (0.00s)
    === RUN   TestValidateInvalidOrderId
    --- PASS: TestValidateInvalidOrderId (0.00s)
    === RUN   TestValidateMoreSKUs
    --- PASS: TestValidateMoreSKUs (0.00s)
    === RUN   TestValidateMoreSKUThanRoomForUpdate
    --- PASS: TestValidateMoreSKUThanRoomForUpdate (0.00s)
    === RUN   TestValidateSuccess
    --- PASS: TestValidateSuccess (0.00s)
    === RUN   TestShipmentUpdateSuccess
    --- PASS: TestShipmentUpdateSuccess (0.00s)
    === RUN   TestShipmentUpdateSuccessDifferentTimestamp
    --- PASS: TestShipmentUpdateSuccessDifferentTimestamp (0.00s)
    === RUN   TestShipmentUpdateSuccessFollowedByValidAndInvalid
    --- PASS: TestShipmentUpdateSuccessFollowedByValidAndInvalid (0.00s)
    === RUN   TestShipmentUpdateSuccessFollowedByValidAndInvalid2
    --- PASS: TestShipmentUpdateSuccessFollowedByValidAndInvalid2 (0.00s)
    PASS
    coverage: 100.0% of statements in .
    ok      github.com/sreeks87/ordermgmt/order/tests       0.345s  coverage: 100.0% of statements in .