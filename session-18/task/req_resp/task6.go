package req_resp

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Number interface {
	int | float64
}

func stringToNumber(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to number: %w", a, err)
	}
	num2, err2 := strconv.ParseFloat(b, 64)
	if err2 != nil {
		return 0, fmt.Errorf("failed to convert %s to number: %w", b, err)
	}
	return num1 / num2, nil
}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	var b = r.URL.Query().Get("b")

	if b == "0" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("%d %s Division by zero is not allowed.", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))))
		if err != nil {
			log.Default().Println(err.Error())
		}
	} else {
		result, errDivide := stringToNumber(a, b)
		if errDivide != nil {
			fmt.Println(errDivide)
		}
		_, err := w.Write([]byte(fmt.Sprintf("Result: %.2f", result)))
		if err != nil {
			log.Default().Println(err.Error())
		}
	}

}
func Task6() {
	fmt.Println("Task6   *****************")
	http.HandleFunc("/divide", handleDivide)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
