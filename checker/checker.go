package checker

import "net/http"

func IsAlive(url string) bool {
	resp, _ := http.Get(url)
	return resp != nil && resp.StatusCode == http.StatusOK
}
