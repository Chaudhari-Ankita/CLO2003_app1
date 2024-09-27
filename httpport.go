package main

func httpport(port int) {

	if port == 80 {
		print("you cant use this port its not safe\n")
	} else if port == 43 {
		println("accessible\n")
	} else {
		println("not accessible ports\n")
	}

}
