### Introduction
Used [RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749).

#### Request
``` http
POST /v1/token HTTP/1.1
Authorization: Basic base64({CLIENT_ID}:{SECRET})
Host: {HOST}
```
Optional ```grant_type={grant}&scope={scope}```

#### Response success
```http
HTTP/1.1 200 OK
 
{
    "access_token": "e8877ceba95cf2fe418874db0dd3b97f23acb0cb9cfcd826d1586316a14f8a40",
    "token_type": "Bearer",
    "scope": "{some scope}",
    "expires_in": 1633508754
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