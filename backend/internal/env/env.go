// https://www.reddit.com/r/golang/comments/ygresd/configuration_in_microservices/
// https://mariocarrion.com/2018/01/31/go-environment-variables.html
// https://github.com/MarioCarrion/env-vars-example/blob/master/main.go

package env

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	CLERK_API_KEY string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var s Specification
	err = envconfig.Process("", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = fmt.Printf(s.CLERK_API_KEY)
	if err != nil {
		log.Fatal(err.Error())
	}
}
