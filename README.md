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
  - [Version](#version)
  - [Step](#step)
- [Contributing](#contributing)

## Prerequisites

### Engine

The game engine is written in [Go](https://go.dev/) and uses the version `1.21`.


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

The game UI runs on [Node.js](https://nodejs.org/) version `20.10.0`.


Install the dependencies inside the `ui` directory with:
```shell
npm install
```

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

## Game UI

The game UI can be found in the `ui` directory. It uses [Solid](https://www.solidjs.com/) and [Tailwind CSS](https://tailwindcss.com/) for styling.

The game UI contains several scripts to lint, format, build and run the project. To check the available scripts, run the following command inside the `ui` directory:
```shell
npm run
```

## WASM API

The WASM binary exports two functions to the global JavaScript object through a property called `engine`. These functions are described in the following sections.

### Version

The `engine.version()` function returns the information about the version of the game engine build:
```jsonc
{
    "go": "1.21.0",                            // Golang version used to build the engine.
    "tag": "v1.0.0",                           // Git tag used to build the engine. 
    "commit": "bd1557f",                       // Hash of the Git commit used to build the engine. 
    "build": "Sun 25 Feb 2024 09:11:20 PM GMT" // Timestamp when the engine was built.
}
```

### Step

The `engine.step()` function is responsible for simulating a game step of the game world physics and behaviors based on player actions. It takes the following argument, which represents the user input:
```jsonc
[
    [
        "Left", // Name of the player action being performed in the current step.
        false   // Status of the player action being performed in the current step.
    ],
    [
        "Right",
        true
    ],
    [
        "Jump",
        true
    ]
]
```

It returns the following structure:
```jsonc
{
    "error": null,     // String of the error that occurred when the step was executed, or null if no error occurred.
    "camera": {        // Configuration of the camera.
        "position": {
            "x": 0.0,
            "y": 0.0
        },
        "rotation": 0.0,
        "scale": {
            "x": 1.0,
            "y": -1.0
        },
        "width": 1000.0,
        "height": 844.0,
        "ppu": 1.0
    },
    "gameObjects": [   // List of game objects present in the camera.
        {
            "id": 1,
            "active": true,
            "tag": "Player",
            "transform": {
                "position": {
                    "x": 500.0,
                    "y": 592.5
                },
                "rotation": 0.0,
                "scale": {
                    "x": 1.0,
                    "y": -1.0
                }
            },
            "rigidBody": {
                "bodyType": 0,
                "drag": 1.0,
                "velocity": {
                    "x": 0.0,
                    "y": -1.74
                },
                "angularVelocity": 0.0,
                "gravityScale": 1.0,
                "collisionDetection": 0,
                "interpolation": 1,
                "autoMass": false,
                "mass": 10.0,
                "angularDrag": 0.0
            },
            "renderer": {
                "width": 96.0,
                "height": 96.0,
                "offset": {
                    "x": -48.0,
                    "y": -48.0
                },
                "layer": "default",
                "image": "images/player/idle/0.png",
                "flipHorizontally": true
            },
            "collider": {
                "density": 1.0,
                "points": [
                    {
                        "x": 0.0,
                        "y": 0.0
                    },
                    {
                        "x": 42.0,
                        "y": 0.0
                    },
                    {
                        "x": 42.0,
                        "y": 84.0
                    },
                    {
                        "x": 0.0,
                        "y": 84.0
                    }
                ],
                "offset": {
                    "x": -21.0,
                    "y": -45.5
                },
                "isTrigger": false,
                "layer": "",
                "material": {
                    "elasticity": 0.0,
                    "friction": 0.2
                }
            }
        },
        {
            "id": 2,
            "active": true,
            "tag": "Platform",
            "transform": {
                "position": {
                    "x": 308.0,
                    "y": 662.0
                },
                "rotation": 0.0,
                "scale": {
                    "y": -1.0,
                    "x": 1.0
                }
            },
            "renderer": {
                "width": 48.0,
                "height": 48.0,
                "offset": {
                    "y": -24.0,
                    "x": -24.0
                },
                "layer": "default",
                "image": "images/platform/forest/grass/3.png",
                "flipHorizontally": null
            }
        }
    ]
}
```

## Contributing

### Branches

- The [main branch](https://github.com/GOOFR-Group/jump-master/tree/main) contains the production code for the game
- To develop a new feature or fix a bug, a new branch should be created based on the main branch

### Issues

- The features and bugs should exist as a [GitHub issue](https://github.com/GOOFR-Group/jump-master/issues) with an appropriate description
- The status of the issues can be seen in the associated [GitHub project](https://github.com/orgs/GOOFR-Group/projects/3/views/4)

### Commits

- Git commits should follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)

### Pull Requests

- To merge the code into production, a pull request should be opened following the existing template to describe it, as well as the appropriate labels
- To merge the pull request, the code must pass all GitHub action checks, be approved by at least one of the code owners, and be squashed and merged

### Deployments

- After the code is merged into the main branch, there is a GitHub action that automatically builds and deploys the code to production

### Releases

- To release a new version of the project, [semantic versioning](https://semver.org/) should be followed
