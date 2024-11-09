import { NextFunction, Request, Response } from 'express';
import Receipt from '../types/ReceiptType.js';
import ReceiptDataController from '../controllers/ReceiptDataController.js';
/*
Manual validation of Request bodies.
Validating just receipt request body data for now, can expand as more routes are defined.
*/

const ValidateRequest = {

  isReceipt: (object: any): object is Receipt =>
  {
    return (
      typeof object === 'object' &&
      Array.isArray(object.items) &&
      object.items.every(
        (item: any) =>
          typeof item === 'object' &&
          typeof item.shortDescription === 'string' &&
          typeof item.price === 'string'
      ) &&
      typeof object.purchaseDate === 'string' &&
      typeof object.purchaseTime === 'string' &&
      typeof object.retailer === 'string' &&
      typeof object.total === 'string'
    );
  },

  validateReceiptSubmission: (
    req: Request<object, object, Receipt>,
    res: Response,
    next: NextFunction
  ): void => {
    const receipt = req.body;
  
    if (ValidateRequest.isReceipt(receipt)) {
      const { items, purchaseDate, purchaseTime, retailer, total } = receipt;

      const date = Date.parse(purchaseDate);
      const timeRegex = /^([01]?[0-9]|2[0-3]):([0-5]?[0-9])(:([0-5]?[0-9]))?$/;
      const numberRegex = /^\d+(\.\d{2})?$/;

      // Validate retailer
      if (retailer.length < 1 || typeof retailer !== 'string') {
        res.status(400).json({
          success: false,
          message: 'Invalid retailer. It should be a non-empty string.',
        });
        return;
      }

      // Validate purchase date
      if (isNaN(date)) {
        res.status(400).json({
          success: false,
          message: 'Invalid purchase date. It should be a valid date string.',
        });
        return;
      }

      // Validate purchase time
      if (!timeRegex.test(purchaseTime)) {
        res.status(400).json({
          success: false,
          message: 'Invalid purchase time. It should be in the format HH:MM',
        });
        return;
      }

      // Validate total amount
      if (!numberRegex.test(total)) {
        res.status(400).json({
          success: false,
          message: 'Invalid total amount. It should be a valid number with up to two decimal places.',
        });
        return;
      }

      // Validate items list
      if (typeof items !== 'object' || items.length < 1) {
        res.status(400).json({
          success: false,
          message: 'Items should be an array with at least one item.',
        });
        return;
      }
      //validate item price
      else{
        for (const el of items) {
          if (!numberRegex.test(el.price)) {
            res.status(400).json({
              success: false,
              message: `Invalid item price: ${el.shortDescription}: $${el.price}. It should be a valid number with up to two decimal places.`,
            });
            return; // Exit the function after sending the response
          }
        }
      }

      next();

    } else {
      res.status(400).json({
        success: false,
        message: 'Invalid Receipt. Receipt data is missing.',
      });
      return;
    }
  },

  validateReceiptId: (req: Request, res: Response, next: NextFunction): void => {
    const id = req.params.id;
    if (typeof id === 'string') {
      if (id.length !== 36) {
        res.status(400).json({
          success: false,
          message: 'Invalid ID. ID must be 36 characters long.',
        });
        return;
      }
      else{
        if(!ReceiptDataController.storage.has(id)){
          res.status(404).json({
            success: false,
            message: 'ID Not Found!',
          });
          return;
        }
        next();
      }      
    } else {
      res.status(400).json({
        success: false,
        message: 'Invalid ID. ID should be a string.',
      });
      return;
    }
  },
};

export default ValidateRequest;
