### Introduction
Used [RFC 7009](https://datatracker.ietf.org/doc/html/rfc7009).

#### Request
``` http
POST /v1/revoke HTTP/1.1
Authorization: Basic base64({CLIENT_ID}:{SECRET})
Host: {HOST}
 
token={token revoked}
```

#### Response success
```http
HTTP/1.1 200 OK
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