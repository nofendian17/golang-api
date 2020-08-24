## Docker Usage

To get started, make sure you have [Docker installed](https://docs.docker.com/docker-for-mac/install/) on your system, and then clone this repository. Add your entire GO project to the `src` folder, then open a terminal and from this cloned respository's root run `docker-compose build && docker-compose up -d`. 

Open up your browser of choice to [http://localhost:80](http://localhost:80) and you should see your GO app running as intended. 

Containers created and their ports are as follows:

- **nginx** - `:80`
- **redis** - `:6379`
- **mongoDB** - `:6379`