package pub

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/binary"
	"time"
)


func Publish() {
	url := "http://127.0.0.1:1234/pub"
	fmt.Println("sending to: ", url)
	tic := time.NewTicker(1*time.Second)
	fails, count := 0, 0

	for {
		select {
		case <-tic.C:
			fmt.Printf("%d req/sec, fails: %d\n", count, fails)
			count = 0
			fails = 0
		default:
			msg := `{"message": "here's a log"}`
			buf := bytes.NewBuffer([]byte{})
			binary.Write(buf, binary.LittleEndian, int32(1))
			binary.Write(buf, binary.LittleEndian, int32(2))
			binary.Write(buf, binary.LittleEndian, int32(len(msg)))
			buf.WriteString(msg)

			req, _ := http.NewRequest("POST", url, buf)
			req.Header.Set("Content-Type", "application/octet-stream")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				if fails == 0 {
					fmt.Println(err.Error())
				}
				fails++
				continue
			}

			if resp.StatusCode == http.StatusOK {
				count++
			} else {
				fails++
			}

			resp.Body.Close()
		}
	}
}
