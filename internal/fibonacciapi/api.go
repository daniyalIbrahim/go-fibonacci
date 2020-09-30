package fibonacciapi

import (
	"encoding/json"
	"github.com/daniyalibrahim/go-fibonacci/internal/entity"
	"github.com/daniyalibrahim/go-fibonacci/pkg/log"
	"github.com/hashicorp/go-uuid"
	"net/http"
	"strconv"
	"sync"
)

/**
A type to store key value pairs for the object type Fibonacci.
*/
type FibonacciHandlers struct {
	sync.Mutex
	Logger log.Logger
	store  map[string]entity.FibonacciObject
}

func NewFibonacciHandlers() *FibonacciHandlers {
	return &FibonacciHandlers{
		store: map[string]entity.FibonacciObject{},
	}

}

/**
A handler function for getting fibonacci number for a given coresponding number
*/
func (h *FibonacciHandlers) FibonacciHandlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.Get(w, r)
		return
	case "POST":
		//h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func (h *FibonacciHandlers) Get(w http.ResponseWriter, r *http.Request) {
	logger := log.New().With(nil, "version", "1.0")

	//AllFiboRequests := make([]FibonacciObject, len(h.store))
	vars := r.URL.Query()
	// check for request parameters
	temp, ok := vars["num"]
	var nthNumber string
	if ok {
		if len(temp) >= 1 {
			nthNumber = temp[0] // The first `?type=model`
		}
	}
	logger.Infof("the input number %v is ", nthNumber)
	// convert the string to int
	newNthNumber, err := strconv.Atoi(nthNumber)
	// convert the int to uint64
	bigNthNumber := uint64(newNthNumber)

	id, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}
	var fibo = &entity.FibonacciObject{ID: id}

	fibo.ModifySecondVal(bigNthNumber)

	if bigNthNumber <= 11 {
		h.Lock()
		FiboValue := fibo.FindFibonacci(bigNthNumber)
		fibo.ModifyThirdVal(FiboValue)
		h.Unlock()
		logger.Infof("the fibonacci number %v for the number %v is ", nthNumber, FiboValue)
	} else if bigNthNumber >= 12 {
		FiboValue := uint64(100)
		fibo.ModifyThirdVal(FiboValue)
		logger.Infof("the fibonacci number %v for the number %v is ", nthNumber, FiboValue)
	}

	fibo.ModifyFourthVal()
	fibo.ModifyFifthVal()

	jsonBytes, err := json.Marshal(fibo)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	logger.Infof("the reply from api is ", string(jsonBytes))

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
