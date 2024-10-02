module github.com/mmorejon/microservices-docker-go-mongodb/website

go 1.19

require (
	github.com/gorilla/mux v1.8.1
	github.com/mmorejon/microservices-docker-go-mongodb/bookings v0.0.0-20231210185510-44dc85ee7569
	github.com/mmorejon/microservices-docker-go-mongodb/movies v0.0.0-20231210185510-44dc85ee7569
	github.com/mmorejon/microservices-docker-go-mongodb/showtimes v0.0.0-20231210185510-44dc85ee7569
	github.com/mmorejon/microservices-docker-go-mongodb/users v0.0.0-20231210185510-44dc85ee7569
)

require go.mongodb.org/mongo-driver v1.17.1 // indirect
