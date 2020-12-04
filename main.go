package main

import (
	"encoding/json"
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"log"
	"net/http"
)

const PORT string = "4004"

type Resources struct {
	Memory *memory.Stats
	Cpu    *cpu.Stats
}

func StatusHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	var res = map[string]interface{}{
		"application-name": "resource monitoring",
		"status":           "OK",
	}

	bytes, _ := json.Marshal(res)
	_, _ = w.Write(bytes)
}

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	memStats, err := memory.Get()
	cpuStats, err := cpu.Get()
	if err != nil {
		log.Fatal(err)
	}

	res := Resources{
		Memory: memStats,
		Cpu:    cpuStats,
	}

	bytes, err := json.Marshal(res)
	_, _ = w.Write(bytes)
}

func main() {
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/resources", ResourcesHandler)

	fmt.Printf("Resouce Monitoring service listening on port %s\n", PORT)
	httpError := http.ListenAndServe(":"+PORT, nil)
	if httpError != nil {
		log.Fatal(httpError)
		return
	}
}
