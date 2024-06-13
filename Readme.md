pull the project into the local system.
Ensure Docker is installed in the system.
open cmd/ terminal in the project folder.
--------------------------------------------------------------------
Run the following commands inside the project folder:
docker build -t go-planets .
docker run -p 8080:8080  go-planets
--------------------------------------------------------------------
To test the API, import the "go-planets.postman_collection.json" file in Postman, and run the request to explore.