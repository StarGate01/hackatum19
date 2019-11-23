#!/bin/sh

# Define default value for app container hostname and port
APP_HOST=${APP_HOST:-app}
APP_PORT_NUMBER=${APP_PORT_NUMBER:-8000}

# Linking Nginx configuration file
ln -s /etc/nginx/sites-available/mattermost /etc/nginx/conf.d/mattermost.conf

# Setup app host and port on configuration file
sed -i "s/{%APP_HOST%}/${APP_HOST}/g" /etc/nginx/conf.d/mattermost.conf
sed -i "s/{%APP_PORT%}/${APP_PORT_NUMBER}/g" /etc/nginx/conf.d/mattermost.conf

# Run Nginx
exec nginx -g 'daemon off;'
