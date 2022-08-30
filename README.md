# Receipt Processor

Build a webservice that fulfils the documented API. The API is defined in `openapi.yml` file, but is summarized below.
We will use the described API to test your solution.

You can use any language and frameworks you prefer, but our engineers must be able to run your application to evaluate
it.

## Tips
* The engineer evaluating your solution may not have a dev environment for your language. Providing a docker file that can run your application helps with this.
* Data does not need to persist after a reboot. This allows you to use in-memory solutions. Databases can be difficult to set up.

---
## Process Receipt

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON

Receipts are processed by being posted to the `/receipts/process` endpoint. Processing receipts calculates the number of
points to be awarded for the receipt. The rules for how points are computed are described below. The response is just 
an ID assigned to the receipt.

The full Receipt schema is defined in `openapi.yml` and examples can be found in the example directory.

## Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`

Returns a JSON object containing the points earned for the receipt. If the ID does not match any receipts then a 404 
should be returned.

# Points Calculations

These rules define the ways a receipt can each points. These should be implemented for any receipt posted to the service.

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of `0.25`
* 5 points for every two items on the receipt.
* 9 points for every 3 items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the purchase is before noon.
* 15 points if the purchase is after noon.


## Example

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "Klarbrunn 12PK 12 FL OZ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

49 Points
* 6 points - retailer name has 6 characters
* 10 points - 4 items
* 9 points - 3 items.
* 3 Points "Emils Cheese Pizza" is 18 characters. `12.25 * 0.2 = 2.45` round up to 3.
* 6 points - for an odd purchase day
* 15 points - purchase was after noon.