# Installation

```
docker-compose build
docker-compose up -d
```

Set the IP to the docker container in the launch.json file

## Attaching the debugger
Set a breakpoint in vscode and start debugging.

## Restarting the application
If you attach the debugger to the application it needs to be restarted for it to be able to re-attach.
```
docker-compose exec app bash -c "supervisorctl restart app"
```
