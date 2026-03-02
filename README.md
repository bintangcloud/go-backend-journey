# Day 1 - Golang Backend 🚀  
## Nested JSON: Order, Customer, and Items (Marshal & Unmarshal)

Today I learned how to create nested JSON structures in Golang using struct and slice, and also how to **marshal** and **unmarshal** JSON.

---

## 📌 Problem

Create a struct `Order` that contains:
- A nested struct `Customer`
- A slice (list) of `Item`
- Convert the struct into JSON (marshal)
- Read JSON back into struct (unmarshal)

---

## 🏗️ Data Structure Design

Order
- ID
- Customer
  - Name
  - Email
- Items (list)
  - Name
  - Price
  - Quantity

---
Ni Putu Bintang Permatasari — Computer Science Student
