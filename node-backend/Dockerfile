FROM node:13
COPY package-lock.json package.json tsconfig.json /app/
WORKDIR /app
ENV NODE_ENV development
RUN npm install
COPY ./node_modules ./node_modules
CMD [ "npm", "start" ]