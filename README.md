# Multi Port Server for Test

## Parameters

| Option | Description       | Default |
| ------ | ----------------- | ------- |
| -p     | Start Port Number | 58000   |
| -n     | Count             | 1       |

## Usage

```bash
$ portserver -p 58000 -n 100 &
Open Port From 58000 to 58099
....

$ curl http://localhost:58000
OK
```

## Build

```bash
$ go build
```
