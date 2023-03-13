# simple-auth-sw

## About

A simple authentication service written in Go, just for learning [Service Weaver](https://github.com/ServiceWeaver/weaver).

## Getting Started

### Prerequisites

- [Docker Compose](https://docs.docker.com/compose/install/)

### Setup

1. Clone this repository.
```bash
git clone https://github.com/wst24365888/simple-auth-sw
cd simple-auth-sw
```
2. Create `db.env` file in the root directory of this project, and fill in the environment variables according to `db.env.example`.
3. Create `weaver.toml` file in the root directory of this project, and fill in the variables according to `weaver.toml.example`.
4. Start the database and the app. The app will be available at <http://localhost:8888>
```bash
docker compose up -d
```

> :warning: **IMPORTANT**: Don't build the app image like this in production, `weaver.toml` is already the deployment configuration file of Service Weaver. This is only used to creat a local linux environment.
