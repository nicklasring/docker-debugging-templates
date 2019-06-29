# Installation

```
docker-compose build
docker-compose up -d
```

Set the IP to the docker container in the launch.json file

## Attaching the debugger
Set a breakpoint in vscode and start debugging.

## Restarting the application
The debugger might not re-attach depending on how your application works. If this problem occurs simply restart the application so it can wait until the debugger is re-attached
```
docker-compose exec app bash -c "supervisorctl restart app"
```
