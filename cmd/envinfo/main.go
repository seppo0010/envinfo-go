package main

import "github.com/seppo0010/envinfo-go"

func main() {
	goversion, _ := envinfo.GetGoVersion()
	println(goversion)
}
