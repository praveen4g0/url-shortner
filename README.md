# URL Shortner  Micro Service

Helps you to build simple shortend version of URL that will accept a URL as an argument over a REST API and return a shortened URL as a result.

## Design flow digaram

![Design-flow Diagram](images/url-shortener.png)

Design flow is simple we would focus on service(Bussiness logic) without having to worry about how it's used.

`eg:   (input) https://www.google.com -> 98sj1-293 (output)
 http://localhost:8000/98sj1-293 -> https://www.google.com`

## Model for Redirect service

```
type Redirect struct {
	Code      string  // to fetch url
	URL       string // URL it self
	CreatedAt int64  // created timestamp (metadata)
}
```

Here I am using 2 serailizers (json,msgPack) so user choice to choose his content-type for communicating b/w apis

And users choice to choose repo here it supports 2 (mongoDB/redis)

### Annotations to model
```
type Redirect struct {
	Code      string `json:"code" bson:"code" msgpack:"code"`
	URL       string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url`
	CreatedAt int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
}

```
We are using validator annotation in addition so to validate `URL` shouldn't be empty and follow particular format

## Shortner service
Here I am defining few interfaces which I will be implementing going forward
#### Service
```
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
```

1. `Find` method takes produced code and gives appropriate redirect url
2. `Store` method takes in riderct model and generates shortend code!


#### Serializer
```
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
```

1. `Decode` method takes in raw bytes and returns redirect struct/object
2. `Encode` mehtod takes in struct and retruns bytes

#### Repository

```
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
```

## Design Flow

it's pretty simple now service is connected to repository interface (implementation I would leave it to users choice)
Now we have build api business logic layer central part now we need to link our service to universal serilazer to could connect with transport http 

repo <--- service ---> serializer  ---> http
