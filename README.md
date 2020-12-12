# s3fzf

[![CircleCI](https://circleci.com/gh/kou-pg-0131/s3fzf.svg?style=shield)](https://circleci.com/gh/kou-pg-0131/s3fzf)
[![Maintainability](https://api.codeclimate.com/v1/badges/1aa323ec22cbd6cae3d4/maintainability)](https://codeclimate.com/github/kou-pg-0131/s3fzf/maintainability)
[![codecov](https://codecov.io/gh/kou-pg-0131/s3fzf/branch/main/graph/badge.svg?token=2W5UVLK4B2)](https://codecov.io/gh/kou-pg-0131/s3fzf)
[![LICENSE](https://img.shields.io/github/license/kou-pg-0131/s3fzf?style=plastic)](./LICENSE)
[![Twitter Follow](https://img.shields.io/twitter/follow/kou_pg_0131?style=social)](https://twitter.com/kou_pg_0131)

## Overview

Fuzzy Finder for AWS S3.

![preview](./demo/preview.gif)

## Installation

```
$ go get -u github.com/kou-pg-0131/s3fzf
```

## Usage

```
$ s3fzf --help
NAME:
   s3fzf - Fuzzy Finder for AWS S3.

USAGE:
   s3fzf [global options]

GLOBAL OPTIONS:
   --bucket value, -b value   The name of the bucket containing the objects
   --profile value, -p value  Use a specific profile from your credential file
   --output value, -o value   File path of the output destination
   --help, -h                 show help (default: false)
```

## LICENSE

[MIT](./LICENSE)
