package utils

import (
	"log"
	"net/http"
	"time"
)

// Response : a struct for json responses
type Response struct {
	Message string `json:"message"`
}

// NewResponse : returns a struct with a message you set as parameter
func NewResponse(m string) Response {
	var r Response
	r.Message = m

	return r
}

// CNInsert : Stands for "ConnectionStatementInsert"
var CNInsert = "toolset_insert:password@/toolset"

// CNSelect : Stands for "ConnectionStatementSelect"
var CNSelect = "toolset_select:password@/toolset"

// ConvertTime : convert the time from go to a string,
// so it complies with mysql standard for DATETIME.
// format used is "YYYY-MM-DD hh:mm:ss"
func ConvertTime(t *time.Time, s *string) {

	if t.IsZero() {
		var tm time.Time

		var st = *s

		tm, err := time.Parse("2006-01-02 15:04:05", st)
		if err != nil {
			log.Panicln(err.Error())
		}

		*t = tm
	} else if len(*s) <= 0 {
		var tm string

		var tt = *t

		tm = tt.Format("2006-01-02 15:04:05")

		*s = tm
	}
}

// TODO: find out how to get the response code from a http response

// LogRequest : this func is used to uniformely log the response of a http request
//		Note: this is not a http handler !
func LogRequest(req *http.Request) {
	log.Printf("Method:\"%s\" on Route:\"%s\"", req.Method, req.URL.Path)
}
