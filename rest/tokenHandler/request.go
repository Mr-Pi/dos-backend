package tokenHandler

import (
	"encoding/base64"
	"github.com/Mr-Pi/dos-backend/database/pgsql"
	"github.com/Mr-Pi/dos-backend/database/redis"
	"github.com/Mr-Pi/dos-backend/permissions"
	"github.com/emicklei/go-restful"
	"log"
	"strings"
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
	username := headerParts[0]
	password := headerParts[1]
	user := pgsql.GETUser(username)
	log.Println(user)
	if user.Password != permissions.HashPasswordOld(password, user.Salt) {
		resp.WriteErrorString(401, "Wrong Password")
		return
	}

	responseObj := tokenResponse{redis.NewAuthToken(headerParts[0])}
	resp.WriteEntity(responseObj)
}
