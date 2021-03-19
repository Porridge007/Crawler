package main

import (
	"crawler/engine"
	"crawler/mockWeb/paser"
)

func main()  {
	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: paser.ParseCityList ,
	})
}
