// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package kakao provides constants for using OAuth2 to access Kakao.
package kakao // import "golang.org/x/oauth2/kakao"

import (
	"golang.org/x/oauth2/endpoints"
)

// Endpoint is Kakao's OAuth 2.0 endpoint.
var Endpoint = endpoints.KaKao
