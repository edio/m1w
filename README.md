# m1w

`m1w` stands for monolith 1-way

> 1-way monolith teleports hero one way to a specific map location.
>
> -- Heroes of Might and Magic documentation

It is a redirection/url-shortener service targeted at companies intranets. Full power and cool benefits of `m1w` are
revealed if `m1w` is deployed to be accessible under a short and memorable DNS name, for example `go`.

Then it may be used to save few key strokes for users, redirecting them to a popular resource:
```
go/w → wiki.mysupercoolcompany.com
```

Think of it as a browser-independent company managed shared bookmarks store:
```
go/codestyle → wiki.mysupercoolcompany.com/articles/engineering/java/codestyle
```

Or as a stable alias to a volatile url, for example, a link to a Google Doc that is changing regularly:
```
go/sprint-goals → docs.mysupercoolcompany.com/document/d/1234567890aBcDeFgHijklmnoPqrstUvwxYz
```

## Build and run as standalone binary

Compiler and tools for Go language are required.

```
$ make
# bin/m1w --db /path/to/database
```

You need root privileges to bind to port 80.

## Build and run in a docker container

For convenience of running `m1w` locally (no hassle with binding to port 80) a simple Dockerfile is provided. Docker is
obviously required.

```
$ make docker
$ docker run --name m1w -d -v /tmp/m1w.db:/m1w.db m1w
```
where `/tmp/m1w.db` is a directory on a host system, where `m1w` state will be stored.

## Usage

1. Bind hostname to `127.0.0.1` or to a docker container IP address
    ```
    # DOMAIN=go
    # echo "`docker inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}" m1w` $DOMAIN" >> /etc/hosts
    ```

2. Create an "entrance" in `m1w`
    ```
    curl -X POST go/g -d 'http://google.com'
    ```

3. "Teleport" by using the entrance
    ```
    curl -i -X GET go/g
    ```

## Features and development

`m1w` is developed for fun in my free time. Some essential features (ui, authentication) are still missing.
