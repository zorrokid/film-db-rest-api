# Running with docker-compose

To run environment with docker-compose

    docker-compose up -d

To tear down

    docker-compose down

# Build docker container image

To build docker container image

    docker build -t "film-db-rest-api" .

To run docker container image

    docker run -dp 8080:8080 film-db-rest-api

# To test api

    curl localhost:8080/movies