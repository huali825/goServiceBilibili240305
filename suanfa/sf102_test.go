package suanfa

import (
	"fmt"
	"testing"
)

type FrequencyTracker struct {
	freq      map[int]int
	freqCount map[int]int
}

func Constructor2() FrequencyTracker {
	return FrequencyTracker{freq: make(map[int]int), freqCount: make(map[int]int)}
}

func (this *FrequencyTracker) Add(number int) {
	oldFreq := this.freq[number]
	this.freq[number]++
	newFreq := this.freq[number]

	this.freqCount[oldFreq]--
	if this.freqCount[oldFreq] == 0 {
		delete(this.freqCount, oldFreq)
	}
	this.freqCount[newFreq]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	oldFreq := this.freq[number]
	if oldFreq == 0 {
		return
	}

	this.freq[number]--
	newFreq := this.freq[number]

	this.freqCount[oldFreq]--
	if this.freqCount[oldFreq] == 0 {
		delete(this.freqCount, oldFreq)
	}
	if newFreq > 0 {
		this.freqCount[newFreq]++
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	_, exists := this.freqCount[frequency]
	return exists
}

func TestSuanfa102(t *testing.T) {
	fmt.Println("前缀树求数组任意区间和")
	s1 := []int{10, 20, 30, 40, 50}
	s2 := s1[2:4]
	fmt.Println(s2)
}
