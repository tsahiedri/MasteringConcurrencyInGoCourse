package Exercise1_MergeSort

func MergeSortConcurrent(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = MergeSortConcurrent(data[:mid])
		done <- true //blocks the statement in line 17 so that it does not proceed until our goroutine has finished.
	}()
	right := MergeSortConcurrent(data[mid:])
	<-done
	return Merge(left, right)
}
