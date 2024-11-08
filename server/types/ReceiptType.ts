export default interface Receipt {
  retailer: string,
  purchaseDate: string,
  purchaseTime: string,
  total: string,
  items: [
    {
      shortDescription: string, 
      price: string
    }
  ]
}