package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("music.yandex.by"),
	)

	// authenticate
	err := c.Post("http://example.com/login", map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		log.Fatal(err)
	}

	fName := "music_list.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("couldn't create the file: %v", fName)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c.OnHTML(".d-track__overflowable-wrapper deco-typo-secondary block-layout", func(e *colly.HTMLElement) {

		writer.Write([]string{
			e.ChildText("title")
		})
	})
}

/*

<div class="d-track__overflowable-wrapper deco-typo-secondary block-layout">
	<div class="d-track__name" title="coordinate">
		<a href="/album/3780264/track/31195884" class="d-track__title deco-link deco-link_stronger"> coordinate </a>
			<span class="d-explicit-mark d-explicit-mark-e " title="Сервис Яндекс.Музыка может содержать информацию, не&nbsp;предназначенную для&nbsp;несовершеннолетних"></span></div>
	<div class="d-track__meta">
		<span class="d-track__artists">
			<a href="/artist/999165" title="Travis Scott" class="deco-link deco-link_muted">Travis Scott</a>
		</span>
    </div>
</div>

*/
