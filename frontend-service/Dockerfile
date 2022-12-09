FROM node:19-alpine3.15

WORKDIR /app
COPY package.json .
RUN apk add --no-cache git
RUN npm install
COPY . .
EXPOSE 8085
CMD ["npm", "start"]