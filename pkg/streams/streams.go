package streams

import "sync"

type Streams struct{}

/// CreateIntStream converts a slice of integers to a channel of integers
/// Representative of a Stream in Java
func (s *Streams) CreateIntStream(input []int) <-chan int {
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

func (s *Streams) CreateStringStream(input []string) <-chan string {
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

func (s *Streams) CreateInterfaceStream(input []interface{}) <-chan interface{} {
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
