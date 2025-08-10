package level2tasks

func Task4() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch)
	}()
	for n := range ch {
		println(n)
	}
}
