// Copyright 2016 The BreakBit Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package scaleway provides a client for using the Scaleway API.

Construct a new Scaleway client, then use the various services on the client to
access differents parts of the Scaleway API. For example:

    // Create a client
    client := scaleway.NewClient(nil)

    // Create a new token
    inBody := &TokenRequest{
        "foo@bar.com",
        "foobar",
        true,
    }

    // Do the action
    token, _, _ := client.Tokens.Create(inBody)

	client.AuthToken = token

*/
package scaleway
