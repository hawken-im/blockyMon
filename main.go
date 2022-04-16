package main

import (
	"blockymon/actions"
	"blockymon/monster"
)

func main() {
	var monster monster.Monster
	monster.Initialize()
	monster.mood = actions.Pet(monster.mood)

}
