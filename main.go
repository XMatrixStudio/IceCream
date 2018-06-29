package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/XMatrixStudio/IceCream/generator"
	"github.com/XMatrixStudio/IceCream/httpserver"
	"gopkg.in/yaml.v2"
)

func main() {
	configFile := flag.String("c", "config/config.yaml", "Where is your config file?")
	toGenerate := flag.Bool("g", false, "generate dist")
	flag.Parse()
	if *toGenerate {
		generator.Generate("default")
		return
	}
	generator.Generate("default")
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Printf("Can't find the config file in %v", *configFile)
		return
	}
	log.Printf("Load the config file in %v", *configFile)
	conf := httpserver.Config{}
	yaml.Unmarshal(data, &conf)
	httpserver.RunServer(conf)
}
