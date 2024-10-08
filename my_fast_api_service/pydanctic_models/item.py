from pydantic import BaseModel, UUID4, StringConstraints, Field
from typing import Annotated


class ItemCreate(BaseModel):
    short_description: str = Field(
        pattern=r"^[A-Za-z0-9\s&,.\\-]+$",
        example="Grocery Store Visit"
    )
    price: str = Field(
        pattern = r"[+-]?([0-9]*[.])?[0-9]+",
        example="12.99"
    )
