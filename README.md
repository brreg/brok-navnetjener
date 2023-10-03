# BRØK Name Server API

## Overview

The BRØK Name Server API serves as the intersection of names, birth dates, and wallet addresses, facilitating seamless interactions between various systems. While public read-access is available but limited, the API allows designated systems to both read and write data.

## Capabilities

### Current Features:

* **New Wallet Creation**: Register a new wallet, including personal details like first name, last name, and social security number.
* **Wallet Query**: Retrieve names and birth dates associated with a particular wallet address.
* **Shareholder Info**: Obtain a list of all companies in which an individual or an organization holds shares.
* **Company Registry**: Browse through all registered companies on Brøk, with a display limit of 25 entities per page.
* **Company Query**: Retrieve details of a specific company by its organization number.
* **Share Count**: Find out the number of shares owned by an individual or organization in a specific company.
* **Shareholder Verification**: Input a list of individuals or organizations to check if they own shares in a specific company.
* **Quality Assurance Tests**: Run tests to validate the functionality of the system.

### Upcoming:

* **Authentication**: Possible integration with BR's API manager for secure authentication.
* **FM Person Integration**: Planned integration with the FM Person in the Name Server.

### Future Work:

* **Data Deletion**: Features to support the removal of data.
* **Data Updates**: Capability to update existing data records.
* **Protected Persons**: Support for confidential information as per FolkeRegister's codes 6 and 7.

## Authentication Policies

* **Write Access**: Restricted to authorized systems.
* **Read Access**: Publicly available but with limitations.

## Data Crawling Protection

* **Rate Limiting**: Potential implementation of throttling measures to control the volume of queries. This is particularly useful for entities requiring high query volumes.

## Development Setup

### Requirements:

* Docker
* Go 1.20+

### Instructions:

```bash
# Clone the repository
git clone git@github.com:brreg/brok-navnetjener.git
# Navigate to the directory
cd brok-navnetjener
# Download the necessary Go modules
go mod download
# Setup local environment
cp .env .env.local
# Start the Docker container
docker run --name navnetjener -d -e POSTGRES_PASSWORD="password" -v ./database/testdata.sql:/docker-entrypoint-initdb.d/testdata.sql -p 6666:6666 postgres -p 6666
```

### Run in Development Mode

* Use the command `go run .` to run the project in development mode.

### Testing

* Run tests using the `go test` command.

### Building

* Create a binary file using the `go build` command.

## Architecture

Please refer to the architecture diagrams to understand how the Name Server fits into the broader ecosystem.

![Diagram 1](https://github.com/brreg/brok-navnetjener/assets/18251869/4929baf9-35b6-4dea-b21c-77d57f185608)

![Diagram 2](https://github.com/brreg/brok-navnetjener/assets/877417/266b0aaa-81d1-4fa6-a1f3-a463f96bcca6)