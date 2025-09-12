# Smart Mirror

## Requirements

You would need

- `git` and `docker` installed

## Get started

1. clone this repo
1. run `make`
1. smart-mirror services will be started exposed on `http://localhost`

## Commands

This repo includes a Makefile that abstracts all commands.

| Command                 | Desc                                 |
| ----------------------- | ------------------------------------ |
| `make`                  | starts all services                  |
| `make start`            | starts all services                  |
| `make stop`             | stops all services                   |
| `make rebuild`          | rebuild all services after an update |
| `make start-frontend`   | start only frontend                  |
| `make start-backend`    | start only backend                   |
| `make stop-frontend`    | stop only frontend                   |
| `make stop-backend`     | stop only backend                    |
| `make rebuild-frontend` | rebuild only frontend                |
| `make rebuild-backend`  | rebuild only backend                 |
| `make open`             | open app in browser                  |
| `make close`            | close browser                        |

## Development

For local development you just run `docker compose up` which will start the frontend and backend services with hot-reloading.
