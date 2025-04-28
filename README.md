Payment Tracking API - README

Project Name: Payment Tracking API

## Introduction

This project is a simple Payment Tracking API built using Golang.
It allows you to track the status of payments, retry payments, and create new payment records. The
project is intended for learning and development purposes and demonstrates basic CRUD
operations with an in-memory store.

## Folder Structure

- **handlers**: Contains logic for handling HTTP requests (Create, Get, Retry).

- **models**: Contains data models (Payment, PaymentStatus).

- **store**: Manages in-memory data (stores payments).

- **config**: Holds configuration information (but not used in this project directly).

- **main.go**: Entry point of the application, where routes are registered.

- **tests**: Folder for unit tests (not implemented yet).

- **README.pdf**: This file.

## Endpoints

1. **POST /payment**:

 - **Function**: Creates a new payment record.
 - **Business Logic**: A new payment is created with the status 'pending' by default. This allows
businesses to track payment statuses and manage payments in a simple in-memory store.

2. **GET /payment/{id}**:

 - **Function**: Retrieves the status of a specific payment using its ID.
 - **Business Logic**: This allows users to check the current status of their payment. For example,
if payment is pending, completed, or failed.

3. **POST /payment/retry/{id}**:

 - **Function**: Retries a failed payment based on the payment ID.
 - **Business Logic**: This is useful for situations where a payment fails due to network issues or
other temporary errors. The status is updated to 'retrying' before attempting the payment again.

## Business Logic

- **Create Payment**: When a payment is created, it is initially marked as 'pending'. This mimics the
real-world behavior of waiting for payment confirmation.

- **Get Payment Status**: This allows the user to check the current status of their payment. In
real-world scenarios, this might be a result of various payment gateway responses (e.g., success,
failure, pending).

- **Retry Payment**: Sometimes, payments fail due to external issues (e.g., network error,
insufficient funds). This endpoint provides functionality to retry the payment, updating the status
accordingly.

## How to Run
1. Clone the repository:
 ```
 git clone https://github.com/your-username/payment-tracking-api.git
 cd payment-tracking-api
 ```
2. Install dependencies (if any):
 This project only uses the standard Golang libraries.
3. Run the application:
 ```
 go run main.go
 ```
## Conclusion
This project serves as a basic API for tracking and retrying payments. It is a good starting point for
learning about APIs and can be extended with more advanced features for production
environments.
