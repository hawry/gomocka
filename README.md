# GOCKS

A very lightweight and simple mocking service. The idea is that a JSON file configures possible endpoints, and the service is either run locally or in a docker container.

<!-- TOC depthFrom:2 -->

- [Configuration](#configuration)
  - [Port](#port)
  - [Mocks](#mocks)
    - [Path](#path)
    - [Method](#method)
    - [Response Code](#response-code)
    - [Response Body](#response-body)
    - [Headers](#headers)
- [Build from source](#build-from-source)
- [Run in Docker](#run-in-docker)

<!-- /TOC -->

## Configuration

```json
{
  "port": 8080,
  "mocks": [
    {
      "path":"/",
      "method":"GET",
      "response_code": 200,
      "response_body": "{\"hello\":\"world\"}",
      "headers": {
        "Content-Type":"application/json",
        "x-trace-id": "thisisatrace"
      }
    },
    {
      "path":"/hello",
      "method":"POST",
      "response_code": 400,
      "response_body":"bad something"
    }
  ]
}
```

### Port
Port to listen to. The service automatically binds to `0.0.0.0:<port>`.

### Mocks
Array of what paths to respond to.

#### Path
Path to listen on. Relative to the root, to access `http://example.com/resource/action`, this attribute should contain `/resource/action`.

#### Method
One of the standard [HTTP Methods][1].

#### Response Code
Response code to return. The response code will be sent regardless of the payload body.

#### Response Body
String representation of the response body. The response body will be sent regardless of header value. To send an empty body, keep this attribute as a blank string: `"response_body":""`.

#### Headers
Headers to set in the response. You can use any string as a header key, and any string as a header value.

## Build from source

Clone this repository and run `make static` to create statically linked binary without CGO. The created binary can then be used as part of the docker build process mentioned above.

## Run in Docker
Before you try to build the docker image, make sure you have compiled a statically linked binary as explained above.

The project contains a lightweight Dockerfile, which only should be about 12MB in size. To create the docker image; just run `docker build -t gock .` in the project directory. Modify `settings.json` before building to make sure the latest configuration is part of the docker image.

To run the docker image: `docker run -d -p 8080:8080 gock`. Replace `8080` with the port number that you've specified in your `settings.json` file.


[1]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods