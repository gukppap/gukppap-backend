# API 명세서

# [GET] /api/v1/urls/shortcut
- shortcut 조회하기 

#### Request 
```json
Content-Type: application/json

{
    "url": "https://github.com/JJerBum"
}
```

#### Response 
```json
// Status OK
Content-Type: application/json

{
    "data": {
        "url": "https://github.com/JJerBum",
        "shortcut": "ai28h2z", 
        "exp": "1704359070"   
    }
}
```

```json
// Bad Reqeust
Content-Type: application/json

{
    "message": "등록되지 않은 url입니다."
}
```

# [POST] /api/v1/urls/shortcut
- shortcut 생성하기

#### Request 
```json
Content-Type: application/json

{
    "url": "https://github.com/JJerBum"
}
```

#### Response 
```json
// Status Created
Content-Type: application/json

{
    "data": {
        "url": "https://github.com/JJerBum",
        "shortcut": "ai28h2z",
        "exp": "1704359070"
    }
}
```

```json
// Bad Reqeust
Content-Type: application/json

{
    "message": "유효하지 않은 url입니다."
}
```

# [PATCH] /api/v1/urls/shortcut
- shortcut 수정하기

#### Request 
```json
Content-Type: application/json

{
    "shortcut": "ai28h2z",
    "want": "29skazl",
}
```

#### Response 

```json
// Status OK
```

```json
// Bad Reqeust
Content-Type: application/json

{
    "message": "없는 shortcut 입니다."
}
```


# [DELETE] /api/v1/urls/shortcut
- shortcut 삭제하기

#### Request 
```json
Content-Type: application/json

{
    "url": "https://github.com/JJerBum"
}
```

#### Response 
```json
// Status OK
```

```json
// Bad Reqeust
Content-Type: application/json

{
    "message": "없는 shortcut 입니다."
}
```