package com.fetchrewards.receiptprocessor.challenge.model;

import java.time.LocalDate;
import java.time.LocalTime;
import java.util.List;

public class Receipt {
    private String retailer;
    private LocalDate purchaseDate;
    private LocalTime purchaseTime;
    private List<ReceiptItem> items;
    private double total;

    public Receipt(String retailer, String purchaseDate, String purchaseTime, String total, List<ReceiptItem> items) {
        this.retailer = retailer;
        this.purchaseDate = LocalDate.parse(purchaseDate);
        this.purchaseTime = LocalTime.parse(purchaseTime);
        this.total = Double.parseDouble(total);
        this.items = items;
    }

    public String getRetailer() {
        return this.retailer;
    }

    public LocalDate getPurchaseDate() {
        return this.purchaseDate;
    }

    public LocalTime getPurchaseTime() {
        return this.purchaseTime;
    }

    public double getTotal() {
        return this.total;
    }

    public int getItemsCount() {
        return this.items.size();
    }

    public List<ReceiptItem> getItems() {
        return this.items;
    }
}
