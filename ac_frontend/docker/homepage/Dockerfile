FROM node:lts

WORKDIR /app
COPY /apps/homepage /app

RUN npm install

EXPOSE 3000
CMD [ "npm", "start" ]