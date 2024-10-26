<p align="center">
  <img alt="Logo" src="https://raw.githubusercontent.com/janczizikow/pit/main/web/src/lib/assets/logo.svg" width="60" />
</p>
<h1 align="center">diablo4pit.web.app</h1>
<p align="center">
  Diablo 4 pit leaderboard website (discontinued)
</p>

## Rationale

In the game (as of making this project) there are no leaderboards and I was curious
to see how various classes/builds perform. Also wanted to experiment with some of the
new tech for me (golang/svelte). It was cool to see
that some people actually used the project and submitted their runs to it!
The website even ended up on google search results page for some time, but it couldn't
surpass [helltides.com](https://helltides.com/) (already quite established website among diablo 4 players).
At some point diablo4pit.web.app disappeared from google
search results, probably due to similarity of the content to helltides.com, which made me decide to
discontinue the project.

## 🚀 Quick start

Before running the project make sure you have installed:

- [go](https://go.dev/doc/install)
- [postgresql](https://www.postgresql.org/)
- [node](https://nodejs.org/en)

If you prefer to run server and database in docker, you only gonna need `node`.

```sh
go mod download
cp .example.env .env # create environment variables
createdb --owner=postgres pit # create database for development
createdb --owner=postgres pit_test # create database for tests
make run
```

To start frontend:

```sh
cd web
yarn
yarn codegen # generate files for api calls from swagger
yarn dev
```

## 🐳 Docker

Alternatively, you can start database and server in docker:

```sh
docker compose up api
```

## 🌱 Seeding database

There are various scripts to seed database in `scripts` folder.

```sh
make psql

\COPY seasons(id, "name", pit, "start", "end")
FROM 'seeds/seasons.csv'
CSV HEADER DELIMITER ',';

\COPY submissions("name", class, tier, mode, video, build, duration, verified, season_id)
FROM 'seeds/season4.csv'
WITH null as E'\'\'' CSV HEADER DELIMITER ',';
```

Check `scripts` folder for more seeds. There are also JavaScript files which can be used to
scrape data from existing websites and output a CSV file, which then can be copied to postgres.

## 👨‍⚖️ License

[MIT](LICENSE)
