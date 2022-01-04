package api

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

func (s *ServerEnvironment) calculatorView(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	//var request CalculatorRequest
	//var response CalculatorResponse

	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("calculator.html")
	t.Execute(w, nil)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Msg("GET::CalculatorView Page Loading")
	//update response

}
func (s *ServerEnvironment) webEvaluate(w http.ResponseWriter, r *http.Request) {
	/*
		//declare response variable
		var request CalculatorRequest
		var response CalculatorResponse

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		rs = fmt.Sprintf("%v", response)
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Msg("POST::evaluate --> JSON:" + rs)
		//update response
	*/
}
