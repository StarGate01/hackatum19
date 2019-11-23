Build: ./app.sh build
Run: ./app.sh up
Shutdown: ./app.sh down



App:
    Host: http://localhost:9200

Portainer
    Host: http://localhost:9201

Postgres
    Host: db:5432
    User: dbuser
    Pass: dbpass
    Database: app

Pgadmin
    Host: http://localhost:9202
    User: pguser
    Pass: pgpass