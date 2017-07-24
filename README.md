# mmc

mmc is a simple commandline currency converter.

# Installation

```sh
go get -u github.com/NARKOZ/mmc
```

# Usage

The usage is similar to Google's Currency converter.

Run from your terminal (case insensitive):

```
mmc 100 USD to AUD
```

or:

```
mmc 12.5 btc in usd
```

The first argument is an amount for conversion, the second and last arguments
are currency codes.

For a list of supported 150+ currencies and their respective codes see
[`data/currencies.json`](https://github.com/NARKOZ/mmc/blob/master/data/currencies.json)

# License

Released under the BSD 2-clause license. See LICENSE.txt for details.
