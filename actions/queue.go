package actions

type ActionQueue struct {
	queue chan Action
}

func NewActionQueue(bufferSize int) *ActionQueue {
	return &ActionQueue{
		queue: make(chan Action, bufferSize),
	}
}

func (aq *ActionQueue) Add(action Action) {
	aq.queue <- action
}

func (aq *ActionQueue) GetAll() []Action {
	var actions []Action
	for {
		select {
		case action := <-aq.queue:
			actions = append(actions, action)
		default:
			return actions
		}
	}
}
