package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text and
// returns this data as a FreqMap.
func ConcurrentFrequency(list []string) FreqMap {
	ch := make(chan FreqMap, len(list))
	for _, str := range list {
		go func(s string) { ch <- Frequency(s) }(str)
	}

	freqMap := FreqMap{}
	for range list {
		for k, v := range <-ch {
			freqMap[k] += v
		}
	}
	return freqMap
}
