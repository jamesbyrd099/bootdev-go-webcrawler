package main

type keyValuePair struct {
	key   string
	value int
}

func sortPairs(pairs []keyValuePair) {
	for index := 0; index < len(pairs); index++ {
		if index == len(pairs)-1 {
			return
		}
		if pairs[index].value > pairs[index+1].value {
			continue
		}
		if pairs[index].value < pairs[index+1].value {
			lowerPair := pairs[index+1]
			upperPair := pairs[index]
			pairs[index] = lowerPair
			pairs[index+1] = upperPair
			sortPairs(pairs)
			return
		}
		if pairs[index].key > pairs[index+1].key {
			lowerPair := pairs[index+1]
			upperPair := pairs[index]
			pairs[index] = lowerPair
			pairs[index+1] = upperPair
			sortPairs(pairs)
			return
		}
	}
}

func sortResults(results map[string]int) []keyValuePair {
	pairs := make([]keyValuePair, len(results))
	for key, value := range results {
		pairs = append(pairs, keyValuePair{
			key:   key,
			value: value,
		})
	}
	sortPairs(pairs)
	return pairs
}
