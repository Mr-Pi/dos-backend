package permissions

import (
	"net/http"
	"github.com/Mr-Pi/dos-backend/database/redis"
)

func ReqToUser(r *http.Request) (username string, rc int) {
	tokens, ok := r.Header["X-Tocken"]
	if !ok {
		rc = 403
		return
	}
	if len(tokens) != 1 {
		rc = 400
		return
	}
	username = redis.AuthTokenToUser(tokens[0])
	if username == "" {
		rc = 401
		return
	}
	rc = 200
	return
}
