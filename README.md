# Go/Gin Design-Driven Design Style Template

## Introduction
This is a repository for creating a DDD style Go template.

Here implement some simple APIs to query MySQL DB and return the result from DB.
And here would apply docker to pack Go API and DB in local. So you can try this project by your local site.

## Features
* A DDD style Go/ Gin API Backend
* Support Swagger (OpenAPI)
* Support Log Middleware
* R/W with MySQL Database via APIs

## Requirements
* Go 1.22.1

## Backend Architecture
Our backend follows the DDD (Domain-Driven Design) structure to construct our Go lang backend system.

## Initialize Go Project
1. Create a project folder
2. Ensure you have Go installed and required versions (check by `go version` in your cli)
3. Use this commad to initialize your project in your project folder: ```go mod init <project-owner>/<project-name>```

## Packages
* Gin
* Viper -> configurations
* Zap -> logger, and open to other repo to use

## Make steps
build: `make build`
run: `make run`
build&run: `make br`