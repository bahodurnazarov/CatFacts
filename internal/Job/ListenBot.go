package Job

import (
	"fmt"
	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
	"io/ioutil"
	"net/http"
	"time"
)

func Route() {
	var timeDuration = 1
	timer := time.NewTicker(time.Second * time.Duration(timeDuration))
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			resp, err := http.Get("http://localhost:1323/bot")

			if err != nil {
				lg.Errl.Fatal(err)
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				lg.Errl.Fatal(err)
			}

			fmt.Println(string(body))
		}
	}
}
