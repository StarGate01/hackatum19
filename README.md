# Rinderhack: Fancy ML Project

## Deployment steps: 

1. Fetch submodules: `git submodule update --init`

2. Chmod: `chmod +x app.sh`

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
    
    Integrations -> Incoming Webhooks -> Add:
        Title: detection
        Description: detection
        Channel: detection
        Lock to this channel: activated

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

Data server
    Host: http://localhost:9203
    Mount path: /data/images