package permissions

import (
	"net/http"
	"github.com/Mr-Pi/dos-backend/database/redis"
	"strings"
	"encoding/base64"
)

func ReqToUser(r *http.Request) (username string, rc int) {
	authHeaders, ok := r.Header["Authorization"]
	if !ok {
		rc = 403
		return
	}
	if len(authHeaders) != 1 {
		rc = 400
		return
	}
	if !strings.HasPrefix(authHeaders[0], "Bearer ") {
		rc = 400
		return
	}
	b64 := strings.Split(authHeaders[0], " ")[1]
	rawHeader, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		rc = 400
		return
	}
	username = redis.AuthTokenToUser(string(rawHeader))
	if username == "" {
		rc = 401
		return
	}
	rc = 200
	return
}
