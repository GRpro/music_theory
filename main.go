package main

import "fmt"

// C D E F G A B
var keys = [12]string{
	"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
}

const whole = 2
const half = 1

// Major
// F F H F F F H
var majorPattern = []int{whole, whole, half, whole, whole, whole, half}

// Minor
// F H F F H F F
var minorPattern = []int{whole, half, whole, whole, half, whole, whole}

func findIdx(note string) int {
	for i := 0; i < len(keys); i += 1 {
		if keys[i] == note {
			return i
		}
	}
	return -1
}

func majorProgression(note string) []string {
	seed := findIdx(note)

	var res []string

	i := seed

	for _, nextSeed := range majorPattern {
		res = append(res, keys[i])
		i = (i + nextSeed) % len(keys)
	}
	res = append(res, keys[i])

	return res
}

// C# (dies) = Db (bimol)
// E# = Fb

func minorProgression(note string) []string {
	seed := findIdx(note)

	var res []string

	i := seed

	for _, nextSeed := range minorPattern {
		res = append(res, keys[i])
		i = (i + nextSeed) % len(keys)
	}
	res = append(res, keys[i])

	return res
}

func majorChord3(note string) []string {
	progression := majorProgression(note)
	return []string{progression[0], progression[2], progression[4]}
}

func minorChord3(note string) []string {
	progression := minorProgression(note)
	return []string{progression[0], progression[2], progression[4]}
}

func main() {

	fmt.Println(majorProgression("C"))
	fmt.Println(majorProgression("G"))
	fmt.Println(majorProgression("D"))
	fmt.Println(majorProgression("A"))

	fmt.Println(minorProgression("C"))
	fmt.Println(majorProgression("A"))

	fmt.Println(majorChord3("A"))
	fmt.Println(minorChord3("A"))
}
