<p>
pull the project into the local system.<br>
Ensure Docker is installed in the system.<br>
open cmd/ terminal in the project folder.
</p>
<br>
Run the following commands inside the project folder:
<ul>
<li>docker build -t go-planets .</li>
<li>docker run -p 8080:8080  go-planets</li>
</ul>
<br>
To test the API, import the "go-planets.postman_collection.json" file in Postman, and run the request to explore.
