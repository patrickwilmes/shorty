# Shorty - Simple url shortener service

Shorty is a dead simple url shortening service written in go.

```go get github.com/patrickwilmes/shorty```

## How it works

In order to create a short url one first needs to obtain a token:
```
POST http://localhost:8080/token
```

Afterward one can simply create a mapping:
```
POST http://localhost:8080/url
Content-Type: application/json

{
  "TargetUrl": "https://google.com",
  "Token": "0eb34fad-ea88-4840-a084-f2d210400527"
}
```

This will return the short url id generated.
Use:
```
GET http://localhost:8080/url/0eb34fad-ea88-4840-a084-f2d210400527
```
to get all your mappings. Use the hash in each mapping + the domain Shorty is
running on as a short url:
```
http://localhost:8080/<THE_SHORT_URL_HASH>
```

### More
To get to know the full api, have a look at the [requests](requests/).
