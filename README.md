# swapi-cli
Golang SWAPI (Star Wars API) cli to output starship and pilot data filtered by film

## Usage
```
Usage of swapi-cli
  -h help
  -film string
    	name of the film in which to generate starship and pilot data (default "Episode IV: A New Hope")
```

## Makefile
Makefile contains commands to build, test, build the docker container, and run.

## Dockerhub
```
docker pull cushind/swapi-cli:tagname
```

## Helm
Please refer to `deployments/helm/swapicli/README.md` for more info.

## Films
- The Phantom Menace
- Attack of the Clones
- Revenge of the Sith
- A New Hope
- The Empire Strikes Back
- Return of the Jedi
- The Force Awakens

## launch.json for VSCode
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": [
                "--film",
                "Return of the Jedi",
            ]
        }
    ]
}
```