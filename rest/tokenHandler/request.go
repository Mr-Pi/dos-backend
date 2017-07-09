package tokenHandler

import (
	"github.com/emicklei/go-restful"
	"strings"
	"encoding/base64"
	"github.com/Mr-Pi/dos-backend/database/redis"
	"fmt"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func RequestToken(req *restful.Request, resp *restful.Response) {
	authHeaders, ok := req.Request.Header["Authorization"]
	if !ok {
		resp.WriteErrorString(400, "Authorization header missing")
		return
	}
	if len(authHeaders) != 1 {
		resp.WriteErrorString(400, "Wrong amount of Authorization headers")
		return
	}
	if !strings.HasPrefix(authHeaders[0], "Basic ") {
		resp.WriteErrorString(400, "Basic Authorization required")
		return
	}
	b64 := strings.Split(authHeaders[0], " ")[1]
	rawHeader, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	headerParts := strings.SplitN(string(rawHeader), ":", 2)
	if len(headerParts) != 2 {
		resp.WriteErrorString(400, "Password is missing")
		return
	}
	// TODO Verify with PostgreSQL
	fmt.Printf("User is %s\n", headerParts[0])
	fmt.Printf("Pass is %s\n", headerParts[1])

	responseObj := tokenResponse{redis.NewAuthToken(headerParts[0])}
	resp.WriteEntity(responseObj)
}
