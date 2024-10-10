package main

import "flag"

var confPath = flag.String("confPath", "./config/local.yaml", "path to config file")

func parseFlags() *string {
	flag.Parse()
	return confPath
}
