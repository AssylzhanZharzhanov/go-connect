# connect

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
