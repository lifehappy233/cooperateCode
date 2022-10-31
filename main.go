package main

import (
	"time"

	"codeforces/handle"
	"codeforces/utils"
)

var process handle.Process

func Start() {
	for _, broadcast := range process.Broadcasts {
		time.Sleep(time.Second * 5)
		process.InitBroadcast(broadcast)
	}

	go process.ProcessMessage()

	for range time.Tick(15 * time.Second) {
		for _, broadcast := range process.Broadcasts {
			messages := process.Refresh(broadcast)
			for _, message := range messages {
				process.MessageChan <- message
			}
			time.Sleep(time.Second * 5)
		}
	}
}

func main() {
	process = handle.Process{
		Broadcasts:  make([]handle.Broadcast, 0),
		MessageChan: make(chan utils.GroupMessage, 20),
	}

	// test 719594145
	// work 770539963
	sigmaBroadcast := handle.NewSigmaBroadcast("770539963")
	sancppBroadcast := handle.NewSancppBroadcast("770539963")
	lifehappyBroadcast := handle.NewLifehappyBroadcast("770539963")
	process.Broadcasts = append(process.Broadcasts, &sigmaBroadcast)
	process.Broadcasts = append(process.Broadcasts, &sancppBroadcast)
	process.Broadcasts = append(process.Broadcasts, &lifehappyBroadcast)

	Start()
}
