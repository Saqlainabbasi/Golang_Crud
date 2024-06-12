# Golang_Crud

In this Golang Crud serve project. Includes "jwt authentication", "MySQL", "gorilla mux", "Gorm package" using "docker-compose" to run the database.

## Instructions

**Step-1:** Fork and clone the repository.

```bash
git clone https://github.com/Saqlainabbasi/Golang_Crud
```

**Step-2:** Download and Install Docker Desktop from [here](https://www.docker.com/products/docker-desktop).

**Step-3:** After installing Docker Desktop, open the terminal and go to the project directory.

```bash
cd your_project_folder
```

**Step-4:** Run the following command to build and run the mysql database.

```bash
docker-compose up --build -d
```

**Note:** Check the doocker yml file for the database credentials.username: root, password: root

**Step-5:** Run the following command to run the server.

```bash
nodemon --exec go run main.go --signal SIGTERM
```

this will run the server and watch for changes in the code and restart the server automatically.
