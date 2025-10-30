FROM golang:1.24.0-alpine as base

#-- base build image 
FROM base as build
WORKDIR /app

COPY go.* .
RUN go mod download

COPY src src
COPY migrations migrations

# Define build arguments
ARG ENVIRONMENT 


# Pass build arguments to environment variables for use in build stage.
ENV ENVIRONMENT=${ENVIRONMENT}

RUN go build src/main.go

#-- test image
FROM build as test
WORKDIR /app

CMD go test ./src/...

#-- production image
FROM base 
WORKDIR /app

COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/main ./
RUN chmod +x main

ARG ENVIRONMENT 


ENV ENVIRONMENT=${ENVIRONMENT}


CMD ./main && echo ""