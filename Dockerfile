#The base Go Image
FROM golang:1.17-alpine

#Create a directory for the app
RUN mkdir /app

#Copy all files from the current directory to the app directory
ADD . /app

#Set Working directory
WORKDIR /app

#COPY THE DEPENDANCIES
COPY go.mod .
COPY go.sum .

#DOWNLOAD DEPENDANCIES
RUN go mod download 

COPY . .


#RUN command as described:
#go build will build an executable file named server in the current directory
RUN go build -o main .

#Expose the Port 
EXPOSE 8382

# RUN the server executable
CMD ["/app/main"]