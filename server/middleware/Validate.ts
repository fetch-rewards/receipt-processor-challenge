import { NextFunction, Request, Response } from 'express';
import Receipt from '../types/receiptType';


/*
Manual validation of Request bodies
Validating just receipt request body data for now, can expand as more routes are defined
*/

const ValidateRequest = {
  validateReceiptSubmission: (req:Request<object,object,Receipt>, res:Response, next:NextFunction) => {
    const receipt = req.body;
    if(receipt){
      const { items, purchaseDate, purchaseTime, retailer, total } = receipt;
      const date = Date.parse(purchaseDate);
      const timeRegex = /^([01]?[0-9]|2[0-3]):([0-5]?[0-9])(:([0-5]?[0-9]))?$/;
      const numberRegex = /^\d+(\.\d{2})?$/;
      if(retailer.length < 1 || typeof retailer !== 'string'){
        return res.status(400).json({
          success: false,
          message: 'Invalid retailer. It should be a non-empty string.'
        });      
      }
      else if(isNaN(date)){
        return res.status(400).json({
          success: false,
          message: 'Invalid purchase date. It should be a valid date string.'
        });
      }
      else if(!timeRegex.test(purchaseTime)){
        return res.status(400).json({
          success: false,
          message: 'Invalid purchase time. It should be in the format HH:MM or HH:MM:SS.'
        });    
      }
      else if(!numberRegex.test(total)){
        return res.status(400).json({
          success: false,
          message: 'Invalid total amount. It should be a valid number with up to two decimal places.'
        });  
      }
      else if(typeof items !== 'object' || items.length < 1){
        return res.status(400).json({
          success: false,
          message: 'Items should be an array with at least one item.'
        });
      }
      else{
        next();
      }
    }
  },  
  validateReceiptId: (req:Request, res:Response, next:NextFunction) => {
    if(typeof req.body == 'string'){
      if(req.body.length !== 36){
        return res.status(400).json({
          success: false,
          message: 'Invalid ID'
        });
      }
      else{
        next();
      }
    }
    else{
      return res.status(400).json({
        success: false,
        message: 'Invalid ID'
      });
    }
  }
}

export default ValidateRequest