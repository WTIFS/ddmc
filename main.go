package main

import (
	"github.com/wtifs/ddmc/config"
	"github.com/wtifs/ddmc/service/cart"
)

func main() {
	cart.CheckCart(config.CartRawBody)
}
