package main

import (
	"fmt"
	"time"
	"math/rand"
    "os"

	"github.com/go-vgo/robotgo"
)

var darcLabsX, darcLabsY int
var altSlackX, altSlackY int
var mainX, mainY int
var waitTime float64

func main() {
	fmt.Print("How long between clicks in min: ")
    fmt.Scanln(&waitTime)
	if (waitTime <= 1.0) {
		fmt.Println("Time interval too small. Please enter time > 1")
		os.Exit(int(waitTime))
	}

	fmt.Println("\nMove mouse TI Slack window")
	delay(4)
	darcLabsX, darcLabsY = robotgo.GetMousePos()
	fmt.Println("pos:", mainX, mainY)
	fmt.Println("\nMove mouse 757dev Slack window")
	delay(4)
	altSlackX, altSlackY = robotgo.GetMousePos()
	fmt.Println("pos:", mainX, mainY)
	autoClick()	
}

func autoClick() {
	count := 1
	iterations := 500
	mainX, mainY = robotgo.GetMousePos()
	robotgo.MoveMouse(mainX, mainY)
	for count <= iterations {
			fmt.Println("\n=========================")
		n := (rand.Float64()) // n will be between 0 and 10
		mainX, mainY = robotgo.GetMousePos()
		fmt.Println("working pos:", mainX, mainY)

		fmt.Println("Click TI Slack:", darcLabsX, darcLabsY)
		robotgo.MoveMouse(darcLabsX, darcLabsY)
		time.Sleep(time.Duration(n)*time.Second)
		robotgo.Click()
		fmt.Println("Click 757dev Slack:", darcLabsX, darcLabsY)
		robotgo.MoveMouse(altSlackX, altSlackY)
		time.Sleep(time.Duration(n)*time.Second)
		robotgo.Click()
		fmt.Println("Click Start Point:", mainX, mainY)
		robotgo.MoveMouse(mainX, mainY)
		robotgo.Click()
		time.Sleep(time.Duration(n)*time.Second)
		robotgo.Click()
		delay((waitTime*60)-30)
		fmt.Println("-- 30 second warning")
		count++
		fmt.Println("=========================\n")
	}

}

func delay(delayTime float64) {
    rand.Seed(time.Now().UnixNano())
	n := (rand.Float64()+0.7) // n will be between 0 and 10
	n=n+delayTime
	fmt.Println("Sleeping for seconds:", n)
	time.Sleep(time.Duration(n)*time.Second)
}
