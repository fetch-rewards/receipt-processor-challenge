import { Router, Request, Response } from 'express';
import Receipt from '../types/receiptType';
import ValidateRequest from '../middleware/Validate';

const router = Router();

//Submits a receipt for processing
router.post('/process', ValidateRequest.validateReceiptSubmission, (req: Request<object,object,Receipt>, res: Response) => {
  
});

//Returns the points awarded for the receipt
router.get('/{id}/points:', ValidateRequest.validateReceiptId, (req:Request, res:Response) => {

});

export default router;