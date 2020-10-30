package concurrency

func Or(channels ...<-chan interface{}) <-chan interface{} { //1

	switch len(channels) {
	case 0: //2
		return nil
	case 1: //3
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() { //4
		defer close(orDone)

		switch len(channels) {
		case 2: //5
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: //6
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(append(channels[3:], orDone)...): //6
			}
		}
	}()
	return orDone
}

func OrDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func Tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range OrDone(done, in) {
			var out1, out2 = out1, out2 //1
			for i := 0; i < 2; i++ {    //2
				select {
				case <-done:
				case out1 <- val:
					out1 = nil //3
				case out2 <- val:
					out2 = nil //3
				}
			}
		}
	}()

	return out1, out2
}

func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{}) // 1
	go func() {
		defer close(valStream)
		for { // 2
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range OrDone(done, stream) { // 3
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
