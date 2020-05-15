package Commands

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Movie struct {
	title   string
	date    time.Time
	dateStr string
	page    string
}

const TrailerSite = "https://www.traileraddict.com"
const dateLayoutUS = "January 2, 2006"
const TrailerURI = "Trailer"

func Trailer(bot *tb.Bot, m *tb.Message) {
	var matchingStr, _ = regexp.Compile(`(?i)^\/trailer\s(.*)`)
	matches := matchingStr.FindStringSubmatch(m.Text)
	if len(matches) <= 1 {
		bot.Send(m.Chat, "Usage: /trailer movie title")
		return
	}
	movieTitle := matches[1]

	document := getPage(fmt.Sprintf("%s/search/%s", TrailerSite, url.QueryEscape(movieTitle)))

	var movies []Movie

	document.Find("#results>ul.result-list>li.result").Each(func(i int, s *goquery.Selection) {
		page, _ := s.Find(".title>a").Attr("href")
		dateStr := s.Find(".date").Text()
		t, _ := time.Parse(dateLayoutUS, dateStr)

		movies = append(movies, Movie{
			title:   s.Find(".title>a").Text(),
			date:    t,
			dateStr: t.Format(dateLayoutUS),
			page:    page,
		})
	})

	sort.SliceStable(movies, func(i, j int) bool {
		return movies[i].date.Unix() < movies[j].date.Unix()
	})

	var inlineKeys [][]tb.InlineButton

	out := fmt.Sprintf("Result search for '%s':\n", movieTitle)
	for _, m := range movies[:5] {
		if len(m.title) != 0 {
			inlineKeys = append(inlineKeys, []tb.InlineButton{
				tb.InlineButton{
					Unique: getUnique(slug.Make(m.page)),
					Text:   fmt.Sprintf("%s (%s)", m.title, m.dateStr),
				},
			})
		}
	}

	bot.Send(m.Chat, out, &tb.ReplyMarkup{InlineKeyboard: inlineKeys})
}

func TrailerOnCallback(c *tb.Callback) string {
	page := strings.Split(strings.TrimSpace(c.Data), ":")[1]

	document := getPage(fmt.Sprintf("%s/%s", TrailerSite, url.QueryEscape(page)))
	movieInfo := document.Find(".movie_info").Find("a.m_title").First()
	trailerPage, _ := movieInfo.Attr("href")

	document = getPage(fmt.Sprintf("%s/%s", TrailerSite, trailerPage))
	regVideoUrl, _ := regexp.Compile(`(?)src\=\"\/\/v.traileraddict.com/\d+`)
	docHtml, _ := document.Html()
	videoUrl := regVideoUrl.FindString(docHtml)

	urlVideo := fmt.Sprintf("https:%s", videoUrl[5:])

	return urlVideo

}

func getUnique(s string) string {
	return fmt.Sprintf("%s:%s", TrailerURI, s)
}

func getPage(url string) goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Filters\\Trailer: Error contacting Trailer website")
		return goquery.Document{}
	}
	defer resp.Body.Close()

	document, _ := goquery.NewDocumentFromReader(resp.Body)
	return *document
}
