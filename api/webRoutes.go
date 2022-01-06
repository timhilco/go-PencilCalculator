package api

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/hlog"
	"github.com/timhilco/go-PencilCalculator/pencilCalculator"
)

func (s *ServerEnvironment) calculatorView(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	//var request CalculatorRequest

	expression := "1+2"
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("GET::CalculatorView Request: " + expression)
	//update response
	pr := evaluateAndRespond(expression, w, r)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("GET::CalculatorView Response: " + pr.String())

}
func evaluateAndRespond(expression string, w http.ResponseWriter, r *http.Request) pencilCalculator.PencilResult {
	var response CalculatorResponse
	ctx := context.Background()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("../api/calculator.html")
	if err != nil {
		msg := fmt.Sprintf("%s", err)
		hlog.FromRequest(r).Fatal().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("Error: HTML Calculator Page Loading: " + msg)
	}
	pencilResult := evaluate(ctx, true, expression)
	response = CalculatorResponse{
		Type:       fmt.Sprintf("%v", pencilResult.Type),
		Result:     fmt.Sprintf("%v", pencilResult.PrValue),
		Expression: expression,
	}
	t.Execute(w, response)

	return pencilResult
}
func (s *ServerEnvironment) webEvaluate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in webEvaluate")

	//declare response variable
	//var request CalculatorRequest
	expression := "2+3"

	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("POST::webEvaluate --> " + expression)
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s=%s\n", key, value)
	}
	expression = r.FormValue("expression")
	pr := evaluateAndRespond(expression, w, r)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("POST::webEvaluate --> " + pr.String())
}
func evaluate(ctx context.Context, isLocal bool, expression string) pencilCalculator.PencilResult {
	var pencilResult pencilCalculator.PencilResult
	if isLocal {
		pencilResult = evaluateLocal(ctx, expression)
	}
	return pencilResult
}
func evaluateLocal(ctx context.Context, expression string) pencilCalculator.PencilResult {

	empJson := `{
		"id": 11,
		"name": { "first": "John", "last" : "Sample"},
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000,
			"units" : 100
		}, {
			"effectiveDate": "01/01/2022",
			"amount": 50000,
			"units" : 110
		}]
	
	}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)
	pencilResult := pencilCalculator.Evaluate(ctx, expression)
	return pencilResult

}

/*
func evaluateRestApi(expression string) pencilCalculator.PencilResult {

}
*/
