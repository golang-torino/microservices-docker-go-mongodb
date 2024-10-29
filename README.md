# Cinema - Example of Microservices in Go with Docker, Kubernetes and MongoDB

## Overview

Cinema is an example project which demonstrates the use of microservices for a fictional movie theater.
The Cinema backend is powered by 4 microservices, all of which happen to be written in Go, using MongoDB for manage the database and Docker to isolate and deploy the ecosystem.

 * Movie Service: Provides information like movie ratings, title, etc.
 * Show Times Service: Provides show times information.
 * Booking Service: Provides booking information.
 * Users Service: Provides movie suggestions for users by communicating with other services.

The Cinema use case is based on the project written in Python by [Umer Mansoor](https://github.com/umermansoor/microservices).

The project structure is based in the knowledge learned in:

* Golang structure: <https://peter.bourgon.org/go-best-practices-2016/#repository-structure>
* Book Let's Go: <https://lets-go.alexedwards.net/>

Container images used support multi-architectures (amd64, arm/v7 and arm64).

## Index

* [Deployment](#deployment)
* [How To Use Cinema Services](#how-to-use-cinema-services)
* [Related Posts](related-posts)
* [Significant Revisions](#significant-revisions)
* [The big picture](#screenshots)

## Deployment

The application can be deployed with Docker Compose ([docs](./docs/localhost.md)).

## How To Use Cinema Services

* [endpoints](./docs/endpoints.md)

## How to observe services

Services are instrumented through OpenTelemetry. If a service is not yet instrumented, you're welcome to add it!

### Setup

Each service is already configured in Docker Compose to support otel signals collection.

Copy the `.env.example` to `.env` and provide all required information.
Pay attention to:
- updating `SERVICE_ENVIRONMENT` if multiple users are using the same otel backend, to help with filtering
- updating `OTEL_EXPORTER_OTLP_ENDPOINT` with your backend endpoint
- updating `OTEL_EXPORTER_OTLP_HEADERS` with your backend authentication token

### Useful links

- [Website](http://localhost)
- [Traefik](http://localhost:8080)
- [Grafana](http://localhost:8081)
- [Jaeger](http://localhost:8082)

## Related Posts

* [Traefik 2 - Advanced configuration with Docker Compose](https://mmorejon.io/en/blog/traefik-2-advanced-configuration-docker-compose/)

## Significant Revisions

* [Microservices - Martin Fowler](http://martinfowler.com/articles/microservices.html)
* [Umer Mansoor - Cinema](https://github.com/umermansoor/microservices)
* [Traefik Proxy Docs](https://doc.traefik.io/traefik/)
* [MongoDB Driver for Golang](https://github.com/mongodb/mongo-go-driver)
* [MongoDB Golang Channel](https://www.youtube.com/c/MongoDBofficial/search?query=golang)

## Screenshots

### Architecture

![overview](docs/images/overview.jpg)

### Homepage

![website home page](docs/images/website-home.jpg)

### Users List

![users list page](docs/images/website-users.jpg)
