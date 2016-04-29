/*
Package scaleway provides a client for using the Scaleway API.

Construct a new Scaleway client,the use the various services on the client to
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
