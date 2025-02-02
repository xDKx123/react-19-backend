package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"slices"
	"strings"
)

type StandardHttpResponseBody struct {
	Data   interface{} `json:"data,null"`
	Errors interface{} `json:"errors,null"`
}

//func StandardHttpResponse(c *gin.Context, data interface{}, errors interface{}, responseCode int) {
//	response := StandardHttpResponseBody{}
//
//	respCode := 200
//	if data != nil {
//		response.Data = data
//	} else {
//		response.Errors = errors
//		respCode = responseCode
//	}
//
//	c.JSON(respCode, response)
//}

func ignoreRouter(c *gin.Context) bool {
	pathsToIgnore := []string{"/", "/ping"}

	if slices.Contains(pathsToIgnore, c.Request.URL.Path) {
		return true
	}
	if strings.HasPrefix(c.Request.URL.Path, "/static") {
		return true
	}

	return false
}
func JSONResponseGinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ignoreRouter(c) {
			c.Next()
			return
		}

		bodyWrapper := &responseWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = bodyWrapper

		c.Next()

		statusCode := c.Writer.Status()

		if bodyWrapper.body.Len() == 0 {
			return
		}

		var originalData interface{}

		err := json.Unmarshal(bodyWrapper.body.Bytes(), &originalData)
		if err != nil {
			c.JSON(statusCode, StandardHttpResponseBody{Errors: "Invalid Response Format"})
			return
		}

		c.JSON(statusCode, StandardHttpResponseBody{Data: originalData})
	}
}

// responseWriter captures the response body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b) // Save response body
	return rw.ResponseWriter.Write(b)
}
