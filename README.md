# Receipt Processor

Build a webservice that fulfils the documented API. The API is described below. A formal definition is provided 
in `api.yml` file, but the information in this readme is sufficient for completion of this challenge. We will use the 
described API to test your solution.

Provide any instructions required to run your application.

Data does not need to persist when your application stops. It is sufficient to store information in memory. There are too many different database solutions, we won't install a database on our system to test your application.

## Language Selection

You can assume our engineers have Go and Docker installed to run your application. Go is our preferred language, but it is not a requirement for this exercise.

If you are using a language other than Go, the engineer evaluating your submission may not have an environment ready for your language. Your instructions should include how to get an environment in any OS that can run your project. For example, if you write your project in Javascript simply stating to "run `npm start` to start the application" is not sufficient, because the engineer may not have NPM. Providing a docker file and the required docker command is a simple way to satisfy this requirement.

---
## API

### Endpoint: Process Receipts

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.

Description:

Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID specifying the id.

The ID returned is the ID that should be passed into `/receipts/{id}/points` to get the number of points the receipt
was awarded.

How many points should be earned are defined by the rules below.

Example Response:
```json
{"id": "7fb1377b-b223-49d9-a31a-5a02701dd310"}
```

## Endpoint: Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.

A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

Example Response:
```json
{ "points": 32 }
```

# Rules

These rules collectively define how many points should be awarded to a receipt.

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of `0.25`
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm


## Examples

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

25 Points
* 6 points: Retailer name has 6 characters
* 10 points: 4 items (2 pairs @ 5 points each)
* 3 Points: "Emils Cheese Pizza" is 18 characters. `12.25 * 0.2 = 2.45` round up to 3.
* 6 points: Purchase day is odd

----

```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```

109 Points

* 50 Points: Total is a round dollar amount
* 25 Points Total is a multiple of `0.25`
* 14 Points: Retailer Name has 14 Alpha-Numeric characters (`&` is not alpha-numeric)
* 10 Points: 2:33pm is between 2:00pm and 4:00pm: 10 points
* 10 Points: 4 items (2 pairs @ 5 points each)