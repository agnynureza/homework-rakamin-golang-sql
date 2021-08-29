# homework-rakamin-golang-sql

## setup
1. clone this repository
2. type ```make run``` to run the apps
3. Hit the Server `localhost:8000/movies/titanic`
4. Unit Test ```make coverage```
5. see coverage all test in html ```make coverage-out```

### Tasks 
We define routes for handling operations:

| Method        | Route                  | Action                                              |
|---------------|------------------------|-----------------------------------------------------|
| GET           | /login                 | create token JWT                                    |
| POST          | /movie                 | create movie                                        |
| GET           | /movie/:slug           | get movie by slug                                   |
| PUT           | /movie/:slug           | update movue by slug                                |
| DELETE        | /movie/:slug           | delete movie by slug                                |

Access API via ```http://localhost:5000/{route}```


1. GET ```/login```

Response:
status code: 200
```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjkyMzU5MjF9.C0mI1Oc_U4E5A2l_qXhToFaw3epgHn1Jj2S1K2EnTdQ",
    "error": false,
    "msg": "success create token"
}
```

2. POST ```/movie ```

Authorization: Bearer {token} 

Request Body: 
```
{
   "title":"boy",
   "slug":"boy",
   "description":"lorem ipsum",
   "duration": 60,
   "image":"titanic poster url"
}
```

Response:
status code : 201
```
{
    "error": false,
    "msg": "success create data",
    "result": {
        "id": 13,
        "title": "boy",
        "Slug": "boy",
        "Description": "lorem ipsum",
        "Duration": 60,
        "Image": "titanic poster url"
    }
}
```

3. GET ```/movie/boy```

Authorization: Bearer {token} 

Response:
status code: 200
```
{
    "error": false,
    "msg": "success retrieve data",
    "result": {
        "id": 13,
        "title": "boy",
        "Slug": "boy",
        "Description": "lorem ipsum",
        "Duration": 60,
        "Image": "titanic poster url"
    }
}
```

4. PUT ```/movie/boy```

Authorization: Bearer {token} 

Response:
status code: 200
```
{
   "title":"boy",
   "slug":"titanic",
   "description":"lorem ipsum",
   "duration": 60,
   "image":"titanic poster url"
}
```

5. DELETE ```/movie/titanic```

Authorization: Bearer {token} 

Response:
status code: 200
```
{
    "error": false,
    "msg": "success"
}
```

### Tech Stack
* [Golang] - programming language
* [Fiber] - web framework with zero memory allocation and performance
* [Gomock] - mocking framework for the Go programming language.
* [JsonWebToken] - Authorization and Authentication 


[Golang]: <https://golang.org/>
[Fiber]: <https://github.com/gofiber/fiber/>
[Gomock]: <https://github.com/golang/mock/>
[JsonWebToken]: <https://jwt.io/>

