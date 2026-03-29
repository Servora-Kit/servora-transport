# servora-transport/server/tcp

TCP server plugin for Servora transport runtime.

## Install

```bash
go get github.com/Servora-Kit/servora-transport/server/tcp@latest
```

## Config Scan (in main)

Use Servora bootstrap runtime scanner directly in `BootstrapAndRun` callback:

```go
err := bootstrap.BootstrapAndRun(flagconf, Name, Version, func(rt *bootstrap.Runtime) (*kratos.App, func(), error) {
    tcpCfg, err := bootstrap.ScanConf[tcpconf.Server](rt)
    if err != nil {
        return nil, nil, err
    }
    _ = tcpCfg
    return wireApp(...)
})
```

## Generate proto

```bash
make gen
```

## Test

```bash
make test
```
