from pydantic import BaseModel, Field

class ItemCreate(BaseModel):
    short_description: str = Field(
        pattern=r"^[A-Za-z0-9\s&,.\\-]+$", example="Grocery Store Visit"
    )
    price: str = Field(pattern=r"[+-]?([0-9]*[.])?[0-9]+", example="12.99")
