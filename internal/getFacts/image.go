package getfacts

import (
	"encoding/json"
	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
	"io/ioutil"
	"net/http"
)

type ImgStr struct {
	Url string `json:"url"`
}

func GetImage() string {
	url := "https://api.thecatapi.com/v1/images/search"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		lg.Errl.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		lg.Errl.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		lg.Errl.Println(err)

	}

	var data []ImgStr
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		lg.Errl.Fatal(jsonErr)
	}
	var sterr string
	for _, values := range data {
		sterr = values.Url
	}
	return sterr
}
