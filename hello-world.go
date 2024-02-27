package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	counter := 1
	for {
		randomWord := randomdata.SillyName() // Generate a random word.
		hash := sha256.Sum256([]byte(randomWord)) // Generate a SHA-256 hash of the random word.
		hashedWord := fmt.Sprintf("%x", hash) // Convert the hash to a hex string.

		log.Infof("hello world - iteration %d at %s - random word: %s - hash: %s", counter, time.Now().Format("2006-01-02 15:04:05"), randomWord, hashedWord)
		time.Sleep(5 * time.Second) // wait for 5 seconds before the next loop
		counter++
	}
}
