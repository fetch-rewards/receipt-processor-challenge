from pydantic import BaseModel, field_validator, Field, UUID4
from typing import List
from datetime import date, time

from pydanctic_models.item import ItemCreate


class ReceiptProcess(BaseModel):
    id: UUID4
    retailer: str = Field(pattern=r"^[A-Za-z0-9\s&.'-]+$")
    purchase_date: date
    purchase_time: time
    items: List[ItemCreate] = Field(..., min_items=1)
    total: str = Field(pattern=r"[+-]?([0-9]*[.])?[0-9]+")


class ReceiptProcessCreate(BaseModel):
    retailer: str = Field(pattern=r"^[A-Za-z0-9\s&.'-]+$", example="M&M Corner Market")
    purchase_date: date
    purchase_time: time = Field(example="13:01")
    items: List[ItemCreate] = Field(..., min_items=1)
    total: str = Field(pattern=r"[+-]?([0-9]*[.])?[0-9]+", example="6.49")

    @field_validator("items")
    def validate_items(cls, v):
        if len(v) < 1:
            raise ValueError("At least one item must be provided.")

        return v


class ShowReceiptProcess(BaseModel):
    id: UUID4

    class Config:
        orm_mode = True
