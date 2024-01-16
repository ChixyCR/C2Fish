#### [English] [[中文]](./README.ZH.md) [[API Document]](./API.md) [[API文档]](./API.ZH.md)

---

<h1 align="center">Spark</h1>

**Spark** is a free, safe, open-source, web-based, cross-platform and full-featured RAT (Remote Administration Tool)
that allow you to control all your devices via browser anywhere.

We **won't** collect any data, thus the server will never self-upgrade. Your clients will only communicate with your
server forever.

---

<div align="center">

|![GitHub repo size](https://img.shields.io/github/repo-size/DGP-Studio/Snap.Genshin?style=flat-square)|![GitHub issues](https://img.shields.io/github/issues/XZB-1248/Spark?style=flat-square)|![GitHub closed issues](https://img.shields.io/github/issues-closed/XZB-1248/Spark?style=flat-square)|
|-|-|-|

|[![GitHub downloads](https://img.shields.io/github/downloads/XZB-1248/Spark/total?style=flat-square)](https://github.com/XZB-1248/Spark/releases)|[![GitHub release (latest by date)](https://img.shields.io/github/downloads/XZB-1248/Spark/latest/total?style=flat-square)](https://github.com/XZB-1248/Spark/releases/latest)|
|-|-|

</div>

---

## Notice

Due to my busy schedule with personal matters and the abuse of this project for cyberattacks, it's going to reach its end of life and will be archived very soon.

I will no longer provide any support for this project, as it is officially abandoned.

---

## Disclaimer

**THIS PROJECT, ITS SOURCE CODE, AND ITS RELEASES SHOULD ONLY BE USED FOR EDUCATIONAL PURPOSES.**
<br />
**ALL ILLEGAL USAGE IS PROHIBITED!**
<br />
**YOU SHALL USE THIS PROJECT AT YOUR OWN RISK.**
<br />
**THE AUTHORS AND DEVELOPERS ARE NOT RESPONSIBLE FOR ANY DAMAGE CAUSED BY YOUR MISUSE OF THIS PROJECT.**

**YOUR DATA IS PRICELESS. THINK TWICE BEFORE YOU CLICK ANY BUTTON OR ENTER ANY COMMAND.**

If you found any security vulnerability, please **DO NOT** open an issue and immediately contact me via [**email**](mailto:i@1248.ink).

---

## Quick start

### binary

* Download executable from [releases](https://github.com/XZB-1248/Spark/releases).
* Following [this](#Configuration) to complete configuration.
* Run executable and browse to `http://IP:Port` to access the web interface.
* Generate a client and run it on your target device.
* Enjoy!

---

## Configuration

Configuration file `config.json` should be placed in the same directory as the executable file.
<br />
Example:

  ```json
  {
      "listen": ":8000",
      "salt": "123456abcdef",
      "auth": {
          "username": "password"
      },
      "log": {
          "level": "info",
          "path": "./logs",
          "days": 7
      }
  }
  ```

* `listen` `required`, format: `IP:Port`
* `salt` `required`, length <= 24
  * after modification, you need to re-generate all clients
* `auth` `optional`, format: `username:password`
  * hashed-password is highly recommended
  * format: `$algorithm$hashed-password`, example: `$sha256$11223344556677AABBCCDDEEFF`
  * supported algorithms: `sha256`, `sha512`, `bcrypt`
  * if you don't follow the format, password will be treated as plain-text
* `log` `optional`
  * `level` `optional`, possible value: `disable`, `fatal`, `error`, `warn`, `info`, `debug`
  * `path` `optional`, default: `./logs`
  * `days` `optional`, default: `7`

---

## Features

| Feature/OS      | Windows | Linux | MacOS |
|-----------------|---------|-------|-------|
| Process manager | ✔       | ✔     | ✔     |
| Kill process    | ✔       | ✔     | ✔     |
| Network traffic | ✔       | ✔     | ✔     |
| File explorer   | ✔       | ✔     | ✔     |
| File transfer   | ✔       | ✔     | ✔     |
| File editor     | ✔       | ✔     | ✔     |
| Delete file     | ✔       | ✔     | ✔     |
| Code highlight  | ✔       | ✔     | ✔     |
| Desktop monitor | ✔       | ✔     | ✔     |
| Screenshot      | ✔       | ✔     | ✔     |
| OS info         | ✔       | ✔     | ✔     |
| Terminal        | ✔       | ✔     | ✔     |
| * Shutdown      | ✔       | ✔     | ✔     |
| * Reboot        | ✔       | ✔     | ✔     |
| * Log off       | ✔       | ❌     | ✔     |
| * Sleep         | ✔       | ❌     | ✔     |
| * Hibernate     | ✔       | ❌     | ❌     |
| * Lock screen   | ✔       | ❌     | ❌     |

* Blank cell means the situation is not tested yet.
* The Star symbol means the function may need administration or root privilege.

---

## Screenshots

![overview](./docs/overview.png)

![terminal](./docs/terminal.png)

![desktop](./docs/desktop.png)

![procmgr](./docs/procmgr.png)

![explorer](./docs/explorer.png)

![overview.cpu](./docs/overview.cpu.png)

![explorer.editor](./docs/explorer.editor.png)

---

## Development

### note

There are three components in this project, so you have to build them all.

Go to [Quick start](#quick-start) if you don't want to make yourself boring.

* Client
* Server
* Front-end

If you want to make client support OS except linux and windows, you should install some additional C compiler.

For example, to support android, you have to install [Android NDK](https://developer.android.com/ndk/downloads).

### tutorial

```bash
# Clone this repository.
$ git clone https://github.com/XZB-1248/Spark
$ cd ./Spark


# Here we're going to build front-end pages.
$ cd ./web
# Install all dependencies and build.
$ npm install
$ npm run build-prod


# Embed all static resources into one single file by using statik.
$ cd ..
$ go install github.com/rakyll/statik
$ statik -m -src="./web/dist" -f -dest="./server/embed" -p web -ns web


# Now we should build client.
# When you're using unix-like OS, you can use this.
$ mkdir ./built
$ go mod tidy
$ go mod download
$ ./scripts/build.client.sh


# Finally we're compiling the server side.
$ mkdir ./releases
$ ./scripts/build.server.sh
```

Then create a new directory with a name you like.
<br />
Copy executable file inside `releases` to that directory.
<br />
Copy the whole `built` directory to that new directory.
<br />
Copy configuration file mentioned above to that new directory.
<br />
Finally, run the executable file in that directory.

---

## Dependencies

Spark contains many third-party open-source projects.

Lists of dependencies can be found at `go.mod` and `package.json`.

Some major dependencies are listed below.

### Back-end

* [Go](https://github.com/golang/go) ([License](https://github.com/golang/go/blob/master/LICENSE))

* [gin-gonic/gin](https://github.com/gin-gonic/gin) (MIT License)

* [imroc/req](https://github.com/imroc/req) (MIT License)

* [kbinani/screenshot](https://github.com/kbinani/screenshot) (MIT License)

* [shirou/gopsutil](https://github.com/shirou/gopsutil) ([License](https://github.com/shirou/gopsutil/blob/master/LICENSE))

* [gorilla/websocket](https://github.com/gorilla/websocket) (BSD-2-Clause License)

* [orcaman/concurrent-map](https://github.com/orcaman/concurrent-map) (MIT License)

### Front-end

* [React](https://github.com/facebook/react) (MIT License)

* [Ant-Design](https://github.com/ant-design/ant-design) (MIT License)

* [axios](https://github.com/axios/axios) (MIT License)

* [xterm.js](https://github.com/xtermjs/xterm.js) (MIT License)

* [crypto-js](https://github.com/brix/crypto-js) (MIT License)

### Acknowledgements

* [natpass](https://github.com/lwch/natpass) (MIT License)
* Image difference algorithm inspired by natpass.

---

### Stargazers over time

[![Stargazers over time](https://starchart.cc/XZB-1248/Spark.svg)](https://starchart.cc/XZB-1248/Spark)

---

## License

[BSD-2 License](./LICENSE)