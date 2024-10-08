import uuid

from pydanctic_models.receipt_process import ReceiptProcessCreate, ReceiptProcess



def new_receipt_process(receipt_process_req: ReceiptProcessCreate, receipt_process: ReceiptProcess):
    receipt_process.id = str(uuid.uuid4())
    receipt_process.retailer = receipt_process_req.retailer
    receipt_process.purchase_date = receipt_process_req.purchase_date
    receipt_process.purchase_time = receipt_process_req.purchase_time
    receipt_process.items = receipt_process_req.items
    receipt_process.total = receipt_process_req.total

    return receipt_process



