package routers

import(
	"net/http"
	"sync"
)

type Balancer struct{
	counter int
	ratio int
	mu *sync.Mutex
	serverlist []*Prox
}

func (lb *Balancer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	lb.mu.Lock()
	defer lb.mu.Unlock()

	ctr := 0
	if lb.counter != 0{
		ctr = 1
	}
	lb.serverlist[ctr].Handle(w, r)
	lb.counter = (lb.counter + 1)%(lb.ratio + 1)
}

func NewBalancer(proxys ...*Prox) *Balancer{
	return &Balancer{
		counter: 0,
		ratio: 4,
		mu: &sync.Mutex{},
		serverlist: proxys, 
	}
}

