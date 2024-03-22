package main


type FrequencyTracker struct {
	slice []int
	freq  map[int]int
}

func Constructor() *FrequencyTracker {
	return &FrequencyTracker{slice: make([]int, 0), freq: make(map[int]int)}
}

func (this *FrequencyTracker) Add(number int) {
	i, ok := this.freq[number]
	if ok != true {
		this.freq[number] = 1
	} else {
		this.freq[number] = i + 1
	}
	//this.slice = append(this.slice, number)

}

func (this *FrequencyTracker) DeleteOne(number int) {
	i, ok := this.freq[number]
	if ok != true {
		return
	} else if i == 1 {
		delete(this.freq, number)
	} else {
		this.freq[number] = i - 1
	}

}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	_, exists := this.freq[frequency]
	return exists
}


