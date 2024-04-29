FROM golang:1.22.2-alpine AS build
# set workdir /app and copy all files to /app
WORKDIR /app
COPY . .
# download dependencies
RUN go mod download
# swag install and init
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH="/go/bin:${PATH}"
RUN swag init -g cmd/main.go -o ./docs/swagger/docs
# mkdir /app/dist then build ./cmd (main) to /app/dist/main
RUN mkdir ./dist
RUN go build -o ./dist/main ./cmd
# copy swagger assets and config
RUN cp -r ./docs ./dist/docs
RUN cp -r ./configs ./dist/configs
RUN ls -l

# running in a final image
FROM alpine:latest AS final
# create /app directory and set as workdir
RUN mkdir /app
WORKDIR /app
# copy /app/dist from build to final /app
COPY --from=build /app/dist .
EXPOSE 8080
RUN ls -l
# add execute permission to main
RUN chmod +x ./main
# set entrypoint to ./main
ENTRYPOINT ["./main"]