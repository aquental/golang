package dndcharacter

import (
	"math"
	"math/rand"
	"time"
)
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}
// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2))
}
// Ability uses randomness to generate the score for an ability
func Ability() int {
	num := 0
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	min := math.MaxInt32
	for i := 0; i < 4; i++ {
		val := rnd.Intn(6) + 1
		num += val
		if val < min {
			min = val
		}
	}
	return num - min
}
// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    constitution := Ability()
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
    
    return c
}
