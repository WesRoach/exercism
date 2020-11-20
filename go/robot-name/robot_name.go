package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot holds a Robot's name
type Robot struct {
	name string
}

const maxNameCombinations = 26 * 26 * 10 * 10 * 10

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXZY")
var numbers = []rune("1234567890")
var usedNames map[string]bool = make(map[string]bool)
var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().Unix()))

// Name retrieves the name of a Robot
func (bot *Robot) Name() (string, error) {
	if len(bot.name) == 0 {
		if err := bot.setName(); err != nil {
			return bot.name, err
		}
	}
	return bot.name, nil
}

// Reset sets a Robot's name to ""
func (bot *Robot) Reset() {
	bot.name = ""
}

func (bot *Robot) setName() error {
	if len(usedNames) >= maxNameCombinations {
		return fmt.Errorf("namespace exhausted")
	}

	runeArray := make([]rune, 5)
	runeArray[0] = letters[rnd.Intn(len(letters))]
	runeArray[1] = letters[rnd.Intn(len(letters))]
	runeArray[2] = numbers[rnd.Intn(len(numbers))]
	runeArray[3] = numbers[rnd.Intn(len(numbers))]
	runeArray[4] = numbers[rnd.Intn(len(numbers))]
	name := string(runeArray)

	for usedNames[name] {
		// pick random runeArray[x] and re-random (or rotate) the value
		// re-random has a 1/26 & 1/10 chance of picking the same value
		// slight performance increase when rotating.
		// When approaching the maxNameCompinations, names will
		// increasingly consist of previous names with an off-by-one character
		randInt := rnd.Intn(len(runeArray))
		if randInt < 2 {
			// runeArray[randInt] = letters[rnd.Intn(len(letters))]
			runeArray[randInt] = rotateLetter(runeArray[randInt])
		} else {
			// runeArray[randInt] = numbers[rnd.Intn(len(numbers))]
			runeArray[randInt] = rotateNumber(runeArray[randInt])
		}
		name = string(runeArray)
	}

	bot.name = name
	usedNames[name] = true
	return nil
}

func rotateLetter(r rune) rune {
	if r == 'Z' {
		return 'A'
	}
	return r + 1

}

func rotateNumber(r rune) rune {
	if r == '9' {
		return '0'
	}
	return r + 1

}
