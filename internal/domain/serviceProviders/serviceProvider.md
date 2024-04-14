# domain/serviceProviders

This section will store the interface of the infra service provider, which involves third-party API service interactions (e.g., connection and calling methods) and DB operations (e.g., DB connections and CRUD methods).

## /serviceProviders/api
* `IApiClient.go` is a API connection interface.
* `IThirdPartyApi.go` is a API interaction methods interface. (You can replace the ThirdParty name to your calling API)

## /serviceProviders/db
* `IDbConnector.go` is a DB connection interface.
* `IMysqlRepositories.go` is a DB operations interface.