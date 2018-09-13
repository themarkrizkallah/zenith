# zenith
**zenith** is a [Go](https://www.golang.org)-based, light-weight CLI [Stellar](https://www.stellar.org) wallet.

## Getting Started

These instructions will get you a copy of **zenith** up and running on your local machine for development and testing purposes.

### Prerequisites
You must have [Go](https://golang.org/doc/install) installed.


### Installing
All you need to do is run:
```
go get github.com/themarkrizkallah/zenith
```

**zenith** is now installed in your `$GOPATH/bin` directory !

**Test Run:**
```
$ zenith
A Go based Stellar wallet!
Here's information on how to use it:

Usage:
  zenith [command]

Available Commands:
  balance     Display the balance of an account
  generate    Generate a new stellar wallet
  help        Help about any command
  pay         Send XLM

Flags:
      --config string   config file (default is $HOME/.zenith.yaml)
  -h, --help            help for zenith
  -t, --tnet            use testnet instead of the real network

Use "zenith [command] --help" for more information about a command.

$ zenith generate
Your public key is: GD5VTJRNKJHDM36JZ4FI2YNZRC56JMSLQVCPMKQC2UD66J6HCNP546UK
Your secret key is: SBJ42JVJSAR7WPK7DR3NET3GRYDMQRQLSDXZY5XOPJ6LV65GOLJWHDH3
Save both of those keys offline somewhere safe.
Note: The account must have a minimum of 1 "XLM" in order to be active.
```

## Built With

* [Go](http://www.golang.org/) - Language used
* [dep](https://golang.github.io/dep/) - Dependency Management
* [stellargo](https://godoc.org/github.com/stellar/go) - The Stellar Go SDK
* [Cobra](https://github.com/spf13/cobra) - Powered the CLI

## License

This project is licensed under the APACHE License - see [LICENSE](LICENSE) for details.

## Acknowledgments

* Tip of the hat to [filidorwiese](https://github.com/filidorwiese/) 's [stellar-wallet](https://github.com/filidorwiese/stellar-wallet) for inspiration on how to structure this project.