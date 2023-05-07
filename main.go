package main

import "github.com/starcorn2020/ds_nostr_go/repositories"

func main() {
	tool := repositories.NewTools()
	tool.GenKey()

}
