# Autocrack

Written at hackaTUM 2019 by
- Marko Stapfner
- Christoph Honal
- Ilias Sulgin
- Michael Gigler

## Dependencies
- `docker`
- `docker-compose`
- `bash`

## Deployment 

Outgoing ports:
- Mattermost: http://localhost:9200/
- Portainer: http://localhost:9201/
- PgAdmin: http://localhost:9202/
- Image server: http://localhost:9203/



- (Chmod: `chmod +x app.sh`)
- Build: `./app.sh build`
- Deploy: `./app.sh up`
- Shutdown `./app.sh down`

## Mattermost setup for testing localhost:9200

Go to http://localhost:9200: 

Credentials: 
- Admin-Account: 
- Email: `admin@example.com`
- Password: `AdminAdmin$1`
- Team-Name: `rinderhack`

Create channels:
- `detection`
- `alerts`

Integrations -> Incoming Webhooks -> Add:
- Title: `detection`
- Description: `detection`
- Channel: `detection`
- Lock to this channel: activated
- Set `WEBHOOK_DETECTION` in `mattermost.env` to hash key

Integrations -> Incoming Webhooks -> Add:
- Title: `alerts`
- Description: `alerts`
- Channel: `alerts`
- Lock to this channel: activated
- Set `WEBHOOK_ALERTS` in `mattermost.env` to hash key

System Console -> Developer:
- Allow untrusted internal connections to: `mattermost-connector`

## Configuration

Edit `core.env` to configure thresholds, `db.env` to configure the database access.
Edit `azure.env` to configure the Azure Accounts


## The Images that should be sent by the Camera
Place them in `./persistent-data/mock`, the `images` folder is for the already predicted images. 


## Microsoft Azure Custom Vision
Visit https://www.customvision.ai/

- Create project 
- Create resource
- Create tags:
    - cracked
    - uncracked
    
Train with some cracked and uncracked images    