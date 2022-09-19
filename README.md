# connect [IN PROGRESS]

#### Full list what has been used:
[Gin](https://github.com/gin-gonic/gin) web framework <br/>
[Elasticsearch](https://github.com/elastic/go-elasticsearch) client for Go <br/>
[Kafka](https://github.com/segmentio/kafka-go) as messages broker<br/>
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed [tracing](https://opentracing.io/) <br/>
[Prometheus](https://prometheus.io/) monitoring and alerting <br/>
[Kibana](https://github.com/labstack/echo) is user interface that lets you visualize your Elasticsearch <br/>
[Docker](https://www.docker.com/) and docker-compose <br/>

## Project structure:
1. /api
    - /proto - contains application proto files
    - /generated - contains generated application and third-party proto files, generation rules should be described in docker-compose.protoc.yml
2. /cmd
    - /main.go - main application starter
3. /internal - contains the application codebase
4. /migrations - contains database migrations
5. /scripts - useful scripts
6. /pkg - contains codebase that could be migrated to another projects
7. version
    - version.sh - scripts for increasing the version of the application that is called from the pipeline
    - version - contains the current version of the application
