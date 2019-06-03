package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//itemCount := 0
	for {
		result := <-out
		//if len(result.Items) == 0 {
		//	continue
		//}

		//for _, item := range result.Items {
		//
		//}
		go func() { e.ItemChan <- result.Items }()

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	//in := make(chan Request)
	go func() {
		for {
			//tell scheduler im ready
			ready.WorkerReady(in)
			request := <-in
			result, err := work(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
