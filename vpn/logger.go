/*
 * App vpn
 *
 * API version: 1.0.0
 * Contact: info@menucha.de
 */

package vpn

import (
	"net/http"
	"time"

	"github.com/menucha-de/logging"
)

var log *logging.Logger = logging.GetLogger("vpn")

// Logger ...
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
