# FROM node:carbon
# Create app directory
# WORKDIR /usr/src/app
# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
# COPY . ./
# RUN npm install
# If you are building your code for production
# RUN npm install --only=production

FROM golang:onbuild
 
#RUN mkdir /app 
#ADD . /app/ 
#WORKDIR /app 
#RUN go build -o main . 
#CMD ["/app/main"]
