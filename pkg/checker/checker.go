package checker

import "net/http"

// IsAlive checks if site is alive.
func IsAlive(url string) bool {
	resp, _ := http.Get(url)
	return resp != nil && resp.StatusCode == http.StatusOK
}
