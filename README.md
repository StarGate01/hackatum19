# Rinderhack: Fancy ML Project

## Deployment steps: 

1. Fetch submodules: `git submodule update --init`

2. Build: `./app.sh build`

3. Deploy: `./app.sh up`

4. Shutdown `./app.sh down`


# mattermost-connector

The Mattermost connector uses Webhook

## 1. Mattermost-Setup 

## 2. Mattermost-Configuration

    Credentials: 

    Admin-Account erstellen: 
    Email: admin@example.com
    Passwort: AdminAdmin$1
    
    Team-Name: rinderhack
    
    Create channel: detection

## Microservice Hosts 

App:
    Host: http://localhost:9200

Portainer
    Host: http://localhost:9201

Postgres
    Host: db:5432
    User: dbuser
    Pass: dbpass
    Database: app
    Docker volume name: postgres-data

Pgadmin
    Host: http://localhost:9202
    User: pguser
    Pass: pgpass
