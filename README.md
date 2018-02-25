# api-um-warsaw-client

Simple client for Warsaw UM Api. More information is available here https://api.um.warszawa.pl/

### Compilation:
```
go get github.com/kaweue/api-um-warsaw-client
```

### Sample usage:
```
api-um-warsaw-client.exe --api-key xxxx-xxxx-xxxx-xxxx-xxx getBusStop znana
api-um-warsaw-client.exe --api-key xxxx-xxxx-xxxx-xxxx-xxx getLinesAtBusStop 5104 01
api-um-warsaw-client.exe --api-key xxxx-xxxx-xxxx-xxxx-xxx getTimeTable 5104 01 155
```
