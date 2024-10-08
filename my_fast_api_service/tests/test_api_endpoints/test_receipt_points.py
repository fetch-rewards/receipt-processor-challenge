from tests.test_config import client

def test_receipt_points():
    payload = {
        "retailer": "Target",
        "purchase_date": "2022-01-01",
        "purchase_time": "13:01",
        "items": [
            {
            "short_description": "Mountain Dew 12PK",
            "price": "6.49"
            },{
            "short_description": "Emils Cheese Pizza",
            "price": "12.25"
            },{
            "short_description": "Knorr Creamy Chicken",
            "price": "1.26"
            },{
            "short_description": "Doritos Nacho Cheese",
            "price": "3.35"
            },{
            "short_description": "   Klarbrunn 12-PK 12 FL OZ  ",
            "price": "12.00"
            }
        ],
        "total": "35.35"
    }

    resp = client.post('/api/v1/receipts/process', json=payload)

    receipt_id = resp.json()["id"]

    resp_two = client.get(f'/api/v1/receipts/{receipt_id}/points')

    assert resp_two.status_code == 200
    assert resp_two.json()["points"] == 28