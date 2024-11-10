package models

// Item represents a single item on a receipt
type Item struct {
    // ShortDescription is a brief description of the item
    // The `json:"shortDescription"` tag ensures this field is serialized/deserialized as "shortDescription" in JSON
    ShortDescription string `json:"shortDescription"`

    // Price represents the item's price
    // The `json:",string"` tag tells Go to parse and output this float as a JSON string
    Price float64 `json:"price,string"`
}

// Receipt represents the structure of a full receipt
type Receipt struct {
    // Retailer is the name of the store or seller
    Retailer string `json:"retailer"`

    // PurchaseDate is the date of purchase in "YYYY-MM-DD" format
    PurchaseDate string `json:"purchaseDate"`

    // PurchaseTime is the time of purchase in "HH:MM" format
    PurchaseTime string `json:"purchaseTime"`

    // Items is a list of Item objects on the receipt
    Items []Item `json:"items"`

    // Total is the total cost of the receipt, parsed as a JSON string and stored as a float
    Total float64 `json:"total,string"`
}
