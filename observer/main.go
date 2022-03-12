package main

func main() {
	listener1 := DataListener{
		CustomerName: "Quang Le",
	}

	subj := &DataSubject{}
	subj.registerObserver(listener1)
	subj.ChangeStatusItem("Finish!")
}
