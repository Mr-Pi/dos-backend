package tokenHandler

import (
	"github.com/emicklei/go-restful"
	"strings"
	"encoding/base64"
	"github.com/Mr-Pi/dos-backend/database/redis"
)

type deleteResponse struct {
	Success bool `json:"success"`
}

func DeleteToken(req *restful.Request, resp *restful.Response) {
	authHeaders, ok := req.Request.Header["Authorization"]
	if !ok {
		resp.WriteErrorString(403, "Authorization header missing")
		return
	}
	if len(authHeaders) != 1 {
		resp.WriteErrorString(400, "Not the right amount of Authorization headers")
		return
	}
	if !strings.HasPrefix(authHeaders[0], "Bearer ") {
		resp.WriteErrorString(400, "Please give us a Bearer token. Not something else")
		return
	}
	b64 := strings.Split(authHeaders[0], " ")[1]
	rawHeader, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		resp.WriteErrorString(400, "Please give us some valid Base64. It's not that hard")
		return
	}
	result := redis.RemoveAuthToken(string(rawHeader))
	responseObj := deleteResponse{ result }
	resp.WriteEntity(responseObj)
}
