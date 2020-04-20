package sangkuriang

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/uniplaces/carbon"
)

var mutex sync.RWMutex

func Suling(db *sql.DB, r *http.Request, id int) {
	mutex.Lock()
	req, _ := r.GetBody()
	times, _ := carbon.NowInLocation("Asia/Jakarta")
	logMainWhen := times.DateTimeString()
	logMainRequest, _ := ioutil.ReadAll(req)
	insForm, err := db.Prepare("INSERT INTO log_rental (log_rental_url, log_rental_request, log_rental_who, log_rental_when) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(r.URL.Path, string(logMainRequest), id, logMainWhen)
	mutex.Unlock()
}
