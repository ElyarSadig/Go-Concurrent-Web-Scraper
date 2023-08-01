# Concurrent Web Scraper
This is a simple web scraper built in Go that retrieves product information from a website concurrently using Goroutines and channels. The scraper uses the popular gocolly library for HTML parsing and a concurrent design to speed up the data collection process.

## How It Works
The scraper collects product information from a [website](https://scrapeme.live/shop/) that lists pokemon products on multiple pages. It concurrently visits each page, extracts relevant product data, and stores it in a slice of PokemonProduct structs. The collected data is then written to a CSV file using another Goroutine, ensuring concurrent writing and efficient data handling.

## Installation
1. Make sure you have Go installed. If not, you can download and install it from the [official website](https://go.dev/).
2. Clone this repository to your local machine.

## Usage
1. Navigate to the project's root directory.
2. Open the terminal and run the following command to execute the scraper : ```go run scraper.go```
3. The scraper will start collecting product information concurrently from multiple pages of the website. After the scraper has finished running, you will find the collected data in a file named products.csv in the project directory.

## Dependencies
This project uses the gocolly library for web scraping. To install the dependency, you can run the following command:
```
go get -u github.com/gocolly/colly
```

## Contributions
Contributions to this project are welcome! If you have any suggestions, bug reports, or feature requests, feel free to open an issue or create a pull request.
