
//Middleware to register receipts + generate points

import Receipt from "../types/ReceiptType.js";
import { Request, Response, NextFunction } from "express"

/*
Map stores receipt in this format
id: {receipt: receiptData, points: pointsAwarded}
*/
interface ReceiptData {
  receipt: Receipt,
  points: number
}

const ReceiptDataController = {
  storage: new Map<string, ReceiptData>(),
  newReceipt: (req:Request, res: Response, next: NextFunction) => {

    //create id for receipt
    const receipt = req.body;
    const receiptId = crypto.randomUUID();
    ReceiptDataController.storage.set(receiptId, {receipt: receipt, points: ReceiptDataController.generatePoints(receipt)})
    res.status(200).json({id: receiptId})

  },

  generatePoints: (receipt: Receipt) => {
    let total = 0;
    const receiptTotal = parseFloat(receipt.total);


    //remove all non alphanumeric from retailer name, add length to total points
    const cleanedRetailerName = receipt.retailer.replace(/[^a-zA-Z0-9]/g, '');


    total += cleanedRetailerName.length;
    //check total amount if round number
    if(receiptTotal == Math.floor(receiptTotal)){
      total += 50;
    }


    //check total amount if divisible by 0.25
    if(receiptTotal % 0.25 == 0){
      total += 25;
    }


    //add points for every two items on receipt
    const pairs = Math.floor(receipt.items.length / 2);
    total += (pairs * 5);


    //add points based on day and hour
    const [hours, minutes] = receipt.purchaseTime.split(":").map(Number);
    const purchaseDate = new Date(receipt.purchaseDate);
    const purchaseDay = purchaseDate.getDay();
    const purchaseHour = hours;
    if(purchaseDay % 2 !== 0){
      total+=6
    }
    if(purchaseHour >= 14 && purchaseHour < 16){
      total+=10
    }


    //loop through items and add points based on item description
    receipt.items.forEach(el => {
      const multipleOfThree = (el.shortDescription.trim().length % 3);
      if(multipleOfThree == 0){
        //convert price from string to float, set decimals to hundreds place, convert outputted string to float again
        const points = Math.ceil(parseFloat(el.price) * 0.2);
        total += points;
      }
    })

    return total
  },

  //get points from storage
  getPoints: (req:Request, res:Response, next: NextFunction) => {
    
    const id = req.params.id;
    const points = ReceiptDataController.storage.get(id)!.points;
    res.status(200).json({points:points})

  }

}

export default ReceiptDataController