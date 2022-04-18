package monster

import "time"

type Growth int

const (
	Egg Growth = iota
	Baby
	Teen
	Adult
	Elder
	Dead
)

type SnapShot struct {
	MonsterData Monster
	TimeStamp   time.Time
}

type Monster struct {
	ID         string
	Name       string
	Age        int
	Health     int
	Mood       int
	Sanitation int
	Hunger     int
	Live       bool
	Sick       bool
	Growth     Growth
	BirthDay   time.Time
}

func (m *Monster) Initialize() {
	m.Age = 0
	m.Health = 100
	m.Mood = 100
	m.Sanitation = 100
	m.Hunger = 100
	m.Live = true
	m.Sick = false
	m.Growth = Egg
}

func (m *Monster) Aging(duration int) (f func()) {
	var count int = 0
	f = func() {
		if !m.Live {
			return
		}
		m.Age += duration
		count++
		if count >= 10 {
			m.Mood -= 1
			m.Hunger -= 1
			m.Sanitation -= 1
			count = 0
		}
		if m.Mood <= 50 || m.Sanitation <= 50 || m.Hunger <= 50 {
			m.Health -= 1
		}
		if m.Health <= 50 {
			m.Sick = true
		}
		if m.Sick {
			m.Health -= 1
		}
		if m.Health <= 0 {
			m.Live = false
		}
	}

	return
}

func (m *Monster) Healthy(healVal int) {
	m.Health += healVal
	if m.Health > 100 {
		m.Health = 100
	}
}

func (m *Monster) Moody(moodVal int) {
	m.Mood += moodVal
	if m.Mood > 100 {
		m.Mood = 100
	}
}

func (m *Monster) Sanny(cleanVal int) {
	m.Sanitation += cleanVal
	if m.Sanitation > 100 {
		m.Sanitation = 100
	}
}

func (m *Monster) Munch(foodVal int) {
	m.Hunger += foodVal
	if m.Hunger > 100 {
		m.Hunger = 100
	}
}
