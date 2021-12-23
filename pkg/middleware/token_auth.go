package middleware

import (
	"encoding/json"
	"github.com/Dann-Go/book-store/internal/domain/responses"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	autHeader = "Authorization"
)

func Token_auth(ctx *gin.Context) {
	header := ctx.GetHeader(autHeader)
	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse("invalid auth header"))
		return
	}
	client := &http.Client{}
	var data = strings.NewReader("token=" + headerParts[1])
	req, err := http.NewRequest("POST", "https://login-demo.curity.io/oauth/v2/oauth-introspect", data)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	req.Header.Set("Authorization", "Basic ZGVtby1nYXRld2F5OmJGZlVVU1ZzV3c4QVlj")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(bodyText, &tmp); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if tmp["active"] == true {
		//ctx.JSON(http.StatusOK, responses.NewServerGoodResponse("valid token"))
		return
	} else {
		//		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewServerUnauthorizedResponse("token isn't active"))
		return
	}
}
