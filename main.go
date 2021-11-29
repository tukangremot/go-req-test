package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

func main() {
	var url = os.Getenv("REQ_TEST_URL")

	if url == "" {
		fmt.Println("Please set the URL in your environment")
	} else {
		countries, err := getCountries(url)
		if err != nil {
			log.Fatal(err)
		}

		err = generateHtmlTable(countries)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getCountries(url string) ([]Country, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var countries []Country
	if err := json.Unmarshal([]byte(body), &countries); err != nil {
		return nil, err
	}

	return countries, nil
}

func generateHtmlTable(data []Country) error {
	var html = `<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>`
	fmt.Printf("<table>")
	for _, d := range data {
		fmt.Printf(html, d.Name, d.DialCode, d.IsoCode, d.Flag)
	}
	fmt.Printf("</table>")

	return nil
}
