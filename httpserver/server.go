package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Request struct {
	Weight string `json:"weight"`
	Steps  string `json:"steps"`
}

type Response struct {
	Calories float64 `json:"calories"`
}

func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", fs.ServeHTTP)

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		var req Request

		json.NewDecoder(r.Body).Decode(&req)

		weight, _ := strconv.ParseFloat(req.Weight, 64)
		steps, _ := strconv.ParseFloat(req.Steps, 64)

		calories := steps * weight * 0.0006

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(Response{
			Calories: calories,
		})
	})

	fmt.Println("server is listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server is not listening")
	}
}
