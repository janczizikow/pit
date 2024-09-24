package middleware

import (
	"net"
	"net/http"

	"github.com/jpillora/ipfilter"
)

// Get the IP address of the server's connected user.
func getUserIP(w http.ResponseWriter, r *http.Request) string {
	var userIP string
	if len(r.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = r.Header.Get("CF-Connecting-IP")
	} else if len(r.Header.Get("X-Forwarded-For")) > 1 {
		userIP = r.Header.Get("X-Forwarded-For")
	} else if len(r.Header.Get("X-Real-IP")) > 1 {
		userIP = r.Header.Get("X-Real-IP")
	} else {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		userIP = ip
	}
	return userIP
}

func IPFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := ipfilter.New(ipfilter.Options{
			BlockedIPs: []string{"173.77.194.16"},
		})
		ip := getUserIP(w, r)
		if net.ParseIP(ip) != nil && ip != "" && !f.Allowed(ip) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
