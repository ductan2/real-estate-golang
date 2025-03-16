package main

import (
	"ecommerce/internal/initialize"

)

func main(){
	 
	engine := initialize.Run()
	engine.Run(":8080")
}