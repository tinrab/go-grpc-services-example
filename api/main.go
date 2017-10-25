package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tinrab/go-grpc-services-example/pb"
	"google.golang.org/grpc"
)

type server struct {
	addition       pb.AdditionServiceClient
	multiplication pb.MultiplicationServiceClient
}

func main() {
	s := server{}
	// Connect to addition service
	conn, err := grpc.Dial("add:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	s.addition = pb.NewAdditionServiceClient(conn)
	// Connect to multiplication service
	conn, err = grpc.Dial("multiply:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	s.multiplication = pb.NewMultiplicationServiceClient(conn)

	// Set up REST service
	r := mux.NewRouter()
	r.HandleFunc("/add/{a}/{b}", s.addHandler)
	r.HandleFunc("/multiply/{a}/{b}", s.multiplyHandler)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalln("Serving failed: %v", err)
	}
}

func (s server) addHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	// Parse arguments
	a, err := strconv.ParseFloat(vars["a"], 64)
	b, err := strconv.ParseFloat(vars["b"], 64)
	if err != nil {
		responseError(w, "Invalid arguments", http.StatusBadRequest)
		return
	}
	// Call service
	if res, err := s.addition.Add(ctx, &pb.AddRequest{A: a, B: b}); err == nil {
		responseJSON(w, res)
		return
	}
	responseError(w, "Could not add", http.StatusInternalServerError)
}

func (s server) multiplyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	// Parse arguments
	a, err := strconv.ParseFloat(vars["a"], 64)
	b, err := strconv.ParseFloat(vars["b"], 64)
	if err != nil {
		responseError(w, "Invalid arguments", http.StatusBadRequest)
		return
	}
	// Call service
	if res, err := s.multiplication.Multiply(ctx, &pb.MultiplyRequest{A: a, B: b}); err == nil {
		responseJSON(w, res)
		return
	}
	responseError(w, "Could not multiply", http.StatusInternalServerError)
}
