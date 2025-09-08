# Domain Email Security Checker

This is a simple Go application that checks if a given email domain has essential DNS records for email security.  
It verifies the following:

- **Valid Domain**
- **MX Records**
- **SPF Records**
- **SPF Record Content**
- **DMARC Records**
- **DMARC Record Content**

---

## Features
- Takes an email domain as input from the terminal.
- Checks DNS for **MX**, **SPF**, and **DMARC** records.
- Displays the results in a clean, readable format.

Example output:
Enter email domain and we will check if it has a:
valid domain
Mx records
SPF records
sprRecord
DMARC
dmarcRecord

Domain: example.com
Has MX: true
Has SPF: true
SPF Record: v=spf1 include:_spf.example.com ~all
Has DMARC: true
DMARC Record: v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com


---

## Installation

1. Make sure you have [Go installed](https://go.dev/doc/install) (version 1.18 or higher recommended).
2. Clone this repository:
git clone https://github.com/Precious-Neche/domain-verifier.git
cd domain-verifier
3. run the project
go run main.go

