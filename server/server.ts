import express, {NextFunction, Request, Response } from 'express';
const app = express();
const PORT = 3000;
const IP = process.env.localIp ? process.env.localIp : '0.0.0.0';


//route to process receipts
app.post('/receipts/process', (req: Request, res: Response) => {
  
})

app.listen(PORT, IP, () => {
  console.log(`Server listening on port ${PORT}`);
});
