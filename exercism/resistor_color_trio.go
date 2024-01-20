package resistorcolortrio

import (
    "strings"
    "strconv"
)

var colorMap = map[string]int{
        "black":  0,
        "brown":  1,
        "red":    2,
        "orange": 3,
        "yellow": 4,
        "green":  5,
        "blue":   6,
        "violet": 7,
        "grey":   8,
        "white":  9,
}


// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
    if len(colors) < 3 {
                return ""
        }
    
    zeroes := colorMap[colors[2]]
	if colors[1] == "black" {
		zeroes++
	}
	zeroesToPad := zeroes % 3
    buckets := []string{"", "kilo", "mega", "giga"}
    result := ""
    for i := 0; i < 2; i++ {
    	if colors[i] != "black" {
			result += strconv.Itoa(colorMap[colors[i]])
    	}       
    }
    return result + strings.Repeat("0", zeroesToPad) + " " + buckets[zeroes / 3] + "ohms"
}

