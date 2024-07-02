# exoplante_microservices
Develop the Golang microservice for managing exoplanets as described, we follow a structured approach. We will use the go-gonic/gin framework for handling HTTP requests, PostgreSQL for data persistence, and implement unit tests for the controller, service, database layers and docker file for deployment.


Installation
To run this microservice locally, follow these steps:

Clone the repository:

bash
Copy code
git clone https://github.com/your_username/exoplanet_microservices.git
cd exoplanet_microservices
Install dependencies:

Ensure you have Go installed. Then, install required dependencies:

bash
Copy code
go mod download
Set up PostgreSQL:

Install PostgreSQL and create a database.
Update the database connection details in config/db.go.
Run the application:

bash
Copy code
go run main.go
The server will start at http://localhost:8080s