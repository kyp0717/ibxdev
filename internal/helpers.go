package internal

import (
	"log"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
