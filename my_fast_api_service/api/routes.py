from fastapi import APIRouter

from api.v1 import receipts
from api.v1 import receipt_points

api_router = APIRouter()

api_router.include_router(receipts.router, prefix="", tags=["receipts"])
api_router.include_router(receipt_points.router, prefix="", tags=["receipt_points"])