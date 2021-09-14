# Clean Gin

Clean architecture template with gin framework, go-fx as dependency container, gorm as orm for database related operations.

To learn about project structure and dependency injection please go through [here](https://medium.com/wesionary-team/dependency-injection-with-go-fx-b698a6585cf0?source=friends_link&sk=26f391ae41c493946ee3434be2ed4971)

## Running the project

-   Make sure you have docker installed.
-   Copy `.env.example` to `.env`
-   Run `docker-compose up -d`
-   Go to `localhost:5000` to verify if the server works.
-   [Adminer](https://www.adminer.org/) Database Management runs at `5001` .

If you are running without docker be sure database configuration is provided in `.env` file and run `go run .`

#### Environment Variables

<details>
    <summary>Variables Defined in the project </summary>

| Key            | Value                    | Desc                             |
| -------------- | ------------------------ | -------------------------------- |
| `SERVER_PORT`  | `5000`                   | Port at which app runs           |
| `ENV`          | `development,production` | App running Environment          |
| `LOG_OUTPUT`   | `./server.log`           | Output Directory to save logs    |
| `DB_USER`      | `username`               | Database Username                |
| `DB_PASS`      | `password`               | Database Password                |
| `DB_HOST`      | `0.0.0.0`                | Database Host                    |
| `DB_PORT`      | `3306`                   | Database Port                    |
| `DB_NAME`      | `test`                   | Database Name                    |
| `JWT_SECRET`   | `secret`                 | JWT Token Secret key             |
| `ADMINER_PORT` | `5001`                   | Adminer DB Port                  |
| `DEBUG_PORT`   | `5002`                   | Port that delve debugger runs in |

</details>

#### Migration Commands

> ⚓️ &nbsp; Add argument `p=host` if you want to run the migration runner from the host environment instead of docker environment.
> Check [#19](https://go_api_deploy_heroku/issues/19) for more details. eg; `make p=host migrate-up`

<details>
    <summary>Migration commands available</summary>

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

</details>


