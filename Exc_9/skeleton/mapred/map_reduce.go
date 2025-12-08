package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce

func (*MapReduce) wordCountMapper(text string) []KeyValue {
	// Regex: keep only letters and spaces
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	cleaned := reg.ReplaceAllString(text, " ")
	cleaned = strings.ToLower(cleaned)

	words := strings.Fields(cleaned)
	kvs := make([]KeyValue, 0, len(words))

	for _, w := range words {
		kvs = append(kvs, KeyValue{Key: w, Value: 1})
	}
	return kvs
}

// Reducer receives a key and all its values and aggregates the sum
func (*MapReduce) wordCountReducer(key string, values []int) KeyValue {
	total := 0
	for _, v := range values {
		total += v
	}
	return KeyValue{Key: key, Value: total}
}

// Run executes the MapReduce algorithm concurrently
func (mr *MapReduce) Run(input []string) map[string]int {
	mapChan := make(chan []KeyValue)
	var wg sync.WaitGroup

	// --- MAP PHASE (concurrent) ---
	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			mapChan <- mr.wordCountMapper(text)
		}(line)
	}

	// Close channel when all mappers finish
	go func() {
		wg.Wait()
		close(mapChan)
	}()

	// --- SHUFFLE PHASE ---
	shuffled := make(map[string][]int)
	for kvlist := range mapChan {
		for _, kv := range kvlist {
			shuffled[kv.Key] = append(shuffled[kv.Key], kv.Value)
		}
	}

	// --- REDUCE PHASE (concurrent) ---
	result := make(map[string]int)
	var wg2 sync.WaitGroup
	var mu sync.Mutex

	for key, vals := range shuffled {
		wg2.Add(1)
		go func(k string, v []int) {
			defer wg2.Done()
			kv := mr.wordCountReducer(k, v)
			mu.Lock()
			result[kv.Key] = kv.Value
			mu.Unlock()
		}(key, vals)
	}

	wg2.Wait()
	return result
}
