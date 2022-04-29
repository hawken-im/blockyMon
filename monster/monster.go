package monster

import (
	"container/list"
	"fmt"
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

type Album struct {
	MonsterData Monster
	TimeStamp   time.Time
}

type Monster struct {
	ID         string
	Name       string
	Age        time.Duration
	Health     int
	Mood       int
	Sanitation int
	Hunger     int
	Live       bool
	Sick       bool
	Growth     Growth
	BirthDay   time.Time
}

func (m *Monster) Initialize(newID string, newName string, newBirthDay time.Time) {
	m.ID = newID
	m.Name = newName
	m.Age = 0
	m.Health = 100
	m.Mood = 40
	m.Sanitation = 100
	m.Hunger = 100
	m.Live = true
	m.Sick = false
	m.Growth = Egg
	m.BirthDay = newBirthDay

}

func (m *Monster) AlbumShot(currentTime time.Time, albumList *list.List) (photo Album) {
	var newPhoto Album
	newPhoto.MonsterData = *m
	newPhoto.TimeStamp = currentTime
	albumList.PushBack(newPhoto)
	return newPhoto
}

func (m *Monster) Grow(currentTime time.Time, albumList *list.List) (growRecord list.Element) {
	duration := currentTime.Sub(m.BirthDay)
	if !m.Live {

		growRecord = list.Element{Value: "dead"}
		return growRecord
	}
	m.Age += duration

	m.Mood -= int(duration.Seconds())
	m.Hunger -= int(duration.Seconds())
	m.Sanitation -= int(duration.Seconds())

	if m.Mood <= 50 || m.Sanitation <= 50 || m.Hunger <= 50 {
		m.Health -= int(duration.Seconds())
	}
	if m.Health <= 50 {
		m.Sick = true
	}
	if m.Sick {
		m.Health -= int(duration.Seconds())
	}
	if m.Health <= 0 {
		m.Live = false
		growRecord = list.Element{Value: "dead"}
		return growRecord
	}

	return
}

func (m *Monster) TakeAction(action string) {
	switch action {
	case "pet":
		m.Moody(10)

	case "clean":
		m.Sanny(10)

	case "feed":
		m.Munch(10)

	case "heal":
		m.Healthy(10)

	default:
		fmt.Println("unknown action")
	}
}

func (m *Monster) Healthy(healVal int) {
	fmt.Println("Heal your monster. Health up: ", healVal)
	m.Health += healVal
	if m.Health > 100 {
		m.Health = 100
	}
	fmt.Println("Health is: ", m.Health)
}

func (m *Monster) Moody(moodVal int) {
	fmt.Println("Pet your monster. Mood up: ", moodVal)
	m.Mood += moodVal
	if m.Mood > 100 {
		m.Mood = 100
	}
	fmt.Println("Mood is: ", m.Mood)
}

func (m *Monster) Sanny(cleanVal int) {
	fmt.Println("Clean your monster. Sanitation up: ", cleanVal)
	m.Sanitation += cleanVal
	if m.Sanitation > 100 {
		m.Sanitation = 100
	}
	fmt.Println("Sanitation is: ", m.Sanitation)
}

func (m *Monster) Munch(foodVal int) {
	fmt.Println("Feed your monster. Hunger up: ", foodVal)
	m.Hunger += foodVal
	if m.Hunger > 100 {
		m.Hunger = 100
	}
	fmt.Println("Hunger is: ", m.Hunger)
}
