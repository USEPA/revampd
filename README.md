# revAMPD

![](https://github.com/USEPA/revampd/workflows/Build/badge.svg)

This is a demo project for the re-engineering of the EPA's Air Markets Program Data (AMPD) website.

Currently, `revampd` is a static file frontend and an API backend.

## Development

### Initial setup

#### Dependencies

- [git][1]
- [docker][2]
- [docker-compose][3]

This project uses [Docker][2] containers to provide a self-contained build and test environment, as well as to locally run the `revampd` web service and static file web server. These instructions assume some familiarity with the command-line (e.g., Windows Terminal under Windows 10, Terminal under macOS, etc.).

#### Clone the repository

Clone the repository and `cd` into it:

```shell
git clone https://github.com/USEPA/revampd
cd revampd
```

#### Tests

Run the initial setup and ensure that all tests pass locally:

```shell
docker-compose run --rm api make
```

### Running local servers

Start the web server and backend API service:

```shell
docker-compose up --abort-on-container-exit --build
```

Log messages will be printed to the console's `stdout`.

Once the `revampd_web_1`, `revampd_db_1`, and `revampd_api_1` instances have started up, a web browser can be directed to `http://localhost:9080` to access the frontend. The API backend can be access directly at `http://localhost:8080`.  If there is no Postgres data directory present, the database will initialize by calling the initialization script and loading the test data. 

Access the database:

```shell
docker exec -it revampd_db_1 sh
psql -U postgres
```

### Stopping and cleaning up the local server

`Ctrl+C` will shutdown the server, if performed in the same console window as it was started. 

Using a different console window — on the same host OS, from the `revampd` project tree — the local server can be stopped and cleaned up with:
```shell
docker-compose down
```

### Adding/updating Go packages

Whenever the packages used by `revampd` change, update the [Dep][4] files:

```shell
docker-compose run --rm api dep ensure -no-vendor
docker-compose build
```

Be sure to commit the `dep`-generated updates to `src/Gopkg.toml` and `src/Gopkg.lock`.

