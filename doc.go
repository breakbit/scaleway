// Copyright 2016 The BreakBit Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package scaleway provides a client for using the Scaleway API.

Construct a new Scaleway client, then use the various services on the client to
access differents parts of the Scaleway API. For example:

    // Create a client
    client := scaleway.NewClient(nil)

    // Create credentials structure
    credentials := scaleway.NewCredentials("foo@bar.com", "foobar")

    // Create new token
    token, _, _ := client.Tokens.Create(credentials, true)

    // Use this token
	client.AuthToken = token.ID

*/
package scaleway
