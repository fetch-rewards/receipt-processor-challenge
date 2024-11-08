
//Middleware to register receipts + generate points

import Receipt from "../types/receiptType"
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
    /*
    One point for every alphanumeric character in the retailer name.
    50 points if the total is a round dollar amount with no cents.
    25 points if the total is a multiple of 0.25.
    5 points for every two items on the receipt.
    If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
    6 points if the day in the purchase date is odd.
    10 points if the time of purchase is after 2:00pm and before 4:00pm.
    */

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

    //loop through items and add points based on item description
    

  },

  getPoints: (id: string) => {

  },

  getReceipt: (id: string) => {

  },

  deleteReceipt: (id: string) => {

  },


}

export default ReceiptDataController