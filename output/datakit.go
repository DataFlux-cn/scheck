package output

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
 send Msg to datakit
*/
type DatakitWriter struct {
	httpURL      string
	pending      []*sample
	maxPending   int
	lastSendTime int64
	samSig       chan *sample
}

func newDatakitWriter(filePath string, maxPending int) *DatakitWriter {
	if !strings.HasPrefix(filePath, "http://") && !strings.HasPrefix(filePath, "https://") {
		filePath = "http://" + filePath
	}

	dk := &DatakitWriter{
		httpURL:      filePath,
		maxPending:   maxPending,
		lastSendTime: time.Now().Unix(),
		samSig:       make(chan *sample, 1),
	}
	go dk.start()
	return dk
}

func (dk *DatakitWriter) Stop() {

}

func (dk *DatakitWriter) ReadMsg(measurement string, tags map[string]string, fields map[string]interface{}, t ...time.Time) {
	var data []byte
	var err error
	data, err = makeMetric(measurement, tags, fields, t...)
	if err != nil {
		return
	}
	dk.samSig <- &sample{data: data}

}

func (dk *DatakitWriter) start() {
	for {
		select {
		case sam := <-dk.samSig:
			dk.pending = append(dk.pending, sam)
			if len(dk.pending) > dk.maxPending {
				dk.ToUpstream(dk.pending...)
				dk.pending = make([]*sample, 0)
			}
		case <-time.After(time.Second * 10):
			if len(dk.pending) != 0 {
				dk.ToUpstream(dk.pending...)
			}
		}
	}
}

func (dk *DatakitWriter) ToUpstream(sams ...*sample) {
	var datas [][]byte
	for _, s := range sams {
		datas = append(datas, s.data)
	}
	data := bytes.Join(datas, []byte{'\n'})

	body, gz, err := buildBody(data)
	if err != nil {
		l.Errorf("build data err %v", err)
		return
	}

	l.Debugf("body:  %s", string(body))

	req, err := http.NewRequest("POST", dk.httpURL, bytes.NewBuffer(body))
	if err != nil {
		l.Errorf("%s", err)
		return
	}

	if gz {
		req.Header.Set("Content-Encoding", "gzip")
	}

	postbeg := time.Now()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.Errorf("%s", err)
		return
	}

	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Errorf("%s", err)
		return
	}

	switch resp.StatusCode / 100 {
	case 2:
		l.Debugf("post %d to %s ok(gz: %v), cost %v, response: %s",
			len(body), dk.httpURL, gz, time.Since(postbeg), string(respbody))
		return

	case 4:
		l.Errorf("post %d to %s failed(HTTP: %d): %s, cost %v, data dropped", len(body), dk.httpURL, resp.StatusCode, string(respbody), time.Since(postbeg))
		return

	case 5:
		l.Errorf("post %d to %s failed(HTTP: %s): %s, cost %v",
			len(body), dk.httpURL, resp.Status, string(respbody), time.Since(postbeg))
		return
	}

}
