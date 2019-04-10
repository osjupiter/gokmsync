package main

/*
#cgo LDFLAGS: -lm -lX11
#include <X11/Xlib.h>
#include <X11/X.h>
#include <X11/Xutil.h>
#include <stdlib.h>
#include <math.h>
#include <stdio.h>
typedef struct {
	int win_x,win_y,root_x,root_y;
}Positions;
Positions p;

Positions getMouse(Display *display,Window root){

    // Get the mouse cursor position
    //int win_x, win_y, root_x, root_y = 0;
    unsigned int mask = 0;
    Window child_win, root_win;
    XQueryPointer(display, root, &child_win,&root_win,&p.root_x, &p.root_y, &p.win_x, &p.win_y, &mask);
    return p;
}
*/
import "C"

import (
	"fmt"
	"time"
)

const (
	Thereshold = 5
)
var fin chan bool

func LunchMaster() {
	fin = make(chan bool)

	go func() {
		for {
			select {
			case <-fin:
				break

			}
		}

	}()

}

func main() {
	dpy:=C.XOpenDisplay(nil)
	root:=C.XDefaultRootWindow(dpy)
	screen:=C.XDefaultScreenOfDisplay(dpy)
	width:=int(C.XWidthOfScreen(screen))
	height:=int(C.XHeightOfScreen(screen))
	fmt.Printf("%d %d\n", width,height)
	for {
		time.Sleep(100*time.Millisecond)
		a := C.getMouse(dpy,root)
		fmt.Printf("%v\n", a)
		var px int= int(a.win_x)
		var py int= int(a.win_y)
		if px<=0+Thereshold|| px>=width-Thereshold-1||py<=0+Thereshold||py>=height-1-Thereshold{
			fmt.Printf("touched!!!\n")
		}
	}


	//	for{
	//			C.XNextEvent(d, &e);
	//			if (e.type == Expose) {
	//				XFillRectangle(d, w, DefaultGC(d, s), 20, 20, 10, 10);
	//				XDrawString(d, w, DefaultGC(d, s), 10, 50, msg, strlen(msg));
	//			}
	//			if (e.type == ButtonPress){
	//				switch (e.xbutton.button){
	//				case Button4:
	//					printf("Scrolled up");
	//					break;
	//				case Button5:
	//					printf("Scrolled down");
	//					break;
	//				default:
	//					printf("cos");
	//				}
	//			}
	//	}

}
