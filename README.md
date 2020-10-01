# LCBO Car Fuel Index

Pull the api.lcbo.com REST API into a ~200MB sqlite database.

## Requirements

- Linux, MacOS, or WSL under Windows.
- [nix](https://nixos.org/download.html)

## Usage

Open a shell with the dependencies.

```sh
$ nix-shell
```

Scrape api.lcbo.com into a sqlite database, as defined in `.env`.

```sh
$ bin/fetch-json # Takes about 1hr
$ bin/make-db      # Takes about 5mins
```

Optionally, explore the data with [metabase](https://metabase.com)

```sh
$ metabase
```
