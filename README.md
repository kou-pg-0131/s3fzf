# s3fzf

[![CircleCI](https://circleci.com/gh/kou-pg-0131/s3fzf.svg?style=shield)](https://circleci.com/gh/kou-pg-0131/s3fzf)
[![Maintainability](https://api.codeclimate.com/v1/badges/1aa323ec22cbd6cae3d4/maintainability)](https://codeclimate.com/github/kou-pg-0131/s3fzf/maintainability)
[![codecov](https://codecov.io/gh/kou-pg-0131/s3fzf/branch/main/graph/badge.svg?token=2W5UVLK4B2)](https://codecov.io/gh/kou-pg-0131/s3fzf)
[![LICENSE](https://img.shields.io/github/license/kou-pg-0131/s3fzf?style=plastic)](./LICENSE)
[![Twitter Follow](https://img.shields.io/twitter/follow/kou_pg_0131?style=social)](https://twitter.com/kou_pg_0131)

## Overview

Fuzzy Finder for AWS S3.

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
   s3fzf <command> [options]

COMMANDS:
   cp  Copy S3 object to local.
   rm  Delete an S3 object.

GLOBAL OPTIONS:
   --help, -h  show help. (default: false)
```

### cp

```
$ s3fzf cp --help
NAME:
   s3fzf cp - Copy S3 object to local.

USAGE:
   s3fzf cp [options]

OPTIONS:
   --bucket value, -b value   name of the bucket containing the objects.
   --profile value, -p value  use a specific profile from your credential file.
   --output value, -o value   file path of the output destination. if '-' is specified, output to stdout.
   --help, -h                 show help. (default: false)
```

### rm

```
$ s3fzf rm --help
NAME:
   s3fzf rm - Delete an S3 object.

USAGE:
   s3fzf rm [options]

OPTIONS:
   --bucket value, -b value   name of the bucket containing the objects.
   --profile value, -p value  use a specific profile from your credential file.
   --no-confirm               skip the confirmation before deleting. (default: false)
   --help, -h                 show help. (default: false)
```

## LICENSE

[MIT](./LICENSE)
