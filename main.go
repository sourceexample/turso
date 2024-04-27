package main

import (
	"fmt"
	"testturso/modTurso"
)

const dburl = "libsql://test1-gitworkone.turso.io"
const dbToken = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MTQyMDA5MzEsImlkIjoiZDRiNTA4ZWItZGMyZi00ZTI4LWE1ZTAtOWM3MjE3NTc3MTVmIn0.d1C9K6ULRdJw2-Q-vNc9Ij_Thc46YO_vHx75D3KIdeqxMpmB-FEF-LXZh-MtrIfQZIfoxEW7RlQKjyRlFNjeDQ"

func main() {
	fmt.Println("Hello, World!")

	truso := modTurso.GetSingleTurso()

	err := truso.Connect(dburl, dbToken)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connect ok")
	}

	err = modTurso.GetSingleProject().Initialize()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Initialize ok")
	}

	inputhandler()
}
