package middlewares

import "net/http"

// this function will verify route for access token
func Auth(handlerFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the access token form the Headers
		// tokenString := r.Header.Get("Authentication")

	}
}
