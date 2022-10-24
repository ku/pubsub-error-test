package p

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// LaggyPushSubscriber returns ack but with 3s delay
func LaggyPushSubscriber(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message struct {
			Data string `json:"data"`
		} `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "empty body")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
	sDec, _ := b64.StdEncoding.DecodeString(d.Message.Data)
	s := string(sDec)

	a := strings.Split(s, ",")
	n := a[0]
	publishedAt, _ := strconv.Atoi(a[1])
	elapsed := time.Now().Unix() - int64(publishedAt)
	fmt.Printf("%s,%d\n", n, elapsed)

	seqnum, _ := strconv.Atoi(n)
	t := time.Unix(int64(publishedAt), 0)

	if elapsed < 1 {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("%d,%d,%d,%s\n", n, elapsed, seqnum, t.String())))
		return
	}

	time.Sleep(5 * time.Second)
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%d,%d,%d,%s\n", n, elapsed, seqnum, t.String())))
}
