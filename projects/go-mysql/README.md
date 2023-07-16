sudo docker run --name=steven-test-mysql -d mysql/mysql-server:latest

sudo docker logs mysql_docker
sudo docker logs steven-test-mysql 2>&1 | grep PASSWORD | sed 's/.*PASSWORD: //'
oDTbtk6:4X%K#1=&k72JU;n#kO875S+V

❯ sudo docker exec -it steven-test-mysql bash
bash-4.4# mysql -uroot -p
Enter password:

mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'yunshengyan';


[READ](https://hevodata.com/learn/docker-mysql/)


---

To start a Go program with MySQL in Docker Compose, you'll need to follow a few steps. Here's a step-by-step guide:

Step 1: Set up your project structure
Create a new directory for your project and navigate into it. Inside this directory, create a file named docker-compose.yml. This file will define the services for your project.

Step 2: Define the Docker Compose configuration
Open the docker-compose.yml file and define the services for your Go program and MySQL database. Here's an example configuration:

yaml
Copy code
version: '3'
services:
  mysql:
    image: mysql:latest
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: root
    volumes:
      - mysql-data:/var/lib/mysql

  go-app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    depends_on:
      - mysql
    environment:
      MYSQL_HOST: mysql
      MYSQL_PORT: 3306
      MYSQL_USER: root
      MYSQL_PASSWORD: root
      MYSQL_DATABASE: mydatabase

volumes:
  mysql-data:
In this configuration, we define two services: mysql for the MySQL database and go-app for your Go program. The mysql service uses the mysql:latest Docker image, exposes port 3306, sets the root password, and creates a named volume mysql-data to persist the data.

The go-app service builds an image using the Dockerfile in the current directory, exposes port 8080, and depends on the mysql service. It also sets environment variables that your Go program can use to connect to the MySQL database.

Step 3: Create a Dockerfile for your Go program
Create a file named Dockerfile in the same directory as the docker-compose.yml file. This file will define how to build the Docker image for your Go program. Here's an example Dockerfile:

Dockerfile
Copy code
FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
This Dockerfile sets the base image to golang:latest, sets the working directory, copies the entire project into the container, and then runs go get and go install to build the Go program. Finally, it sets the command to run your Go program (assuming the executable is named app).

Step 4: Write your Go program
Write your Go program inside the project directory. Make sure to import the necessary packages for connecting to MySQL using a library like go-sql-driver/mysql. Refer to the official documentation for details on how to connect and interact with MySQL in Go.

Step 5: Start the services
Open a terminal, navigate to your project directory, and run the following command:

Copy code
docker-compose up
Docker Compose will start the mysql and go-app services defined in the docker-compose.yml file. It will build the Go program's Docker image, set up the MySQL container, and establish the necessary connections.

Your Go program will be accessible at http://localhost:8080, and it can communicate with the MySQL database using the environment variables defined in the docker-compose.yml file.

That's it! You've successfully started a Go program with MySQL in Docker Compose.