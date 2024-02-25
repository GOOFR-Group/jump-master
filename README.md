# Jump Master

Jump Master is a simple game where your only objective is to get to the top. Just try not to fall and do not get angry.

### Player Actions

| Action              | Description           |
|---------------------|-----------------------|
| Left Arrow          | Move left             |
| Right Arrow         | Move right            |
| Space               | Jump (can be charged) |
| Space + Left Arrow  | Jump left             |
| Space + Right Arrow | Jump right            |

# Development

The following section focuses on the development part of the game, including prerequisites, how to build and run the code, and how to contribute.

### Table of Contents
- [Prerequisites](#prerequisites)
  - [Engine](#engine)
  - [UI](#ui)
- [Web App](#web-app)
- [Game Engine](#game-engine)
- [Game UI](#game-ui)
- [WASM API](#wasm-api)
- [Contributing](#contributing)

## Prerequisites

### Engine

The game engine is written in Go and uses the version `1.21`.


The `GOPRIVATE` environment variable should be set to download the private go modules.  
Allow all the private repositories of the organization:
```shell
go env -w GOPRIVATE=github.com/goofr-group
```


In order for Go to successfully access the Git server, you should use SSH.  
To do this, [configure SSH](https://docs.github.com/en/authentication/connecting-to-github-with-ssh) in your GitHub account and use SSH instead of HTTPS in your global Git configuration:
```shell
git config --global url."git@github.com:goofr-group".insteadOf https://github.com/goofr-group
```

### UI

// TODO

## Web App

The game can be built and run in a web application that uses a WASM created by the game engine. To simplify this process, there is a makefile in the root directory that contains the following commands.


To build the game web app into the `dist` directory:
```shell
make build
```


To run the game locally:
```shell
make dev
```


For other commands:
```shell
make help
```

## Game Engine

The game engine can be found in the `engine` directory. It contains the Go code that runs the game physics simulation and behaviors. It builds a WASM binary that is used by the game UI to send the player input and render the objects in the world.

The game engine contains a makefile that defines a set of tasks that can be run to help with development. The available targets can be checked by running the following command inside the `engine` directory:
```shell
make help
```

### Structure

## Game UI

## WASM API

## Contributing

- The [main branch](https://github.com/GOOFR-Group/jump-master/tree/main) contains the production code for the game
- To develop a new feature or fix a bug, a new branch should be created based on the main branch
- The features and bugs should exist as a [GitHub issue](https://github.com/GOOFR-Group/jump-master/issues) with an appropriate description
- The status of the GitHub issues can be seen in the associated [GitHub project](https://github.com/orgs/GOOFR-Group/projects/3/views/4)
- Git commits should follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
- To merge the code into production, a pull request should be opened following the existing template to describe it, as well as the appropriate labels
- To merge the pull request, the code must pass all GitHub action checks, be approved by at least one of the code owners, and be squashed and merged
- After the code is merged into the main branch, there is a GitHub action that automatically builds and deploys the code to production
- To release a new version of the project, [semantic versioning](https://semver.org/) should be followed
