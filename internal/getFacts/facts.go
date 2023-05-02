package getfacts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
)

type scheme struct {
	Fact string `json:"fact"`
}

func GetFacts() string {
	url := "https://catfact.ninja/fact"
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

	scheme1 := scheme{}
	jsonErr := json.Unmarshal(body, &scheme1)
	if jsonErr != nil {
		lg.Errl.Fatal(jsonErr)
	}

	return scheme1.Fact
}
