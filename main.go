package main

import (
	"log"
	"time"

	"github.com/jimjibone/lightwavego"
)

func main() {

	log.Println("LightwaveRF API")
	defer lightwavego.Shutdown()

	lightOn, _ := lightwavego.NewBuffer([]byte{0x9, 0xf, 0x3, 0x1, 0x5, 0x9, 0x3, 0x0, 0x1, 0x2})

	lwtx := lightwavego.NewLwTx()

	lwtx.SendBuffer(lightOn)
	time.Sleep(4 * time.Second)
}
