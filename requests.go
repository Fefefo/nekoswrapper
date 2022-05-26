package nekoswrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetEndpoints() (Endpoints, error) {
	resp, err := http.Get(Link + "endpoints")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var end Endpoints
	err = json.NewDecoder(resp.Body).Decode(&end)
	if err != nil {
		return nil, err
	}
	return end, nil
}

func GetResource(str string, num ...int) (Resources, error) {
	var n int
	if len(num) == 0 || num[0] < 1 {
		n = 1
	} else if num[0] > 20 {
		n = 20
	} else {
		n = num[0]
	}
	resp, err := http.Get(fmt.Sprintf("%s%s?amount=%d", Link, str, n))
	if err != nil {
		return nil, err
	}
	var res map[string]Resources
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.New("this endpoint doesn't exist")
	}
	defer resp.Body.Close()
	if val, ok := res["results"]; ok {
		return val, nil
	}
	return nil, errors.New("no results")
}
