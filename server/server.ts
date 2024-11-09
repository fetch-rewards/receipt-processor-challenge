import express, {Request, Response, NextFunction} from 'express';
import router from './routes/ReceiptRoutes.js';
const app = express();
const PORT = process.env.PORT ? process.env.PORT : 3000;
const IP = process.env.localIp ? process.env.localIp : '';

// Parse JSON Bodies
app.use(express.json()); 

// Route Receipts
app.use('/receipts', router)

// Global Error Handler
// app.use((err:Error, req:Request, res:Response, next:NextFunction) => {
//   // Log the error (useful for debugging)
//   console.error('Error:', err.message);

//   // Set the HTTP status code from the error or default to 500
//   const statusCode = err.statusCode || 500;

//   // Send a JSON response with error details
//   res.status(statusCode).json({
//     success: false,
//     message: err.message || 'Internal Server Error',
//   });
// });


app.listen(PORT, IP, () => {
  console.log(`Server listening on port ${IP}:${PORT}`);
});
