package async

type Async struct {
	Coroutine func(done chan bool)
	Started   bool
	Done      chan bool
}

type Coroutine string
type AsyncMap map[Coroutine]Async

var Routines AsyncMap = AsyncMap{}

func Create(id string, coroutine func(done chan bool)) Coroutine {
	_id := Coroutine(id)
	if _, ok := Routines[_id]; ok {
		return _id
	}

	a := Async{
		Coroutine: coroutine,
		Started:   false,
		Done:      make(chan bool),
	}
	Routines[_id] = a
	return _id
}

func (id Coroutine) Call() {
	if _, ok := Routines[id]; !ok {
		return
	}
	if Routines[id].Started {
		return
	}

	routine := Routines[id]
	routine.Started = true
	Routines[id] = routine
	go Routines[id].Coroutine(Routines[id].Done)
}

func (id Coroutine) Remove() {
	delete(Routines, id)
}

func (id Coroutine) CallOnce() {
	id.Call()
	go func() {
		for {
			select {
			case <-Routines[id].Done:
				id.Remove()
			}
		}
	}()
}
