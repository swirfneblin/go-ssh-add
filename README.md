# SSH Key Adder

This is a Go repository that executes a script to add SSH keys located in the `~/.ssh` folder of the device.
The `add_ssh_keys.sh` script reads all available keys in the folder and adds them using `ssh-add`.
The script is run by `main.go`, which will be compiled into a binary later.

## Prerequisites

Make sure you have the following software installed on your system:

- [Go](https://go.dev/doc/install)
- [OpenSSH](https://www.openssh.com/)

## Installation

1. Clone this repository to your local development environment and navigate to the project directory:

```
git clone https://github.com/swirfneblin/go-ssh-add.git
cd ssh-key-adder
```

## Usage

Run the add_ssh_keys.sh script to add the available SSH keys:

```
go build -o ssh-my
sudo mv ssh-my /usr/local/bin/
```

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues to report problems or suggestions.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit).
