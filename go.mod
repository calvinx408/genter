module github.com/openjw/genter

go 1.12

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/fatih/color v1.7.0
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/panjf2000/gnet v1.0.0-rc.3
	github.com/shiyanlou/termbox-go v0.0.0-20170327202638-7994c181db77
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/urfave/negroni v1.0.0
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/sys v0.0.0-20191105231009-c1f44814a5cd // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/utils v0.0.0-20191218082557-f07c713de883 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.24.0
