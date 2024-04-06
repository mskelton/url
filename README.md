# url

Parse URLs from a string.

## Installation

You can install url by running the install script which will download
the [latest release](https://github.com/mskelton/url/releases/latest).

```bash
curl -LSfs https://go.mskelton.dev/url/install | sh
```

Or you can build from source.

```bash
git clone git@github.com:mskelton/url.git
cd url
go install .
```

## Usage

```bash
echo 'This is https://foo.com url' | url
```
