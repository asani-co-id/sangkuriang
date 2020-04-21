package sangkuriang

import (
	"bytes"
	"database/sql"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/uniplaces/carbon"
)

func GetBody(r *http.Request) string {
	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	newStr := buf.String()
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(b)
	if !gjson.Valid(newStr) {
		return ""
	} else {
		return newStr
	}
}

func Suling(db *sql.DB, r *http.Request, id int) {
	times, _ := carbon.NowInLocation("Asia/Jakarta")
	logMainWhen := times.DateTimeString()
	request := GetBody(r)
	insForm, err := db.Prepare("INSERT INTO log_rental (log_rental_url, log_rental_request, log_rental_who, log_rental_when) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	go insForm.Exec(r.URL.Path, request, id, logMainWhen)
}
