//PRESS ^C TO STOP EXECUTING!!!

package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func Zzleeping(lag int) fmt.Stringer {

	lag = rand.Intn(lag)

	time.Sleep(time.Duration(lag) * time.Microsecond)

	return (time.Duration(lag) * time.Microsecond)

}

func logging(uid int, got, sent string, timer fmt.Stringer) {

	log.WithFields(log.Fields{
		"uid":       uid,
		"sent":      sent,
		"got":       got,
		"game time": timer,
	}).Info("NEW ROUND")

}

func pingpong(str string, ch chan string, uid int) {

	logging(uid, "PING", "PONG", Zzleeping(1000000))

	if str == "PING" || str == "START GAME" {
		ch <- "PONG"
	}

	if str == "PONG" {
		ch <- "PING"
	}

}

func main() {

	var chFirst chan string = make(chan string)
	var chSecond chan string = make(chan string)

	go pingpong("START GAME", chFirst, 1)

	for {
		go pingpong(<-chFirst, chSecond, 2)
		go pingpong(<-chSecond, chFirst, 1)
	}

	fmt.Scanln()

}
