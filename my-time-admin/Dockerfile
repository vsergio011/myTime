# pull official base image
FROM node:15.6.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm cache clean --force
RUN npm install
RUN npm install -g npm-check-updates
RUN npm install react-bootstrap bootstrap@4.6.0

# add app
COPY . ./

# start app
CMD ["npm", "start"]  