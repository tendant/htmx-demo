# Start postgres using docker

    cd docker/
    ./start-postgres.sh
    
# Create database 

    CREATE Role htmx_demo WITH PASSWORD 'pwd';
    grant htmx_demo to postgres;
    CREATE DATABASE htmx_demo_db ENCODING ‘UTF8’ OWNER htmx_demo;
    GRANT ALL PRIVILEGES ON DATABASE htmx_demo_db TO htmx_demo;
    ALTER ROLE htmx_demo WITH LOGIN;
    
# Migrate database 

    make migrate

    