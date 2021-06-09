# Go-Matomo

A simple Go SDK for sending data to a Matomo instance. This is useful for tracking user events on the server-side and publishing them to a Matomo site.

## Configuration

First, you must have a working Matomo installation. Not sure how to get started with Matomo? Refer to their [website](https://matomo.org/).

Sending analytics and events up is pretty straight-forward. Most data does not require any authentication, so the only configuration the SDK requires at this time is the domain of the installation. This is passed in through the environment to your running Go code:

`MATOMO_DOMAIN=https://matomo.mydomain.com`

Note the full protocol and domain. If Matomo is behind a different port, include that as well. Also note the lack of `/matomo.php` at the end. The SDK will take care of that for you.

When included, the SDK will run its `init` function to set up the configuration and read the domain. If one is not provided, you will see an error in the terminal. We will not panic or cause your application to stop due to this configuration error, so it is your responsibility to monitor for it.

If you are running the SDK on a system that is only responsible for reporting to a single site, then you can also specify the Site ID. This will allow the SDK to fill that in for you and also allows using the environment for different situations (such as testing the same Docker image from local to development to production):

`MATOMO_SITE_ID=1`

## Usage

Upon startup, the SDK will call it's `init` func, which calls its `Setup` func. This prepares the SDK for usage. When you have an event to send up, you will populate a `matomo.Parameters{}` struct. Most fields are optional. If they are `nil`, they will not be included. Since pointers are used to denote presence (as default values in Go are interpreted as present values by Matomo), you will want to use the `*Ptr` helper functions. For example:

```go
params := Parameters{
  RecommendedParameters: &RecommendedParameters{
    ActionName: matomo.StringPtr("visit"),
    URL: matomo.StringPtr("/users/me"),
    VisitorID: Smatomo.tringPtr("UNIQUE_HEX_VALUE"),
  },
  UserParameters: &UserParameters{
    UserID: matomo.StringPtr("mytestuser@mysite.com"),
  },
  EventTrackingParameters: &EventTrackingParameters{
    Category: matomo.StringPtr("site visit"),
    Action: matomo.StringPtr("loaded"),
    Name: matomo.StringPtr("user profile load"),
    Value: matomo.Float64Ptr(1.0),
  },
}
err := matomo.Send(&params)
```

The `RecommendedParamters.VisitorID` field may trip you up. If you have the user's DB identifier or some other identifying material, you can convert it to a 16 character hex. Or you can generate a pseudorandom one with the IP of the user, or whatever other identifying information you are comfortable collecting.

## Contributing

Contributors are welcome. You should raise an issue or communicate with us prior to committing any significant effort to ensure that your desired changes are compatible with where we want this library to go. Read more in the CONTRIBUTING.md document.
