/*
 Copyright 2018 Padduck, LLC
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 	http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package shared

import (
	"github.com/gin-gonic/gin"
	"github.com/pufferpanel/apufferi/http"
	netHttp "net/http"
	"strings"
)

func NotImplemented (c *gin.Context) {
	http.Respond(c).Fail().Status(netHttp.StatusNotImplemented).Message("not implemented").Send()
}

func CreateOptions(options ...string) gin.HandlerFunc {

	replacement := make([]string, len(options) + 1)

	copy(replacement, options)

	replacement[len(options)] = "OPTIONS"
	response := strings.Join(replacement, ",")

	return func (c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", response)
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", response)
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(netHttp.StatusOK)
	}
}