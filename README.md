# iridium

![](https://github.com/bradj/iridium/workflows/CI/badge.svg)

## Requirements

1. [sql-migrate](https://github.com/rubenv/sql-migrate)
1. [sqlboiler v4](https://github.com/volatiletech/sqlboiler)

## Getting Started

1. `git clone git@github.com:bradj/iridium.git`
1. Copy `sqlboiler.toml.example` to `sqlboiler.toml` and change values accordingly
1. Copy `config.toml.example` to `config.toml` and change values accordingly
1. `make db-start`      - start postgres docker container
1. `make db-migrate`    - migrate container to latest db version
1. `make db-generate`   - generate CRUD structs ([sqlboiler](https://github.com/volatiletech/sqlboiler))
1. `make`               - build Iridium binary
1. `make run`           - run Iridium binary
