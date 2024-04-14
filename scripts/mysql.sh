#bin/bash
docker pull mysql:latest
docker images | grep mysql

docekr run --name mysql-cafe -p 3306:3306 -e MYSQL_ROOT_PASSWORD="$MYSQL_ROOT_PASSWORD" -d mysql:latest
docker ps