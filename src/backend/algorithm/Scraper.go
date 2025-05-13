package algorithm

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// var elementsCombination [][]string
// var elementsImage map[string]string
// var elementsTier map[string]int

func Scraper() ([][]string, map[string]string, map[string]int) {
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"

	// Fetch-nya berhasil ga bang
	response, e := http.Get(url)
	if e != nil {
		log.Fatalf("Failed to fetch URL: %v", e)
	}
	defer response.Body.Close()

	// Permintaan HTTP-nya berhasil ga bang
	if response.StatusCode != 200 {
		log.Fatalf("HTTP error: %d", response.StatusCode)
	}

	// Parse HTML-nya berhasil kah
	doc, e := goquery.NewDocumentFromReader(response.Body)
	if e != nil {
		log.Fatalf("Failed to parse HTML: %v", e)
	}

	elementsCombination := [][]string{
		{"Air", "-", "-"},
		{"Earth", "-", "-"},
		{"Water", "-", "-"},
		{"Fire", "-", "-"},
		{"Time", "-", "-"},
	}

	elementsImage := make(map[string]string)
	elementsTier := make(map[string]int)

	doc.Find("div.mw-content-ltr.mw-parser-output").Each(func(i int, tier *goquery.Selection) {

		currentTier := -1
		tier.Find("table.list-table.col-list.icon-hover").Each(func(j int, table *goquery.Selection) {
			// currentTier := tierList[currentTier]
			table.Find("tr").Each(func(k int, row *goquery.Selection) {
				if k == 0 {
					return // skip header
				}

				cols := row.Find("td")
				if cols.Length() != 2 {
					return
				}

				imageLink := ""
				row.Find("span span a").First().Each(func(_ int, a *goquery.Selection) {
					if href, exists := a.Attr("href"); exists {
						if idx := strings.Index(href, ".svg"); idx != -1 {
							imageLink = href[:idx+len(".svg")]
						}
					}
				})

				elementName := strings.TrimSpace(cols.Eq(0).Text())
				combinationCell := cols.Eq(1)

				elementsImage[elementName] = imageLink
				if currentTier >= 0 {
					elementsTier[elementName] = currentTier
				} else {
					elementsTier[elementName] = 0
				}

				var combinations []string
				if combinationCell.Find("li").Length() > 0 {
					combinationCell.Find("li").Each(func(_ int, li *goquery.Selection) {
						combinations = append(combinations, strings.TrimSpace(li.Text()))
					})
				} else {
					raw := strings.TrimSpace(combinationCell.Text())
					if raw != "" {
						combinations = append(combinations, raw)
					}
				}

				for _, combo := range combinations {
					parts := strings.Split(combo, "+")
					if len(parts) == 2 {
						ing1 := strings.TrimSpace(parts[0])
						ing2 := strings.TrimSpace(parts[1])
						if len(ing1) <= 20 {
							elementsCombination = append(elementsCombination, []string{elementName, ing1, ing2})
						}
					} else {
						ing := strings.TrimSpace(combo)
						if len(ing) <= 20 {
							elementsCombination = append(elementsCombination, []string{elementName, ing, ""})
						}
					}
				}
			})

			currentTier += 1
		})
	})
	return elementsCombination, elementsImage, elementsTier
}
