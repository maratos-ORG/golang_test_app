package backend

import (
	"fmt"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

type pingRequest struct {
	ip       string
	response string
	wait     *sync.WaitGroup
}

var (
	requests = make(chan *pingRequest)
	ips      = make(map[string]*pingRequest)
	mutex    sync.Mutex
)


func pinger() {
	for req := range requests {
		go func(req *pingRequest) {
			output, err := exec.Command("ping", "-c", "10", req.ip).Output()
			var outputStr string
			if err != nil {
				outputStr = "Error: " + err.Error()
			} else {
				outputStr = string(output)
			}

			req.response = outputStr
			req.wait.Done()
		}(req)
	}
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Current time is:", time.Now())
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "missing ip", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	if req, ok := ips[ip]; ok {
		fmt.Println("IN LINE")
		mutex.Unlock()
		req.wait.Wait()
		fmt.Fprintln(w, string(req.response))
		return
	}

	req := &pingRequest{
		ip:       ip,
		response: "",
		wait:     &sync.WaitGroup{},
	}
	req.wait.Add(1)
	ips[ip] = req
	fmt.Println(ips)

	mutex.Unlock()

	go func() { requests <- req }()

	req.wait.Wait()
	fmt.Fprintln(w, string(req.response))
	mutex.Lock()
	delete(ips, ip)
	mutex.Unlock()
}
