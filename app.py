from flask import Flask, request, jsonify
import re
import uuid
from datetime import datetime, time
import math

app = Flask(__name__)

receipts = {}

def check_condition1(receipt):
    retailer_name = receipt['retailer']
    points = sum(1 for char in retailer_name if char.isalnum())
    return points

def check_condition2(receipt):
    points = 0
    total = float(receipt['total'])
    if total.is_integer():
        points += 50
    return points

def check_condition3(receipt):
    points = 0
    total = float(receipt['total'])
    if total % 0.25 == 0:
        points += 25
    return points

def check_condition4(receipt):
    return (len(receipt['items']) // 2) * 5

def check_condition5(receipt):
    points = 0
    for item in receipt['items']:
        trimmed_len = len(item['shortDescription'].strip())
        if trimmed_len % 3 == 0:
            price = float(item['price'])
            points += math.ceil(price * 0.2)
    return points

def check_condition6(receipt):
    points = 0
    purchase_date = datetime.strptime(receipt['purchaseDate'], '%Y-%m-%d')
    if purchase_date.day % 2 == 1:
        points += 6
    return points

def check_condition7(receipt):
    points = 0
    purchase_time = datetime.strptime(receipt['purchaseTime'], '%H:%M').time()
    if time(14, 0) < purchase_time < time(16, 0):
        points += 10
    return points

def calculate_points(receipt):
    points = 0
    points += check_condition1(receipt)
    points += check_condition2(receipt)
    points += check_condition3(receipt)
    points += check_condition4(receipt)
    points += check_condition5(receipt)
    points += check_condition6(receipt)
    points += check_condition7(receipt)

    return points

@app.route('/receipts/process', methods=['POST'])
def process_receipts():
    receipt = request.get_json()

    # Validate receipt data
    if not is_valid_receipt(receipt):
        return jsonify({'error': 'Invalid receipt data'}), 400

    receipt_id = str(uuid.uuid4())

    receipts[receipt_id] = receipt

    response = {'id': receipt_id}

    return jsonify(response), 200

@app.route('/receipts/<string:receipt_id>/points', methods=['GET'])
def get_points(receipt_id):
    receipt = get_receipt(receipt_id)
    if receipt is None:
        return jsonify({'error': 'Receipt not found'}), 404

    points = calculate_points(receipt)
    response = {'points': points}

    return jsonify(response), 200

def is_valid_receipt(receipt):
    # Validate retailer name
    if not isinstance(receipt['retailer'], str):
        return False

    # Validate purchase date
    try:
        datetime.strptime(receipt['purchaseDate'], '%Y-%m-%d')
    except ValueError:
        return False

    # Validate purchase time
    try:
        datetime.strptime(receipt['purchaseTime'], '%H:%M')
    except ValueError:
        return False

    # Validate items
    if not isinstance(receipt['items'], list) or len(receipt['items']) < 1:
        return False

    for item in receipt['items']:
        if not is_valid_item(item):
            return False

    # Validate total
    if not isinstance(receipt['total'], str) or not bool(re.match(r'^\d+\.\d{2}$', receipt['total'])):
        return False

    return True

def is_valid_item(item):
    if not isinstance(item['shortDescription'], str) or not bool(re.match(r'^[\w\s\-]+$', item['shortDescription'])):
        return False

    if not isinstance(item['price'], str) or not bool(re.match(r'^\d+\.\d{2}$', item['price'])):
        return False

    return True

def get_receipt(receipt_id):
    return receipts.get(receipt_id)

def save_receipt(receipt_id, receipt):
    receipts[receipt_id] = receipt

if __name__ == '__main__':
    app.run(debug=True)
