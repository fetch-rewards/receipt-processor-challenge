from flask import Flask, request, jsonify
import re
import uuid
from datetime import datetime, time
import math

app = Flask(__name__)

receipts = {}

class Scorer:

    @classmethod
    def compute_score(cls, receipt):
        points = 0
        points += cls._alpha_numeric_score(receipt)
        points += cls._is_integer_total_score(receipt)
        points += cls._is_total_divisible_by_quarter_score(receipt)
        points += cls._num_items_score(receipt)
        points += cls._is_trimmed_length_divisible_by_3_score(receipt)
        points += cls._is_day_odd_score(receipt)
        points += cls._is_time_between_2_and_4_score(receipt)
        return points

    @staticmethod
    def _alpha_numeric_score(receipt):
        retailer_name = receipt["retailer"]
        return sum([1 for char in retailer_name if char.isalnum()])

    @staticmethod
    def _is_integer_total_score(receipt):
        total = float(receipt["total"])
        return 50 if total.is_integer() else 0
    
    @staticmethod
    def _is_total_divisible_by_quarter_score(receipt):
        total = float(receipt["total"])
        return 25 if total % 0.25 == 0 else 0

    @staticmethod
    def _num_items_score(receipt):
        return (len(receipt['items']) // 2) * 5

    @staticmethod
    def _is_trimmed_length_divisible_by_3_score(receipt):
        points = 0
        for item in receipt['items']:
            trimmed_len = len(item['shortDescription'].strip())
            if trimmed_len % 3 == 0:
                price = float(item['price'])
                points += math.ceil(price * 0.2)
        return points

    @staticmethod
    def _is_day_odd_score(receipt):
        purchase_date = datetime.strptime(receipt["purchaseDate"], "%Y-%m-%d")
        return 6 if purchase_date.day % 2 == 1 else 0
    
    @staticmethod
    def _is_time_between_2_and_4_score(receipt):
        purchase_time = datetime.strptime(receipt['purchaseTime'], '%H:%M').time()
        return 10 if time(14) < purchase_time < time(16) else 0

@app.route('/receipts/process', methods=['POST'])
def process_receipts():
    #Get json object, generate id and store in dictionary
    receipt = request.get_json()
    receipt_id = str(uuid.uuid4())
    receipts[receipt_id] = receipt
    response = {'id': receipt_id}

    return jsonify(response), 200

@app.route('/receipts/<string:receipt_id>/points', methods=['GET'])
def get_points(receipt_id):
    receipt = get_receipt(receipt_id)
    if receipt is None:
        return jsonify({'error': 'Receipt not found'}), 404

    #Compute score for the recipt id
    points = Scorer.compute_score(receipt)
    response = {'points': points}

    return jsonify(response), 200

def get_receipt(receipt_id):
    return receipts.get(receipt_id)

def save_receipt(receipt_id, receipt):
    receipts[receipt_id] = receipt

if __name__ == '__main__':
    app.run(debug=True)
