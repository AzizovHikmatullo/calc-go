package application

import (
	"encoding/json"
	"github.com/AzizovHikmatullo/calc-go/pkg/calculation"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Config *Config
}

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

func New() *Application {
	return &Application{
		Config: ConfigFromEnv(),
	}
}

type calculateRequest struct {
	Expression string `json:"expression"`
}

type calculateResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(calculateResponse{Error: "Method not allowed"})
		return
	}

	var req calculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(calculateResponse{Error: "Expression is not valid"})
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(calculateResponse{Error: err.Error()})
		return
	}

	log.Println("New calculation. Expression:", req.Expression, " Result:", result)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(calculateResponse{Result: result})
}
