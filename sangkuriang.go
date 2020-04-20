package sangkuriang

import (
	"bytes"
	"database/sql"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/uniplaces/carbon"
)

func GetBodyJson(r *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	if !gjson.Valid(newStr) {
		return ""
	} else {
		return newStr
	}
}

func suling(db *sql.DB, r *http.Request, id int) {
	times, _ := carbon.NowInLocation("Asia/Jakarta")
	logMainWhen := times.DateTimeString()
	logMainRequest := GetBodyJson(r)
	insForm, err := db.Prepare("INSERT INTO log_rental (log_rental_url, log_rental_request, log_rental_who, log_rental_when) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	go insForm.Exec(r.URL.Path, logMainRequest, id, logMainWhen)
}