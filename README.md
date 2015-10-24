# CMPE-272-Assignment-2
##Location and trip planner service in Go.


This is a CRUD Location Service.The location service shall have the following REST endpoints to store and retrieve locations.

####API Used:
Google Map API

##Installation
#####To run the location Crud Service you need to install
```
1. mgo driver for MongoDb driver
  >  go get gopkg.in/mgo.v2
2. httprouter
   Download HttpRouter from GitHub and Run
  > go get
```
## Usage

Clone the repository CMPE-273-Assignment-2

###Start the  server:

```
cd CMPE-273-Assignment-2
go run locationService.go
```
##Execute cURL request through shell for CRUD Operations

### 1. Create New Location

HTTP Request Used : POST

####Sample cURL command for POST Request:
```
> curl -H "Content-Type: application/json" -X POST -d '{"name" : "John Smith","address" : "123 Main St","city" : "San   Francisco","state" : "CA","zip" : "94113"}' http://127.0.0.1:8080/locations
```

####Sample Response:
```
  {
  "id": "11066",
  "name": "John Smith",
  "address": "123 Main St",
  "city": "San Francisco",
  "state": "CA",
  "zip": "94113",
  "coordinate": {
    "lat": "37.791762",
    "lng": "-122.394340"
  }
}
```
### 2. Get Existing Location

HTTP Request Used : GET

####Sample cURL command for GET Request:
```
  > curl -H "Content-Type: application/json" -X GET -d http://127.0.0.1:8080/locations/11066
```
####Sample Response:
```
 {
    "id":"11066",
    "name":"John Smith",
    "address":"123 Main St",
    "city":"San Francisco",
    "state":"CA",
    "zip":"94113",
    "coordinate":{
        "lat":"37.791762",
        "lng":"-122.394340" 
        }
  }
```
### 3. Update New Location

HTTP Request Used : PUT

####Sample cURL command for PUT Request:
```
> curl -H "Content-Type: application/json" -X PUT -d '{"address" : "1600 Amphitheatre Parkway","city" : "Mountain View","state" : "CA","zip" : "94043"}' http://127.0.0.1:8080/locations/11066
```
####Sample Response:
```
{
  "id": "11066",
  "name": "John Smith",
  "address": "1600 Amphitheatre Parkway",
  "city": "Mountain View",
  "state": "CA",
  "zip": "94043",
  "coordinate": {
    "lat": "37.422035",
    "lng": "-122.084124"
  }
}
```

### 4. Delete Location

HTTP Request Used : DELETE

####Sample cURL command for DELETE Request:
```
> curl -H "Content-Type: application/json" -X DELETE http://127.0.0.1:8080/locations/11066
```
####Sample Response:
```
---HTTP Response Code: 200
```
  


