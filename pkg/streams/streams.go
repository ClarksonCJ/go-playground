package streams

import "sync"

/// CreateIntStream converts a slice of integers to a channel of integers
/// Representative of a Stream in Java
func CreateIntStream(input []int) <-chan int {
	out := make(chan int, len(input))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		for _, i := range input {
			out <- i
		}
	}()
	wg.Wait()
	return out
}

func CreateStringStream(input []string) <-chan string {
	out := make(chan string, len(input))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		for _, i := range input {
			out <- i
		}
	}()
	wg.Wait()
	return out
}

func CreateInterfaceStream(input []interface{}) <-chan interface{} {
	out := make(chan interface{}, len(input))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		for _, i := range input {
			out <- i
		}
	}()
	wg.Wait()
	return out
}
