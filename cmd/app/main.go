package main

import "flag"

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "")
	flag.Parsed()

}
