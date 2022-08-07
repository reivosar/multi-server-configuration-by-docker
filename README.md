# Summary
This project is an attempt to take a great design idea from the following URL and bring it to life using Docker.

https://github.com/donnemartin/system-design-primer/blob/master/README-ja.md

I use Docker because I wanted to run it without depending on AWS, GCP, etc. If practicality is a consideration, it is wise to consider other options besides Docker.

# Configuration
![image027](https://user-images.githubusercontent.com/2153822/183266236-c4f89db3-b049-4cf6-b2cc-01c10bab5450.png)

When a client accesses the following address, it is dispatched by a Nginx reverse proxy to the application running in the Docker container.

* http://localhost/
```
Dispatch to web-container:3000
```
* http://localhost/web
```
Dispatch to web-container:3000
```
* http://localhost/read-api
```
Dispatch to read-api-container:80
```
* http://localhost/write-api
```
Dispatch to write-api-container:3000
```

And what is unique about this configuration is that the Web, Read API, and Write API are each separate Docker images.

Under normal circumstances, it is the reading of the resource that is overloaded and should be managed separately from the writing to the resource. In other words, scale-up and scale-out must be able to be performed at different times and in a manner that does not affect the respective systems.

To make this possible, each Docker image should at least be separate.

# Usage
1. Clone this project
```
git clone https://github.com/reivosar/multi-server-configuration-by-docker.git
```
2. Execute the following commands
```
 docker-compose up 
```
3. Access the following URL
* http://localhost/web (WEB)
* http://localhost/read-api (READ-API)
* http://localhost/write-api (WRITE-API)
