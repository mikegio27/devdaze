# resume-app
A go API that serves my resume data from a Mongo DB.

### Running Locally
To run the application locally run:
* `docker compose up -d`
* * This starts the mongo container and the mongo express container to view the DB in a convenient UI.
* `./srcipts/load_db.py`
* * This script loads the DB with the JSON data into a collection called `resume`
* `go run src/main.go`
* * Start running the API

### K8s stuff..