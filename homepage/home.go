package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello World From homepage Package"

//Handlers has something to do with dependency injection
type Handlers struct {
	logger *log.Logger
}

//Home creates a function for when the homepage has been accessed
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utif-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

//Logger handles the log event for the home page
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		//Do this following next call which calls the Home Function
		defer h.logger.Printf("Request Processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
		//Looks like you have to use printf in golang -- couldnt use "" + time.Now()
	}
}

//SetupRoutes sets the default controller for the home page -- "/"
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	//pass the request of "/" to this.Logger function and pass this.Home function as the next call
	mux.HandleFunc("/", h.Logger(h.Home))
}

//New used for returning a pointer and Make used when you return the actual Type

//NewHandlers returns a pointer to a new Hander instance
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
