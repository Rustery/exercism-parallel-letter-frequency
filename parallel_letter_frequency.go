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

// ConcurrentFrequency concurrently counts the frequency of each rune in a given texts and returns this
// data as a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	m := FreqMap{}
	results := make(chan FreqMap, 10)

	for _, s := range texts {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	for range texts {
		for k, v := range <-results {
			m[k] += v
		}
	}

	return m
}
