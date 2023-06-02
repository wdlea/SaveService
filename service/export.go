package service

// Starts the exporting loop to the database, returns the *kill* channel
// which, when written to, will terminate all goroutines created. NOTE: that this
// channel needs to be written to only once.
func (S *SaveService[GameState_T]) StartExport(num_workers uint, buf_size uint) (kill chan int) {
	work := make(chan *ServiceEntry[GameState_T], buf_size)
	kill = make(chan int)

	for i := 0; i < int(num_workers); i++ {
		go exportWorker[GameState_T](work, kill)
	}

	go S.allocateWork(work, kill)

	return kill
}

func (S *SaveService[GameState_T]) allocateWork(work chan *ServiceEntry[GameState_T], kill chan int) {
	var currentEntry uint64 = 0
	for {
		entry := S.entries[uint64(currentEntry)]

		if entry.changed {
			work <- entry
		}
		currentEntry = (currentEntry + 1) % uint64(len(S.entries))

		select {
		case <-kill:
			return

		default:
			continue
		}
	}
}

func (E *ServiceEntry[GameState_T]) exportEntry() {
	E.mu.Lock()
	defer E.mu.Unlock()
	if E.changed {

	}
}

func exportWorker[GameState_T IGameState](work chan *ServiceEntry[GameState_T], kill chan int) {
	select {
	case current := <-work:
		{
			current.exportEntry()
		}
	case <-kill:
		{
			return
		}
	}
}
