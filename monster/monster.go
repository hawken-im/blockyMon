package monster

import (
	"time"
)

type Growth int

const (
	Egg Growth = iota
	Baby
	Teen
	Adult
	Elder
	Dead
)

type Monster struct {
	age        int
	health     int
	mood       int
	sanitation int
	hunger     int
	live       bool
	sick       bool
	growth     Growth
}

func (m *Monster) Initialize() {
	m.age = 0
	m.health = 100
	m.mood = 100
	m.sanitation = 100
	m.hunger = 100
	m.live = true
	m.sick = false
	m.growth = Egg
}

func (m *Monster) Aging(t time.Time) {
	m.age += t.Second()
	m.mood -= 1

}

func (m *Monster) Moody(moodVal int) {
	m.mood += moodVal
	if m.mood>100 then m.mood = 100
}
