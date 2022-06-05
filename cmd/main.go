package main

import (
	"github.com/drabadan/gostealthclient"
	multicharconnector "github.com/drabadan/multicharconnector/internal/script"
)

func main() {
	gostealthclient.Bootstrap(multicharconnector.Script)
}
