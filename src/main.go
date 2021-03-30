package main

import (
	cli "./cli"
)

func main () {
	cli.Println(`{red}
	          _    _            _   
	    /\   | |  (_)          | |  
	   /  \  | | ___ _ __   ___| |_ 
	  / /\ \ | |/ / | '_ \ / _ \ __|
	 / ____ \|   <| | | | |  __/ |_ 
	/_/    \_\_|\_\_|_| |_|\___|\__|
`)

	cli.Execute()
}