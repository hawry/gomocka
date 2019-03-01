# GOCKS

A very lightweight and simple mocking service. The idea is that a JSON file configures possible endpoints, and the service is either run locally or in a docker container.

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

## Run in Docker
The project contains a lightweight Dockerfile, which only should be about 12MB in size. To create the docker image; just run `docker build -t gock .` in the project directory. Modify `settings.json` before building to make sure the latest configuration is part of the docker image.

To run the docker image: `docker run -d -p 8080:8080 gock`. Replace `8080` with the port number that you've specified in your `settings.json` file.


[1]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods