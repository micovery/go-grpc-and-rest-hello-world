// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package greeter

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RestHelloWorld() func(c echo.Context) error {
	return func(c echo.Context) error {
		message := struct {
			Message string
		}{
			Message: "Hello World",
		}

		defer func() {
			c.Response().Header().Set("trailer1", "trailer-value1")
			c.Response().Header().Set("trailer2", "trailer-value2")
		}()

		c.Response().Header().Set("trailer", "trailer1, trailer2")
		c.Response().Header().Set("content-type", echo.MIMEApplicationJSONCharsetUTF8)

		enc := json.NewEncoder(c.Response())
		enc.Encode(message)
		c.Response().WriteHeader(http.StatusOK)
		c.Response().Flush()

		return nil

	}
}
