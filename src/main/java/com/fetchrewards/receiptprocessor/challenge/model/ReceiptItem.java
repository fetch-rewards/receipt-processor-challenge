package com.fetchrewards.receiptprocessor.challenge.model;

public class ReceiptItem {
    private String shortDescription;
    private double price;

    public ReceiptItem(String shortDescription, String price) {
        this.shortDescription = shortDescription;
        this.price = Double.parseDouble(price);
    }

    public String getShortDescription() {
        return this.shortDescription;
    }

    public double getPrice() {
        return this.price;
    }
}
