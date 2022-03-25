package main

import (
	"github.com/drabadan/gostealthclient"
	"github.com/drabadan/internal/minercrafter"
)

func main() {
	gostealthclient.Bootstrap(minercrafter.Script)
}
