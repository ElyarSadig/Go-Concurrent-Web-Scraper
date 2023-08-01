package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/gocolly/colly"
)

const BASE = "https://scrapeme.live/shop/page/"

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func scrapePage(pageNum int, pokemonProducts *[]PokemonProduct, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	c := colly.NewCollector()

	pageToScrap := BASE + strconv.Itoa(pageNum)

	// scraping logic
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		m.Lock()
		*pokemonProducts = append(*pokemonProducts, pokemonProduct)
		m.Unlock()
	})

	c.Visit(pageToScrap)
}

func writeCSV(ch <-chan PokemonProduct) {
	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for pokemonProduct := range ch {
		// converting a PokemonProduct to an array of strings
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}

	writer.Flush()
}

func main() {
	// initializing the slice of structs to store the data to scrape
	pokemonProducts := []PokemonProduct{}
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(48)
	ch := make(chan PokemonProduct)

	for i := 1; i < 49; i++ {
		go scrapePage(i, &pokemonProducts, &wg, &m)
	}

	go writeCSV(ch)

	wg.Wait()

	for _, p := range pokemonProducts {
		ch <- p
	}

	wg.Wait()
	fmt.Println("CSV writing completed.")
}
