# LCBO Car Fuel Index

Pull the [LCBO REST API](https://api.lcbo.com) into a ~200MB sqlite database. Originally intended to be used to figure out the least expensive volume of ethanol at a specific location, but can be used to explore other interesting correlations.

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
$ bin/fetch all # Takes about 1hr
$ bin/tune-db
```

Optionally, explore the data with [metabase](https://metabase.com)

```sh
$ metabase
```

## Questions

- [ ] Most represented country, alcohol, etc by total volume, price adjusted, etc.
- [ ] Wine vs Beer popularity
- [ ] Geographic distributions