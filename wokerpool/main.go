package wokerpool

import (
	"os"
	"github.com/labstack/gommon/log"
	"net/http"
	"fmt"
	"time"
	"bytes"
	"encoding/json"
)

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue = os.Getenv("MAX_QUEUE")
)

type Job struct {
	PayLoad PayLoad
}

type PayLoad struct {
	UpdateToS3 func() error
	storageFolader string

}

type S3Bucket struct {
}

type PayloadCollection struct {
	WindowsVersion string `json:"version"`
	Token string 		  `json:"token"`
	Payloads []PayLoad    `json:"data"`
}

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool:pool, maxWorkers:maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i ++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch()  {

	for {
		select {
		case job := <- JobQueue:
			go func(job Job) {

				jobChannel := <- d.WorkerPool

				jobChannel <- job

			}(job)
		}
	}

}



var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit	chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool:workerPool,
		JobChannel:make(chan Job),
		quit: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {

		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <- w.JobChannel:
				if err := job.PayLoad.UpdateToS3; err != nil {
					log.Error("Error handling to S3:  %s ", err)
				}
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}

	}()
}


func (w Worker) Stop()  {
	go func() {
		w.quit <- true
	}()
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var content = &PayloadCollection{}

	// ,...


	for _, payload := range content.Payloads {
		work :=Job{PayLoad:payload}

		JobQueue <- work
	}

	w.WriteHeader(http.StatusOK)

}

func (p *PayLoad) UploadToS3() error {

	storage_path := fmt.Sprintf("%v/%v", p.storageFolader, time.Now().UnixNano())


	bucket := S3Bucket{}

	b := new(bytes.Buffer)

	encodeErr := json.NewEncoder(b).Encode(bucket)


	fmt.Println(storage_path, encodeErr)

	return nil
}













