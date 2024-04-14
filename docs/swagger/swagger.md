# Swagger
I use Swagger to be our API documents.

## How to install it 
Refer to official documentation to get started: [swaggo/swag](https://github.com/swaggo/swag)

1. Firstly, add all of annotations and comments to API source code. For example, Add API service information in main.go, and add API information in controller.

2. Install Swag:
```
go install github.com/swaggo/swagger/cmd/swag@latest
```

3. Execute `swag init` in the root of your project. E.g., in this project, we need to execute `swag init` in `./go-template`. However, we need to specify the main.go location. So, we need to improve the script as:
```
swag init -g ./cmd/main.go
```
Besides, we can specify the output files location, such as:
```
swag init -g ./cmd/main.go -o ./docs/swagger/docs
```

4. Go to the router location, here I wrote in `./init/init.go`. Then import the required Swagger libraries:
```
import (
    sysinit "kaichihcodeme.com/go-template/init"
	logger "kaichihcodeme.com/go-template/pkg/zap-logger"
)

// Then add handler into router to specify the swagger routes
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

5. Go to `./cmd/main.go`, import the required dependencies:
```
import (
    _ "kaichihcodeme.com/go-template/docs/swagger/docs" // the location of output swagger documentations
)
```

## Swagger Scripts
I wrote the swagger creating and updating scripts in `./scripts/swagger.sh`.
Once you have to build up the application with swagger, you need to run this script to update the Swagger documentations.

And if you use VS Code to develop, you can refer to `./.vscode`, I specify a launch mode called: `Debug Gin w/ Swagger`, which will debug our application with update swagger via preLaunchTasks.