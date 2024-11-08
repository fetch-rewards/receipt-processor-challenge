import { NextFunction, Request, Response } from 'express';
import Receipt from '../types/receiptType';

/*
Manual validation of Request bodies.
Validating just receipt request body data for now, can expand as more routes are defined.
*/

const ValidateRequest = {
  validateReceiptSubmission: (
    req: Request<object, object, Receipt>,
    res: Response,
    next: NextFunction
  ): void => { // Ensure return type is `void`
    const receipt = req.body;

    if (receipt) {
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
        return; // End the request-response cycle
      }

      // Validate purchase date
      if (isNaN(date)) {
        res.status(400).json({
          success: false,
          message: 'Invalid purchase date. It should be a valid date string.',
        });
        return; // End the request-response cycle
      }

      // Validate purchase time
      if (!timeRegex.test(purchaseTime)) {
        res.status(400).json({
          success: false,
          message: 'Invalid purchase time. It should be in the format HH:MM or HH:MM:SS.',
        });
        return; // End the request-response cycle
      }

      // Validate total amount
      if (!numberRegex.test(total)) {
        res.status(400).json({
          success: false,
          message: 'Invalid total amount. It should be a valid number with up to two decimal places.',
        });
        return; // End the request-response cycle
      }

      // Validate items
      if (typeof items !== 'object' || items.length < 1) {
        res.status(400).json({
          success: false,
          message: 'Items should be an array with at least one item.',
        });
        return; // End the request-response cycle
      }

      next();

    } else {
      res.status(400).json({
        success: false,
        message: 'Receipt data is missing.',
      });
      return; // End the request-response cycle
    }
  },

  validateReceiptId: (req: Request, res: Response, next: NextFunction): void => { // Ensure return type is `void`
    const { body } = req;

    if (typeof body === 'string') {
      if (body.length !== 36) {
        res.status(400).json({
          success: false,
          message: 'Invalid ID. ID must be 36 characters long.',
        });
        return; // End the request-response cycle
      }

      next();
      
    } else {
      res.status(400).json({
        success: false,
        message: 'Invalid ID. ID should be a string.',
      });
      return; // End the request-response cycle
    }
  },
};

export default ValidateRequest;
