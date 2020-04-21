package sangkuriang

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
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

func Suling(db *sql.DB, r *http.Request, id int) {
	b := bytes.NewBuffer(make([]byte, 0))
	io.TeeReader(r.Body, b)
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(b)
	fmt.Println(b.String)
	// times, _ := carbon.NowInLocation("Asia/Jakarta")
	// logMainWhen := times.DateTimeString()
	// insForm, err := db.Prepare("INSERT INTO log_rental (log_rental_url, log_rental_who, log_rental_when) VALUES (?, ?, ?)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// go insForm.Exec(r.URL.Path, id, logMainWhen)
}
