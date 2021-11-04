# front_end_app

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

### How to generate grpc_web and grpc client code
- Download grpc_web package from:  https://github.com/grpc/grpc-web/releases/tag/1.2.1
- Place the file in your path, for example, on my windows machine my path is at D:/path
- Rename the file to protoc-gen-grpc-web
- cd to ac_backend directory
- protoc --proto_path=proto --proto_path=./third_party --js_out=import_style=commonjs:../ac_frontend/apps/admin/src/client --grpc-web_out=import_style=commonjs,mode=grpcweb:../ac_frontend/apps/admin/src/client user_registry.proto

*** In this case user_registry client code has generated, you might need others as well. Just change the .proto name.
*** Each time the proto file has changed, this new file must be regenerated and pushed to master.