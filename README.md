Simulator (task) of a backend server for processing and converting donations to Go

# Donate Processor Simulator (Go)

A simple and reliable backend server simulator for processing and converting incoming donations for the streaming platform.

## Project Description

The program simulates the flow of incoming donations in different currencies, performs strict data validation, and converts successful transactions into rubles (RUB). 

### Implemented functionality:
* Currency validation: Only `USD` and `EUR` are supported. Any other currencies are rejected by the security system.
* Security limits: Minimum donation is 1 unit, maximum one—time donation is 5,000 units. All transactions outside the limits are blocked.
* Error Handling: Validation errors are logged, but do not cause the server to crash — processing of the following donations continues normally (`continue').
* Guaranteed session closure: Using the `defer' mechanism to log session closure under any shutdown scenarios.

---

## Code Architecture

The logic is divided into independent layers:
1. `randomCurrency()` is a slice—based random currency generator for simulating real traffic.
2. `convertDonateToRUB()' is an isolated validation and conversion function.
3. `main()` is an entry point that simulates an infinite transaction processing cycle with an interval of 1 second.

---

## How to run a project locally

Make sure that you have [Go] installed(https://go.dev /).

1. Clone the repository:
``bash
git clone [https://github.com/pepesoul/go-donation-processor.git](https://github.com/pepesoul/go-donation-processor.git)
