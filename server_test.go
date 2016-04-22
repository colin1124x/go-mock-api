package mock_api

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRun(t *testing.T) {

	r := map[string]map[string]string{
		"GET": map[string]string{
			"greet": "hello",
		},
	}

	q := make(chan bool)
	defer func() { q <- true }()

	s := New()
	s.Map(r)

	e := s.Run(":8000", r, q)
	if e != nil {
		t.Error(e)
		t.Fail()
	}

	c := http.Client{}
	req, e := http.NewRequest("GET", "http://127.0.0.1:8000/greet", nil)
	if e != nil {
		t.Error(e)
		t.Fail()
	}
	res, e := c.Do(req)
	if e != nil {
		t.Error(e)
		t.Fail()
	}
	defer res.Body.Close()
	b, e := ioutil.ReadAll(res.Body)
	if e != nil {
		t.Error(e)
		t.Fail()
	}
	if string(b) != "hello" {
		t.Errorf("response must be 'hello' but: %s", b)
	}

}
