FROM golang:latest

# Setup general environment
ENV SHELL bash
ENV WORKON_HOME /usr/local/app

# Make project folder
RUN mkdir WORKON_HOME

# Copy all files
ADD . ${WORKON_HOME}

# Change work directory
WORKDIR ${WORKON_HOME}

# Write version file
RUN echo "1.0" > version.dat

# Compile and execute the script
RUN go build server.go
CMD ["./server"]

# Open port
EXPOSE 1080