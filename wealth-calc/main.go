package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math"
	"net/http"
	"strconv"
)

//go:embed static/*
var staticFiles embed.FS

type YearDetail struct {
	Year         int     `json:"year"`
	StartBalance float64 `json:"startBalance"`
	Interest     float64 `json:"interest"`
	AnnualIncome float64 `json:"annualIncome"`
	EndBalance   float64 `json:"endBalance"`
	Formula      string  `json:"formula"`
}

type YearsToReachResp struct {
	Years   int          `json:"years"`
	Details []YearDetail `json:"details"`
}

type AfterYearsResp struct {
	FinalAmount float64      `json:"finalAmount"`
	Details     []YearDetail `json:"details"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

func formatMoney(v float64) string {
	s := strconv.FormatFloat(math.Round(v), 'f', 0, 64)
	n := len(s)
	if n <= 3 {
		return s
	}
	var result []byte
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, s[i])
	}
	return string(result)
}

func formatRate(r float64) string {
	return fmt.Sprintf("%.4f", r)
}

func buildFormula(startBalance, rate, annualIncome float64) string {
	interest := math.Round(startBalance * rate)
	return fmt.Sprintf("%s × %s + %s = %s + %s = %s",
		formatMoney(startBalance),
		formatRate(1+rate),
		formatMoney(annualIncome),
		formatMoney(startBalance+interest),
		formatMoney(annualIncome),
		formatMoney(startBalance+interest+annualIncome),
	)
}

func simulate(startBalance, annualIncome, rate float64, maxYears int) []YearDetail {
	var details []YearDetail
	balance := startBalance
	for y := 1; y <= maxYears; y++ {
		interest := balance * rate
		endBalance := balance + interest + annualIncome
		details = append(details, YearDetail{
			Year:         y,
			StartBalance: balance,
			Interest:     math.Round(interest),
			AnnualIncome: annualIncome,
			EndBalance:   math.Round(endBalance),
			Formula:      buildFormula(balance, rate, annualIncome),
		})
		balance = math.Round(endBalance)
	}
	return details
}

func parseFloatParam(r *http.Request, key string, defaultVal float64) (float64, error) {
	s := r.URL.Query().Get(key)
	if s == "" {
		return defaultVal, nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %s", key, s)
	}
	return v, nil
}

func parsePositiveFloatParam(r *http.Request, key string) (float64, error) {
	s := r.URL.Query().Get(key)
	if s == "" {
		return 0, fmt.Errorf("missing required param: %s", key)
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %s", key, s)
	}
	if v < 0 {
		return 0, fmt.Errorf("%s must be non-negative", key)
	}
	return v, nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func handleHowManyYears(w http.ResponseWriter, r *http.Request) {
	target, err := parsePositiveFloatParam(r, "target")
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	monthly, err := parsePositiveFloatParam(r, "monthly")
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	bonus, err := parsePositiveFloatParam(r, "bonus")
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	rate, err := parseFloatParam(r, "rate", 0)
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	initial, err := parseFloatParam(r, "initial", 0)
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}

	annualIncome := monthly*12 + bonus
	startBalance := initial

	if annualIncome <= 0 && startBalance <= 0 {
		writeJSON(w, 400, ErrorResp{Error: "收入或初始本金至少需要一个大于0"})
		return
	}

	if startBalance >= target {
		writeJSON(w, 200, YearsToReachResp{Years: 0, Details: nil})
		return
	}

	// 如果每年收入为0且有利率，也需要考虑
	if annualIncome == 0 && startBalance > 0 && rate > 0 {
		// 纯复利增长：startBalance * (1+rate)^n >= target
		years := math.Ceil(math.Log(target/startBalance) / math.Log(1+rate))
		details := simulate(startBalance, annualIncome, rate, int(years))
		writeJSON(w, 200, YearsToReachResp{Years: int(years), Details: details})
		return
	}

	if annualIncome == 0 && rate == 0 {
		writeJSON(w, 400, ErrorResp{Error: "当前参数下永远无法达到目标金额"})
		return
	}

	// 逐年模拟
	balance := startBalance
	for y := 1; y <= 500; y++ {
		interest := balance * rate
		balance = balance + interest + annualIncome
		if math.Round(balance) >= math.Round(target) {
			details := simulate(startBalance, annualIncome, rate, y)
			writeJSON(w, 200, YearsToReachResp{Years: y, Details: details})
			return
		}
	}

	writeJSON(w, 400, ErrorResp{Error: "500年内无法达到目标金额"})
}

func handleAfterYears(w http.ResponseWriter, r *http.Request) {
	years, err := parsePositiveFloatParam(r, "years")
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	yearsInt := int(years)
	if yearsInt <= 0 {
		writeJSON(w, 400, ErrorResp{Error: "years must be a positive integer"})
		return
	}
	if yearsInt > 500 {
		writeJSON(w, 400, ErrorResp{Error: "years too large, max 500"})
		return
	}

	monthly, err := parsePositiveFloatParam(r, "monthly")
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	bonus, err := parsePositiveFloatParam(r, "bonus")
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	rate, err := parseFloatParam(r, "rate", 0)
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}
	initial, err := parseFloatParam(r, "initial", 0)
	if err != nil {
		writeJSON(w, 400, ErrorResp{Error: err.Error()})
		return
	}

	annualIncome := monthly*12 + bonus
	details := simulate(initial, annualIncome, rate, yearsInt)
	finalAmount := details[len(details)-1].EndBalance

	writeJSON(w, 200, AfterYearsResp{FinalAmount: finalAmount, Details: details})
}

func main() {
	subFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(subFS)))
	mux.HandleFunc("/api/how-many-years", handleHowManyYears)
	mux.HandleFunc("/api/after-years", handleAfterYears)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
