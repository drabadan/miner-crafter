package minercrafter

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
)

func Script() interface{} {
	if <-sc.Connected() {
		for {
			if <-sc.Dead() {
				log.Println("Char is dead! Stopping...")
				break
			}

			sc.UseSkill("Hiding")
			time.Sleep(time.Second * 8)
		}
	}

	log.Println("Script finished!")
	return nil
}
