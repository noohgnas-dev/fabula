# fabula

Yet another of instagram stories downloader. Downloading instagram stories without account. Not working if the instagram target are private.

Thanks to https://storiesig.info/en/ for providing the website and the API.

*Note: There is a rate limiter from Instagram website.*

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

*Note: There is a `Makefile` if you need to create the binary file by each operating system.*

*Note: Sometimes kinda buggy, just re-run until work.*

Build:
```bash
$ go build
```

Usage:

Show help

```bash
$ ./fabula --help
Usage of ./fabula:
  -username string
        the instagram username
```

Run instagram story downloader

```bash
$ ./fabula -username=john.doe
2023/02/28 01:29:29 main.go:36: =======================================
2023/02/28 01:29:29 main.go:37: = Name          : John Doe(@john.doe)
2023/02/28 01:29:29 main.go:38: = Followers     : 1000
2023/02/28 01:29:29 main.go:39: = Followings    : 500
2023/02/28 01:29:29 main.go:40: = Public Email  : john.doe@mail.com
2023/02/28 01:29:29 main.go:41: =======================================
2023/02/28 01:29:29 main.go:42: Found the user with 6 stories
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567834.mp4
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567434.mp4
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229566634.jpg
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567334.jpg
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567634.jpg
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567734.jpg
2023/02/28 01:29:31 main.go:63: Downloaded.. 1677518969229567634.jpg
2023/02/28 01:29:32 main.go:63: Downloaded.. 1677518969229567334.jpg
2023/02/28 01:29:36 main.go:63: Downloaded.. 1677518969229567734.jpg
2023/02/28 01:29:36 main.go:63: Downloaded.. 1677518969229566634.jpg
2023/02/28 01:29:39 main.go:63: Downloaded.. 1677518969229567834.mp4
2023/02/28 01:29:42 main.go:63: Downloaded.. 1677518969229567434.mp4
2023/02/28 01:29:42 main.go:68: All stories has been downloaded!
```

Check the `./stories` folder where the story contents has been downloaded.