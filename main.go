package main

import "flag"

func main()  {
	var client bool

	flag.BoolVar(&client, "client", false, "detail display")

	flag.Parse()

	if client{
		LunchSlave()
	}else{
		LunchMaster()
	}

	return

}
