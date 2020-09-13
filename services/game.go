package services

import (
	"log"
	"net/http"
	"os"
	"strings"
	"tv-guide/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func GetGames() []models.Game {
	res, err := http.Get(os.Getenv("GAMES_LIST"))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	games := []models.Game{}
	doc.Find(".table-games tbody tr").Each(func(i int, s *goquery.Selection) {
		game := models.Game{}
		s.Find("td").Each(func(ix int, td *goquery.Selection) {
			switch ix {
			case 0:
				date := td.Text()
				if len(td.Find(".gametoday").Text()) > 0 {
					date = strings.Replace(date, " "+td.Find(".gametoday").Text(), "", 1)
				}

				game.Date = date
			case 1:
				td.Find(".team").Each(func(ixTeam int, team *goquery.Selection) {
					switch ixTeam {
					case 0:
						game.TeamA = team.Text()
					case 1:
						game.TeamB = team.Text()
					}
				})

				game.League = td.Find(".extra").Text()
			case 2:
				td.Find("a").Each(func(ixChannel int, ch *goquery.Selection) {
					switch ixChannel {
					case 0:
						game.Channel = strings.TrimSpace(ch.Text())
					}
				})
			}
		})

		if len(game.TeamA) > 0 {
			games = append(games, game)
		}
	})

	return games
}
