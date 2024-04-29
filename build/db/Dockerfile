FROM mysql:latest

# Add a database
ENV MYSQL_DATABASE cafe

# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./sql/init.sql /docker-entrypoint-initdb.d/