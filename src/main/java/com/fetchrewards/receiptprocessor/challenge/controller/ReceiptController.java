package com.fetchrewards.receiptprocessor.challenge.controller;

import com.fetchrewards.receiptprocessor.challenge.model.Receipt;
import com.fetchrewards.receiptprocessor.challenge.model.ReceiptItem;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.server.ResponseStatusException;

import java.time.LocalTime;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

@RestController
@RequestMapping("/receipts")
public class ReceiptController {

    private final int POINTS_TOTAL_IS_WHOLE_DOLLAR = 50;
    private final int POINTS_TOTAL_IS_MULTIPLE_OF_QUARTERS = 25;
    private final int POINTS_EVERY_PAIR_OF_ITEMS = 5;
    private final int POINTS_PURCHASE_DATE_IS_ODD = 6;
    private final int POINTS_PURCHASE_TIME_IS_WITHIN_INTERVAL = 10;
    private final double POINTS_MULTIPLIER_ITEM_LENGTH_IS_MULTIPLE_OF_THREE = 0.2;
    private final LocalTime PURCHASE_TIME_INTERVAL_START = LocalTime.parse("14:00");
    private final LocalTime PURCHASE_TIME_INTERVAL_END = LocalTime.parse("16:00");

    private HashMap<String, Integer> receiptPoints = new HashMap<>();

    @GetMapping("{id}/points")
    public ResponseEntity<Object> getPoints(@PathVariable String id) {
        if (receiptPoints.containsKey(id)) {
            Map<String, String> data = new HashMap<>();
            data.put("points", receiptPoints.get(id).toString());
            return new ResponseEntity<>(data, HttpStatus.OK);
        }
        throw new ResponseStatusException(
                HttpStatus.NOT_FOUND, "Receipt ID not found."
        );
    }

    @PostMapping("/process")
    public ResponseEntity<Object> postReceipt(@RequestBody Receipt receipt) {
        // Here you can process the POST request data
        //return "Received POST data: " + data;
        System.out.println("received postttt");
        Map<String, String> data = new HashMap<>();
        data.put("id", processReceipt(receipt));
        return new ResponseEntity<Object>(data, HttpStatus.OK);

    }

    private String processReceipt(Receipt receipt) {
        String receiptId = UUID.randomUUID().toString();
        int points = calculatePoints(receipt);
        receiptPoints.put(receiptId, points);
        return receiptId;
    }

    private int calculatePoints(Receipt receipt) {
        int points = 0;
        points += getNumberOfAlnumChars(receipt.getRetailer());
        if (receipt.getTotal() % 1 == 0) points += POINTS_TOTAL_IS_WHOLE_DOLLAR;

        if (receipt.getTotal() % 0.25 == 0) points += POINTS_TOTAL_IS_MULTIPLE_OF_QUARTERS;

        points += (receipt.getItemsCount() / 2) * POINTS_EVERY_PAIR_OF_ITEMS;

        if (receipt.getPurchaseDate().getDayOfMonth() % 2 == 1) points += POINTS_PURCHASE_DATE_IS_ODD;

        if (receipt.getPurchaseTime().isAfter(PURCHASE_TIME_INTERVAL_START)
                && receipt.getPurchaseTime().isBefore(PURCHASE_TIME_INTERVAL_END))
            points += POINTS_PURCHASE_TIME_IS_WITHIN_INTERVAL;

        for (ReceiptItem item : receipt.getItems()) {
            if (item.getShortDescription().trim().length() % 3 == 0)
                points += (int) Math.ceil(item.getPrice() * POINTS_MULTIPLIER_ITEM_LENGTH_IS_MULTIPLE_OF_THREE);
        }

        return points;
    }

    private int getNumberOfAlnumChars(String retailer) {
        int count = 0;
        for (int i = 0; i < retailer.length(); i++) {
            char c = retailer.charAt(i);
            if (Character.isLetterOrDigit(c)) {
                count++;
            }
        }
        return count;
    }
}
