/****************************************************************************
 * Copyright 2020, Optimizely, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/

// Package middleware
package middleware

import (
	"errors"
	"net/http"
)

func InvalidRequestHost(allowedHosts []string, allowedPort string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := GetLogger(r)
		logger.Debug().Strs("allowedHosts", allowedHosts).Str("allowedPort", allowedPort).Str("Host", r.Host).Str("X-Forwarded-Host", r.Header.Get("X-Forwarded-Host")).Str("Forwarded", r.Header.Get("Forwarded")).Msg("Request failed allowed hosts check")
		RenderError(errors.New("invalid request host"), http.StatusNotFound, w, r)
	})
}
