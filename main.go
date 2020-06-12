package main

import (
	"errors"
	"ginframe/app/model"
	"ginframe/config"
	"ginframe/router"
	"github.com/spf13/pflag"
	"log"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)

func main()  {
	pflag.Parse()
	if err := config.Init(*conf); err != nil {
		panic(err)
	}

	isSuc := model.GetInstance().InitPool()
	if !isSuc {
		log.Println("init database pool failure...")
		panic(errors.New("init database pool failure"))
	}

	router.RunServer()
}
