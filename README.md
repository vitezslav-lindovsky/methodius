# Methodius

### At the interview
They: *"Describe what different HTTP methods do."*  
Me: *"They will be doing what you make them do..."*  
They: *"...what?"*  

Although there is a standard RFC-7231 with recommendations, it depends on you how you want to use them.

### So what are HTTP methods for?

The typical HTTP 1.x request looks like this:  
The first line contains the method, path, and HTTP version.  
Then it's followed by headers, an empty line, and data.  

```
POST /status HTTP/1.1
Host: vitezslav-lindovsky.cz:8080
User-Agent: curl/8.5.0
Accept: */*
Content-Length: 5
Content-Type: application/x-www-form-urlencoded

happy
```

Server - Nginx, Apache, or our Golang app Methodius - will then parse this request,
and make some decisions based on the values (Method, path, Host header, ...), or forward it to another application (PHP app, Python app, ...).

After that, it will create a response. The response looks similar:  
HTTP version, status code, and status message on the first line.  
Then headers, an empty line, and body.  

```
HTTP/1.1 200 OK
Date: Sat, 26 Apr 2025 16:40:06 GMT
Content-Length: 19
Content-Type: text/plain; charset=utf-8

Updated key: status
```

Methodius creates random mapping of HTTP methods to actions. Plus there is one extra to quit.  
The main goal is to show that HTTP methods are not strictly defined, so it's kept as simple as possible.

```
methodius --rfc # will start in classic, expected mappping
methodius -p 8888 # will start on port 8888 instead of default 8080
methodius -v # will show what is it doing for each method
```

### Installation

```
go install github.com/vitezslav-lindovsky/methodius@latest
```
