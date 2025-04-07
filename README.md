# Exinity | Symbol Data Provider

# Project Layout 

## Project Milestones (metric: x=hour)
- Understanding the needs and design architecture as a template (1x)
- Development of the architecture and apply the edge cases (2x)
- Development of services (4x)
- Testing and Documentation (1x)
- Terraform & K8S Files (1x)
- Deployment and get ready for sharing (1x)

## Project ToDo's
- [x] Understanding the needs and requirements.
- [x] Create a basic project layout and git repo.
- [x] System design and skeleton development.
- [x] Docker integration database connection
- [x] Database and schema design
- [x] Development of internal services
- [x] Development of worker processor 
- [x] Development of gRPC server
- [x] Development of unit tests
- [x] Terraform & K8S files.
- [x] Documentation
- [x] Sharing with Exinity Team.

## Missing parts due to time constraint
- Unit test cases can not completed well due to time constraint
- Integration, benchmark, contract, e2e tests are also could not be added due to time constraint
- I tried to keep the design as simple as I can due to time constraint.

# How to use ?

## Accessing the project 

You can access the project via [github](https://github.com/AkyurekDogan/exinity-task). 

## How to run ?

### Step-1

You can run the code in different ways however, in every case you need the following setup 

- `install docker` if you do not have

### Step-2

You can choose one of the following options in `Step-2`

#### Running the code (Recommended)

- `install the golang version ~1.23.2`
- You can use any IDE to see the code however, I used `VSCode` so if you are `VSCode` user, you can use the `.vscode/lunch` configuration to run the code in `debug mode`
- Run the following make file command (you can see the makefile at the root directory of the code directory) `make run-local-docker-db`. This command will initiate local docker `postgres` instance with some initial data which is defined in the `init.sql` file in the `/scripts/database` folder.
- The database connection is ready at `localhost:5432` with `username:postgress` and `password:mypassword123!`. You can also see these details in `./config.yml` file
- Both you can run the `make run` command to run the service or use the IDE debugger functionality.
- The service will be running on `localhost:1989`
- You can check the `http://localhost:1989/swagger/index.html` url to access the swagger file or you can use the postman or curl collections that defined in another steps.

#### Using as a docker compose

- You can see the `docker-compose.yml` file in the root directory of the repository. 
- I recommend to change the database connection in the `./config.yml` file `database.host: go-exinity-task-postgress` so the service can access in the compose base port and host.
- run the `make run-docker-compose` command 
- The service will be running on the `localhost:3000` (ports are used different to understand the traffic root better)

### Step-3

Depending on the previous steps the service must be running on either `localhost:1989` or `localhost:3000` so you can use the following commands to use the service 

- `http://localhost:1989/partner?id={partner_id} GET`
- `http://localhost:1989/match?material_type={material_type}&lat={lat}&long={long} GET`

### Step-4

You can see the definitions of the services as follows;

#### HTTP:GET /partner 

Example `curl` command 
```
curl --location 'http://localhost:1989/partner?id=pk12kkk123'
```

##### 200 OK
```
{
    "status_code": 200,
    "message": "OK",
    "data": {
        "id": "pk12kkk123",
        "name": "K Engineering - Neukoln/Berlin",
        "location": {
            "lat": 52.475945,
            "long": 13.446991
        },
        "radius": {
            "value": 30,
            "metric": "km"
        },
        "rating": {
            "value_avg": 5
        },
        "skills": [
            "wood",
            "carpet"
        ]
    }
}
```

##### 400 Bad Request
```
{
    "status_code": 400,
    "message": "[Bad Request] invalid or insufficient input",
    "error": [
        "0-partner id must be provided"
    ]
}
```

#### HTTP:GET /match 

Example `curl` command 

```
curl --location 'http://localhost:1989/match?material_type=wood&lat=52.451320&long=13.337652'
```

##### 200 OK
```
{
    "status_code": 200,
    "message": "OK",
    "data": {
        "filter": {
            "material_type": "wood",
            "location": {
                "lat": 52.45132,
                "long": 13.337652
            }
        },
        "matches": [
            {
                "partner_id": "py12yyy123",
                "name": "Y Engineering - Spandau/Berlin",
                "location": {
                    "lat": 52.537976,
                    "long": 13.204127
                },
                "radius": {
                    "value": 30,
                    "metric": "km"
                },
                "distance": {
                    "value": 17.5,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 10
                },
                "skills": [
                    "wood",
                    "tiles"
                ],
                "rank": 1
            },
            {
                "partner_id": "pl12lll123",
                "name": "L Engineering - Prenzlauer/Berlin",
                "location": {
                    "lat": 52.540394,
                    "long": 13.423423
                },
                "radius": {
                    "value": 30,
                    "metric": "km"
                },
                "distance": {
                    "value": 13.53,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 9
                },
                "skills": [
                    "wood",
                    "carpet"
                ],
                "rank": 2
            },
            {
                "partner_id": "po12ooo123",
                "name": "O Engineering - Lichtenberg/Berlin",
                "location": {
                    "lat": 52.535544,
                    "long": 13.497926
                },
                "radius": {
                    "value": 30,
                    "metric": "km"
                },
                "distance": {
                    "value": 19.94,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 9
                },
                "skills": [
                    "wood",
                    "carpet"
                ],
                "rank": 3
            },
            {
                "partner_id": "pt12ttt123",
                "name": "T Engineering - Schoneberg/Berlin",
                "location": {
                    "lat": 52.49092,
                    "long": 13.359833
                },
                "radius": {
                    "value": 10,
                    "metric": "km"
                },
                "distance": {
                    "value": 4.94,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 8
                },
                "skills": [
                    "wood",
                    "carpet",
                    "tiles"
                ],
                "rank": 4
            },
            {
                "partner_id": "pz12zzz123",
                "name": "Z Engineering - Hauptbahnhof/Berlin",
                "location": {
                    "lat": 52.526804,
                    "long": 13.365678
                },
                "radius": {
                    "value": 20,
                    "metric": "km"
                },
                "distance": {
                    "value": 8.75,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 7
                },
                "skills": [
                    "wood",
                    "tiles"
                ],
                "rank": 5
            },
            {
                "partner_id": "pn12nnn123",
                "name": "N Engineering - Adlershof/Berlin",
                "location": {
                    "lat": 52.437278,
                    "long": 13.534422
                },
                "radius": {
                    "value": 30,
                    "metric": "km"
                },
                "distance": {
                    "value": 21.82,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 7
                },
                "skills": [
                    "wood",
                    "carpet",
                    "tiles"
                ],
                "rank": 6
            },
            {
                "partner_id": "pp12ppp123",
                "name": "P Engineering - Rodow/Berlin",
                "location": {
                    "lat": 52.422061,
                    "long": 13.495607
                },
                "radius": {
                    "value": 30,
                    "metric": "km"
                },
                "distance": {
                    "value": 17.76,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 6
                },
                "skills": [
                    "wood",
                    "carpet",
                    "tiles"
                ],
                "rank": 7
            },
            {
                "partner_id": "pk12kkk123",
                "name": "K Engineering - Neukoln/Berlin",
                "location": {
                    "lat": 52.475945,
                    "long": 13.446991
                },
                "radius": {
                    "value": 30,
                    "metric": "km"
                },
                "distance": {
                    "value": 12.39,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 5
                },
                "skills": [
                    "wood",
                    "carpet"
                ],
                "rank": 8
            },
            {
                "partner_id": "pb12bbb123",
                "name": "B Engineering - Falkensee/Brendenburg",
                "location": {
                    "lat": 52.560051,
                    "long": 13.078569
                },
                "radius": {
                    "value": 40,
                    "metric": "km"
                },
                "distance": {
                    "value": 30.99,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 5
                },
                "skills": [
                    "wood",
                    "carpet",
                    "tiles"
                ],
                "rank": 9
            },
            {
                "partner_id": "px12xxx123",
                "name": "X Engineering - Gesundrunen/Berlin",
                "location": {
                    "lat": 52.550385,
                    "long": 13.380968
                },
                "radius": {
                    "value": 20,
                    "metric": "km"
                },
                "distance": {
                    "value": 11.75,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 4
                },
                "skills": [
                    "wood",
                    "carpet",
                    "tiles"
                ],
                "rank": 10
            },
            {
                "partner_id": "pc12ccc123",
                "name": "C Engineering - Hennigsdorf/Brendenburg",
                "location": {
                    "lat": 52.645989,
                    "long": 13.199716
                },
                "radius": {
                    "value": 40,
                    "metric": "km"
                },
                "distance": {
                    "value": 26.04,
                    "metric": "km"
                },
                "rating": {
                    "value_avg": 4
                },
                "skills": [
                    "wood",
                    "carpet",
                    "tiles"
                ],
                "rank": 11
            }
        ]
    }
}
```

##### 400 Bad Request

```
{
    "status_code": 400,
    "message": "[Bad Request] invalid or insufficient input",
    "error": [
        "0-material type must be provided",
        "1-lattitute must be provided",
        "2-longtitute must be provided"
    ]
}
```

