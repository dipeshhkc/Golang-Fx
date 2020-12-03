## Golang FX Skeleton


#### Create .env file following .env.example

#### Migration Commands

| Command            | Desc                                           |
| -------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

#### Change the configuration on docker_compose.yml as per your requirement and run the following docker commands:

- `docker-compose build` 
- `docker-compose up` 

