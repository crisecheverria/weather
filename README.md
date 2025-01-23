# Weather App CLI

A weather CLI built with Go. Followed [tutorial](https://www.youtube.com/watch?v=zPYjfgxYO7k) on Youtube and changed a couple of things.

![Weather App CLI](https://github.com/crisecheverria/weather-app-golang/blob/main/terminal.png)

# Use

Install `go` and `git` if you don't have them.

```bash
git clone https://github.com/criech/weather-app-cli.git
cd weather-app-cli
go build
```

You can pass a location as argument (default is `Gothenburg`):

```bash
./weather London
```

You will need to create your own WEATHERAPI_KEY key for the weather API. You can do that [here](https://openweathermap.org/api).

```bash
export WEATHERAPI_KEY=YOUR_KEY
./weather London
```

# Run

```bash
./weather London
```

# Install

In Linux you can use `cp` to copy the binary to `/usr/local/bin` and run it from anywhere:

```bash
cp weather /usr/local/bin
weather London
```

In Mac you can use `mv` to move the binary to `/usr/local/bin` and run it from anywhere:

```bash
mv weather ~/usr/local/bin
weather London
```

# Run tests

```bash
go test
```
