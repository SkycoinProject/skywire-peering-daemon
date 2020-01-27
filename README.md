## skywire-peering-daemon

##### skywire-peering-daemon is a [Skywire](https://github.com/SkycoinProject/skywire-mainnet) daemon running functionality

#### Overview
A skywire-peering-daemon facilitates the setup of a local Skywire network via **stcp** transports by advertising a visor to other visors in the local network.  

### Build and run

#### Requirements

`skywire-pering-daemon` requires a version of [golang](https://golang.org/) with [go modules](https://github.com/golang/go/wiki/Modules) support.

#### Build

```
# Clone.
$ git clone git https://github.com/SkycoinProject/skywire-peering-daemon.git
$ cd skywire-peering-daemon

# Build
$ make build 

$ make install
```

#### Run

```
# Run daemon
$ ./skywire-peering-daemon [publickey] [path/to/namedpipe]
```

The daemon uses the default port 3000 to send broadcasts and listen to incoming broadcasts from other peering-daemons in a local network. It forwards packects received to [Skywire](https://github.com/SkycoinProject/skywire-mainnet) via **named pipes**.
