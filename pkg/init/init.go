package init

import (
	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
	"github.com/joho/godotenv"
)

func Init() {

	err := godotenv.Load("../../.env")

	if err != nil {
		lg.Errl.Fatal("Error loading .env file")
	}
}
