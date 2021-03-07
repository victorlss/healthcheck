package checker

import "net/http"

func IsAlive(url string) bool {
	resp, _ := http.Get(url)
	return resp.StatusCode == http.StatusOK
}
