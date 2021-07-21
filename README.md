# setup-quotel-db

This is a package for creating the Postgres database for the Quotes/Authors, found in the private repo https://github.com/Skjaldbaka17/Quotel-Data-JSON, that the Quotel-SLS-API uses https://github.com/Skjaldbaka17/quotel-sls-api/tree/main.

## Requirements

* [Golang](https://golang.org)
* The Quotel-Data in the directory above this one: https://github.com/Skjaldbaka17/Quotel-Data-JSON

## Setup

Clone the https://github.com/Skjaldbaka17/Quotel-Data-JSON repo into the parent directory of this project. 

Run the following to get the quotel data

```bash
    mkdir ../Quotel-Data-JSON
    git clone https://github.com/Skjaldbaka17/Quotel-Data-JSON ../Quotel-Data-JSON
```

Then to create the Postgres DB create a `.env` file with `DATABASE_URL=YOUR_DB_URL` and then run:

```bash
    go mod tidy
    make setup
```
