package main

import "github.com/go-vgo/robotgo"

func LunchSlave(){

}

func Control(){

}


func main() {
	robotgo.ScrollMouse(10, "up")
	robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
	robotgo.GetMousePos()
}