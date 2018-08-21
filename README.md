# WarpKrub
## Intro
### This is for learning purpose only. But feel free to report us if this repo should be removed, add more warps also.

## Docker Build
```sh
docker build -t warpkrub . -f Dockerfile.multi-stg
```

## Docker run
```sh
docker run -p 8080:8080 warpkrub 
```
Your warps portal will ready at `localhost:8080`
## WarpKrub
`/` Welcome page

`/version` Display current version

`/warps` Shows all warps