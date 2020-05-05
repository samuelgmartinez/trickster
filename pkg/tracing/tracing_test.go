/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tracing

import (
	"net/http"
	"strconv"
	"testing"

	"google.golang.org/grpc/codes"
)

func TestHTTPToCode(t *testing.T) {

	tests := []struct {
		code     int
		expected codes.Code
	}{
		{
			http.StatusMovedPermanently, codes.OK,
		},
		{
			http.StatusNotFound, codes.NotFound,
		},
		{
			http.StatusBadRequest, codes.InvalidArgument,
		},
		{
			http.StatusServiceUnavailable, codes.Unavailable,
		},
		{
			http.StatusInternalServerError, codes.Internal,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v := HTTPToCode(test.code)
			if v != test.expected {
				t.Errorf("expected %d got %d", test.expected, v)
			}
		})
	}

}