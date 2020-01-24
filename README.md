# skywire-peering-daemon

##### skywire-peering-daemon is a [Skywire](https://github.com/SkycoinProject/skywire-mainnet) daemon running functionality

## Build and run

### Requirements

`skywire-pering-daemon` requires a version of [golang](https://golang.org/) with [go modules](https://github.com/golang/go/wiki/Modules) support.

### Build

```
# Clone.
$ git clone git https://github.com/SkycoinProject/skywire-peering-daemon.git
$ cd skywire-peering-daemon

# Build
$ make build 

$ make install
```

### Run

```
# Run daemon
$ ./skywire-peering-daemon [publickey] [path/to/namedpipe]
```

The daemon uses the default port 3000 to send broadcasts and listen to incoming broadcasts. 
