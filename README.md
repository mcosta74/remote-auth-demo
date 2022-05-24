# remote-auth-demo
simple project that shows how to use traefik's ForwardAuth middleware

## Build & Run

You need to install Docker and run

```sh
> docker compose up -d
```

At this point you have traefik listening at port 80.

Now you can send requests

### Missing Authorization Header
```sh
> curl -v -H Host:cheers.docker.localhost http://localhost/cheers                             *   Trying 127.0.0.1:80...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET /cheers HTTP/1.1
> Host:cheers.docker.localhost
> User-Agent: curl/7.79.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 401 Unauthorized
< Content-Length: 28
< Content-Type: text/plain; charset=utf-8
< Date: Tue, 24 May 2022 12:23:11 GMT
< 
* Connection #0 to host localhost left intact
missing authorization header
```

### Malformed Header
curl -v -H Host:cheers.docker.localhost -H Authorization:blahh http://localhost/cheers      *   Trying 127.0.0.1:80...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET /cheers HTTP/1.1
> Host:cheers.docker.localhost
> User-Agent: curl/7.79.1
> Accept: */*
> Authorization:blahh
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 401 Unauthorized
< Content-Length: 28
< Content-Type: text/plain; charset=utf-8
< Date: Tue, 24 May 2022 12:26:16 GMT
< 
* Connection #0 to host localhost left intact
invalid authorization header
### Invalid Token
```sh
> curl -v -H Host:cheers.docker.localhost -H Authorization:"Bearer foo" http://localhost/cheers
*   Trying 127.0.0.1:80...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET /cheers HTTP/1.1
> Host:cheers.docker.localhost
> User-Agent: curl/7.79.1
> Accept: */*
> Authorization:Bearer foo
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 401 Unauthorized
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
< Date: Tue, 24 May 2022 12:24:17 GMT
< 
* Connection #0 to host localhost left intact
invalid token
```

### Valid Token
```sh
curl -v -H Host:cheers.docker.localhost -H Authorization:"Bearer 1122334455" http://localhost/cheers
*   Trying 127.0.0.1:80...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET /cheers HTTP/1.1
> Host:cheers.docker.localhost
> User-Agent: curl/7.79.1
> Accept: */*
> Authorization:Bearer 1122334455
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Length: 16
< Content-Type: text/plain; charset=utf-8
< Date: Tue, 24 May 2022 12:24:55 GMT
< 
* Connection #0 to host localhost left intact
cheers mcosta74!
```