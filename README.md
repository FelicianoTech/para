# para - Packages and Releases Analytics

[![Build Status](https://circleci.com/gh/halseylabs/para.svg?style=shield)](https://circleci.com/gh/halseylabs/para) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/halseylabs/para/master/LICENSE)

`para` is a useful tool for software package manager analytics.
It gives you metrics on how your packages are doing on Brew and GitHub as well as helps you find available names.


## Installing

### Debian Package (.deb) Instructions

Download the `.deb` file to the desired system.

For graphical systems, you can download it from the [GitHub Releases page][gh-releases].
Many distros allow you to double-click the file to install.
Via terminal, you can do the following:

```bash
wget "https://github.com/halseylabs/para/releases/download/v0.1.0/para-v0.1.0_amd64.deb"
sudo dpkg -i para-v0.1.0_amd64.deb
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### Linux Snap

```bash
sudo snap install para
```

### macOS / Brew

```bash
sudo brew install halseylabs/tap/para
```


## Usage

Run `para help` to see all commands available.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/halseylabs/para/releases
