from fastapi import APIRouter, status

from pydanctic_models.receipt_process import ReceiptProcessCreate, ShowReceiptProcess, ReceiptProcess
from api.helpers.receipt_process import new_receipt_process
from data_store import receipts_storage

router = APIRouter()


@router.post("/api/v1/receipts/process", response_model=ShowReceiptProcess, status_code=status.HTTP_201_CREATED)
def create_receipt_process(receipt_process: ReceiptProcessCreate):
    receipt_process = new_receipt_process(receipt_process, ReceiptProcess)

    receipts_storage.append({receipt_process.id: receipt_process})

    return receipt_process