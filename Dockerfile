FROM golang

# Setup general environment
ENV SHELL bash
ENV WORKON_HOME /usr/local/app

# Copy needed files
COPY src ${WORKON_HOME}

# Write version file
RUN echo "1.0" > version.dat

# Compile and execute the script
CMD ["go build server.go"]
CMD ["./server"]

# Open port
EXPOSE 1080