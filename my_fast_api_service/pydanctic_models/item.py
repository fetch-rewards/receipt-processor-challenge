from pydantic import BaseModel, UUID4, StringConstraints, Field
from typing import Annotated

# class Item(BaseModel):
#     id: UUID4
#     short_description: Annotated[str, StringConstraints(pattern=r"^[\\w\\s\\-]+$")]
#     price: Annotated[str, StringConstraints(pattern=r"^\\d+\\.\\d{2}$")]

class ItemCreate(BaseModel):
    short_description: str = Field(
        pattern=r"^[A-Za-z0-9\s&,.]+$",
        description="A brief description of the receipt.",
        example="Grocery Store Visit"
    )
    price: str = Field(
        pattern = r"[+-]?([0-9]*[.])?[0-9]+",
        description="The price of the item, formatted as a string with two decimal places.",
        example="12.99"
    )
