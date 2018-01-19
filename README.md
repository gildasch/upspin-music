A music player for Upspin

Run it with:

```
$ go run main.go -config /path/to/your/upspin/config
```

And open your browser at http://localhost:8080/listen/a-user@mail.com/path/to/a/folder

![Screenshot](screenshot.png)

Additional configuration:

```
Listen to another port:
$ PORT=1234 go run main.go -config /path/to/your/upspin/config

Set baseURL:
$ go run main.go -config /path/to/your/upspin/config -baseURL https://myserver.com/music
```

It uses [Aplayer music player](https://github.com/MoePlayer/APlayer)
