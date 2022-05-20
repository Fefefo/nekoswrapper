package nekoswrapper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var client *http.Client

func GetEndpoints() (Endpoints, error) {
	resp, err := client.Get(Link + "endpoints")
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
	if len(num) == 0 || num[0] < 1 {
		num[0] = 1
	} else if num[0] > 20 {
		num[0] = 20
	}
	resp, err := client.Get(fmt.Sprintf("%s%s?amount=%d", Link, str, num[0]))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var res Resources
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
