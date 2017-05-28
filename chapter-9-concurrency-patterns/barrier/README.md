# Barrier concurrency pattern

Barrier is a very common pattern, especially when we have to wait for more than one response from different Goroutines before letting the program continue.

## Situation

We have a microservices application where one service needs to compose its response by merging the responses of another three microservices.

Our Barrier pattern could be a service that will block its response until it has been composed with the results returned by one or more different Goroutines (or services). To block we could use an unbuffered channel in Go.

## Example

We are going to write a very typical situation in a microservices application - an app that performs two HTTP GET calls and joins them in a single response that will be printed on the console.

Our app must perform each request in a different Goroutine and print the result on the console if both responses are correct. If any of them returns an error, then we print just the error.

### Acceptance criteria

- Print on the console the merged results of the two calls to `http://httpbin.org/headers` and `http://httpbin.org/user-agent` URLs.
- If any of the calls fails, it must not print any result - just the error message (or error messages if both calls failed).
- The output must be printed as a composed result when both calls have finished. It means that we cannot print the result of one call and then the other.