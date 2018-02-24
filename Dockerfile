FROM julia:latest

# Install JuliaFEM package and it's dependencies
RUN apt-get update && \
    apt-get -y install hdf5-tools build-essential node
RUN julia -e 'Pkg.add("JuliaFEM")'

# Get golang image
FROM golang

# Copy all files
COPY ./main.go /go/src/github.com/juliafem/
COPY ./manta /go/src/github.com/juliafem/manta

# Set workdir
WORKDIR /go/src/github.com/juliafem/

# Build
RUN go get ./... && go build -o MantaUI
RUN cd /manta/frontend
RUN npm build

# Run command
CMD ./MantaUI
