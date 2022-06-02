### Introduction
Used [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662).

#### Request
``` http
POST /v1/introspect HTTP/1.1
Authorization: Basic base64({CLIENT_ID}:{SECRET})
Host: {HOST}
 
token={token introspection}
```

#### Response success
```http
HTTP/1.1 200 OK

{
    "active": true,
    "code": "68",
    "exp": 1633510245,
    "token_type": "Bearer"
}
```

#### **Response "Unauthorized"**
```http 
HTTP/1.1 401 Unauthorized
 
Unauthorized
```

#### **Response "Forbidden"**
```http 
HTTP/1.1 403 Forbidden
 
Forbidden
```