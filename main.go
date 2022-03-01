package main

import (
	"sync"

	//covergencia "github.com/daniielsantos/channels/convergencia"
	divergencia "github.com/daniielsantos/channels/divergencia"
)

var wg sync.WaitGroup

func main() {
	//covergencia.GetConvergencia()
	divergencia.GetDivergencia()
}
