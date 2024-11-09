import { Router, Request, Response } from 'express';
import Receipt from '../types/ReceiptType.js';
import ValidateRequest from '../middleware/Validate.js';
import ReceiptDataController from '../controllers/ReceiptDataController.js';
const router = Router();

//Submits a receipt for processing
router.post('/process', ValidateRequest.validateReceiptSubmission, ReceiptDataController.newReceipt);

//Returns the points awarded for the receipt
router.get('/:id/points', ValidateRequest.validateReceiptId, ReceiptDataController.getPoints);

export default router;