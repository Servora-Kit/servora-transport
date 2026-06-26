module github.com/Servora-Kit/servora-transport/server/tcp

go 1.26.1

replace github.com/go-kratos/kratos/v3 v3.0.0 => github.com/go-kratos/kratos/v3 v3.0.0-20260621094049-2726761cdd77

require (
	github.com/Servora-Kit/servora v0.7.6
	github.com/go-kratos/kratos/v3 v3.0.0
	google.golang.org/protobuf v1.36.12-0.20260120151049-f2248ac996af
)

require (
	github.com/go-playground/form/v4 v4.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
