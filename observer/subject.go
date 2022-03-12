package main

type Observable interface {
	register(obs Observer)
	unregister(obs Observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

func (ds *DataSubject) ChangeStatusItem(data string) {
	ds.field = data

	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) unregisterObserver(o DataListener) {
	var newobs []DataListener
	for _, obs := range ds.observers {
		if o.CustomerName != obs.CustomerName {
			newobs = append(newobs, obs)
		}
	}
	ds.observers = newobs
}

func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onFinish(ds.field)
	}
}
