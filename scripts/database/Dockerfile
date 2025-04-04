# Use the official PostgreSQL image from the Docker Hub
FROM kartoza/postgis:15

# Add the init.sql script to the Docker image
COPY init.sql /docker-entrypoint-initdb.d/

# Set environment variables for PostgreSQL
ENV POSTGRES_DB=exinity_task
ENV POSTGRES_USER=postgress
ENV POSTGRES_PASSWORD=mypassword123!

# To build and run this container run the following commands

# docker build -t go-stackit-cost-calculator-postgress .
# docker run --name go-stackit-cost-calculator-postgress -p 5432:5432 -d go-stackit-cost-calculator-postgress

