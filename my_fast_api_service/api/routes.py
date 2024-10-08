from fastapi import APIRouter

from api.v1 import receipts

api_router = APIRouter()

api_router.include_router(receipts.router, prefix="", tags=["receipts"])