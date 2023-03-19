# Random-Mouse

Move mouse randomly, Keep the computer active.

## How to install
```bash
go install github.com/dinghenc/random-mouse@v1.0.3
```

## How to use
```bash
Usage of random-mouse:
  -duration duration
        time duration of random-mouse exit (default 2h0m0s)
  -fresh duration
        mouse fresh time (default 1m0s)
  -time string
        absolute time of random-mouse exit, eg: (2006-01-02 15:04:05)
```

## Examples
```bash
# specify a refresh time of 1 second, and the specific time ends
random-mouse -fresh=1s -time="2006-01-02 15:04:05"

# specify a refresh time of 2 min, and end after a certain period of time
random-mouse -fresh=2m -duration=2h30m
```

## Todo
1. split methods to different package - done
2. support multiple display screen
