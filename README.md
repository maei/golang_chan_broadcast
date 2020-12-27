# KAMeri ServiceDevice

## MISSING
- add validation for all input fields
- Permission model still need to be tested

## Minor Updates
Version  | Date | Summary
------------- | ------------- | -------------
v4 | 01.12.2020 | add connected deviceID endpoint
v3 | 26.07.2020 | add device connect/disconnect permissions 
v2 | 22.07.2020 | add created_at field to session_id
v1 | 14.07.2020 | init documentation

## Description and Functionality
The ServiceDevice is mainly responsible for managing devices 

## Postman Collection
https://www.getpostman.com/collections/0b71e1a4d8eee6704a21

### Permission Description
We can check 
- json body
- url query parameter
- url filter parameter
for permission on role level

### Device

#### Endpoints
No | Endpoint  | Type | Description| JWT | PERMISSION
------------- |------------- | ------------- | ------------- | --------------  | --------------
1 | /api/device/types  | GET | Get all device Types | YES | SUPERUSER, ADMIN, DEFAULT
2 | /api/device/get/:deviceID  | GET | Get device information for given deviceID | YES | SUPERUSER, ADMIN, DEFAULT
3 | /api/device/create  | POST | Create a new device | YES | SUPERUSER, ADMIN
4 | /api/device/auth  | POST | authenticate a device with credentials | NO | SUPERUSER, ADMIN, DEFAULT
5 | /api/device/:deviceID/connect  | POST | connect a device to a specific device | YES | SUPERUSER, ADMIN, DEFAULT
6 | /api/device/:deviceID/disconnect  | POST | disconnect a device from a specific device | YES | SUPERUSER, ADMIN, DEFAULT
7 | /api/device/:orgaID/devices  | GET | get all devices from a specific orgaID | YES | SUPERUSER, ADMIN, DEFAULT
8 | /api/internal/getReceivingNode/:sendingNodeId  | GET | get the deviceID of the connected device | NO | SUPERUSER, ADMIN, DEFAULT

##### 8 Get connected deviceID
get all infos about connected node 
###### Mandatory Fields
no json body

###### Mandatory header-information
no JWT required

###### URL
`/api/internal/getReceivingNode/:sendingNodeId`

###### Permission
1. `SUPERUSER`
2. `ADMIN` 
3. `DEFAULT`

###### Golang http request

```
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://device:5003/api/internal/getReceivingNode/b02ae8c4-33f2-11eb-a4de-00059a3c7a00"
  method := "GET"

  payload := strings.NewReader(``)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```

###### CURL example as SUPERUSER
```
curl --location --request GET 'http://localhost:5003/api/device/internal/device/b02ae8c4-33f2-11eb-a4de-00059a3c7a00' \
--data-raw ''
```

###### Response examples
*200 OK*
```
{
    "connected_device_id": "b4c343ea-33f4-11eb-8509-00059a3c7a00"
}
```

*400 Bad Request*
```
{
    "message": "please connect to a device"
}
```


*400 Bad Request*
```
{
    "message": "no documents in result"
}
```

*500 Internal Server Error*
```
{
    "message": "something failed"
}
```

##### 1 Device Types
Gives back all devices types which are selectable
###### Mandatory Fields
no json body

###### Mandatory header-information
Valid JWT must be provided with request (see CURL)

###### URL
`/api/device/types`

###### Permission
1. `SUPERUSER`
2. `ADMIN` 
3. `DEFAULT`

All roles can access this endpoint with a valid token

###### CURL example as SUPERUSER
```
curl --location --request GET '46.101.179.215:8000/api/device/types' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1OTQ2NDM4NjEsIm5iZiI6MTU5NDY0Mzg2MSwianRpIjoiZjIxZmNiODMtZmQ3OC00YmYwLThhODQtOTBiZjJiOTk5OGM2IiwiZXhwIjoxNTk0NjQ0NzYxLCJpZGVudGl0eSI6IjQxMTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNhMDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZGFiNmUwIn19.mUInDcz82p9TflmtI7nXIWxI3_pkT0n0O_x8xFx5Vfg' \
--header 'Content-Type: application/json' \
--data-raw '{
}'
```
###### Response example
```
{
    "message": {
        "0": "UNKNOWN",
        "1": "BCI_DEVICE",
        "2": "ROBOT",
        "3": "POLAR_HEART_WRIST"
    }
}
```
##### 2 Device Information
Get device information for given deviceID

###### Mandatory Fields
no json body

###### Mandatory header-information
Valid JWT must be provided with request (see CURL)

###### URL
`/api/device/get/:deviceID`
`/api/device/get/d7cb2ce5-c446-11ea-82fd-00ffb1b25530`

###### Permission
1. `SUPERUSER` -> can see every device
2. `ADMIN` -> can see device if its of his organisation
3. `DEFAULT` -> can see device if its of his organisation

###### CURL example as SUPERUSER
```
curl --location --request GET '46.101.179.215:8000/api/device/get/602d0c19-cbfd-11ea-a978-00ffb1b25530' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.
eyJpYXQiOjE1OTQ2NDM4NjEsIm5iZiI6MTU5NDY0Mzg2MSwianRpIjoiZjIxZmNiODMtZmQ3OC
00YmYwLThhODQtOTBiZjJiOTk5OGM2IiwiZXhwIjoxNTk0NjQ0NzYxLCJpZGVudGl0eSI6IjQx
MTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2
Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNh
MDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZG
FiNmUwIn19.mUInDcz82p9TflmtI7nXIWxI3_pkT0n0O_x8xFx5Vfg'
```
###### Response example
```
{
    "message": {
        "orga_id": "1",
        "device_id": "602d0c19-cbfd-11ea-a978-00ffb1b25530",
        "device_name": "TEST_1",
        "device_connected": true,
        "device_type": "BCI_DEVICE",
        "session_ids": [
            {
                "session_id": "zYAuLj7Mg",
                "created_at": "2020-07-22T09:27:43.0097199Z"
            }
        ],
        "device_rules": [
            "UNKNOWN",
            "ROBOT"
        ],
        "receiver_device_id": "asdsadsad",
        "created_at": "2020-07-22T09:26:08.889Z",
        "updated_at": "2020-07-22T09:27:16.5658406Z"
    }
}
```
###### ERROR Codes/Messages
401 - no permission
404 - no device found with given obi
500 - error while encoding payload
500 - something failed

##### 3 Create Device
Create a new device

###### Mandatory Fields
JSON-KEY  | DATATYPE | REQUIREMENT
------------- | ------------- | -------------
orga_id  | string | yes
device_name  | string | yes
password  | string | yes
device_type  | int | yes
device_rules | []int8 | yes

###### Mandatory header-information
Valid JWT must be provided with request (see CURL)

###### URL
`/api/device/create`

###### Permission
1. `SUPERUSER` -> can create a device for every organisation
2. `ADMIN` -> can create a device for his own organisation
3. `DEFAULT` -> cannot create devices

###### CURL example as SUPERUSER
```
curl --location --request POST 'http://localhost:5003/api/device/create' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1OTQ3NTc0MzcsIm5iZiI6MTU5NDc1NzQzNywianRpIjoiMTJhZDY1MTgtNmNkZS00ZWVjLWEwZDMtNzM5NDc3NzVhYmRhIiwiZXhwIjoxNTk0NzU4MzM3LCJpZGVudGl0eSI6IjQxMTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNhMDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZGFiNmUwIn19.Z_d_aMcmAky380qConK1KqdXDwUjMEkNVma6YGZZ2h0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orga_id": "2",
    "device_name":"TEST_3",
    "password":"test",
    "device_type": 1,
    "device_rules": [0,2]
}'
```
###### Response example
```
{
    "orga_id": "2",
    "device_name":"TEST_3",
    "password":"test",
    "device_type": 1,
    "device_rules": [0,2]
}
```
###### ERROR Codes/Messages
-

##### 4 Auth Device
Authenticate a device
IMPORTANT: The device has to be first connected to another device to enter this endpoint

###### Mandatory Fields
JSON-KEY  | DATATYPE | REQUIREMENT
------------- | ------------- | -------------
device_id  | string | yes
password  | string | yes

###### Mandatory header-information
no header-information

###### URL
`/api/device/auth`

###### Permission
no permissions, endpoint checks credentials provided in json-body

###### CURL example as SUPERUSER
```
curl --location --request POST 'http://localhost:5003/api/device/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "device_id":"d7cb2ce5-c446-11ea-82fd-00ffb1b25530",
    "password":"test"
}'
```
###### Response example
```
{
    "message": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uX2lkIjoic0VyUVdubkdSIiwicmVjZWl2ZXJfZGV2aWNlX2lkIjoiYXNkc2Fkc2FkIiwiZGV2aWNlX2lkIjoiZDdjYjJjZTUtYzQ0Ni0xMWVhLTgyZmQtMDBmZmIxYjI1NTMwIiwiZXhwIjoxNTk0NjAwODI5fQ.dL6aagIz42U1XmVmCCBxMjISg1gvmIRihVbgJm7iYes"
}
```
###### JWT information

```
{
  "session_id": "sErQWnnGR",
  "receiver_device_id": "asdsadsad",
  "device_id": "d7cb2ce5-c446-11ea-82fd-00ffb1b25530",
  "exp": 1594600829
}
```
###### ERROR Codes/Messages
-

##### 5 Connect Device
Connect to a device

###### Mandatory Fields
JSON-KEY  | DATATYPE | REQUIREMENT
------------- | ------------- | -------------
receiver_device_id  | string | yes

###### Mandatory header-information
Valid JWT must be provided with request (see CURL)

###### URL
`/api/device/:deviceID/connect`

###### Permission
ALL

###### CURL example as SUPERUSER
```
curl --location --request POST 'http://localhost:5003/api/device/d7cb2ce5-c446-11ea-82fd-00ffb1b25530/connect' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1OTQ3NTYwNjQsIm5iZiI6MTU5NDc1NjA2NCwianRpIjoiN2E4YmIzMGUtYTEzNS00OWUyLTg3OGEtYTIwN2YwYzFkZDljIiwiZXhwIjoxNTk0NzU2OTY0LCJpZGVudGl0eSI6IjQxMTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNhMDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZGFiNmUwIn19.QtAre_i5u1E5iKwpBHFbwik27PXcrBrzZVFjcjehYj4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "receiver_device_id": "asdsadsad"
}'
```
###### Response example
```
{
    "message": "device successfully connected"
}
```

###### ERROR Codes/Messages
ERROR 400

1 If the device type not match a 400 ERROR will thrown
```
{
    "message": "you try to connect devices that have different rules"
}
```
2 If you provide deviceIDs for sender and receiver which does not exist 
```
{
    "message": "no device found with given deviceID"
}
```
3 If you try to connect devices which are already connected
```
{
    "message": "device is already connected"
}
```
4 If you try to connect a device to a device which is already connected
```
{
    "message": "receiver_device is already connected to another device"
}
```

ERROR 500

1 If sth bad happened internally :) 
```
{
    "message": "error while changing status of the device"
}
```

##### 6 Disconnect Device
Disconnect a device

###### Mandatory Fields
No payload necessary 

###### Mandatory header-information
Valid JWT must be provided with request (see CURL)

###### URL
`/api/device/:deviceID/connect`

###### Permission
ALL

###### CURL example as SUPERUSER
```
curl --location --request POST 'http://localhost:5003/api/device/d7cb2ce5-c446-11ea-82fd-00ffb1b25530/disconnect' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1OTQ3NTYwNjQsIm5iZiI6MTU5NDc1NjA2NCwianRpIjoiN2E4YmIzMGUtYTEzNS00OWUyLTg3OGEtYTIwN2YwYzFkZDljIiwiZXhwIjoxNTk0NzU2OTY0LCJpZGVudGl0eSI6IjQxMTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNhMDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZGFiNmUwIn19.QtAre_i5u1E5iKwpBHFbwik27PXcrBrzZVFjcjehYj4'
```
###### Response example
```
{
    "message": "device successfully disconnected"
}
```

##### ERROR Codes

ERROR 400

1. If you provide a deviceID which not exist, valid etc
```
{
    "message": "no device found with given deviceID"
}
```

2. If you try to disconnect a device which is not connected yet
```
{
    "message": "device is not connected"
}
```

ERROR 500

1 If sth bad happened internally :) 
```
{
    "message": "error while changing status of the device"
}
```

##### 7 Get all devices
Get all devices for an organisation

###### Mandatory Fields
No payload necessary 

###### Mandatory header-information
Valid JWT must be provided with request (see CURL)

###### URL
`/api/device/:orgaID/devices`

###### Permission
1. `SUPERUSER` -> can show all devices for every organisation
2. `ADMIN` -> can show all devices of his organisation
3. `DEFAULT` -> can show all devices of his organisation

###### CURL example as SUPERUSER
```
curl --location --request GET 'http://localhost:5003/api/device/6/devices' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1OTQ3NTc0MzcsIm5iZiI6MTU5NDc1NzQzNywianRpIjoiMTJhZDY1MTgtNmNkZS00ZWVjLWEwZDMtNzM5NDc3NzVhYmRhIiwiZXhwIjoxNTk0NzU4MzM3LCJpZGVudGl0eSI6IjQxMTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNhMDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZGFiNmUwIn19.Z_d_aMcmAky380qConK1KqdXDwUjMEkNVma6YGZZ2h0'
```
###### Response example
```
{
    "message": [
        {
            "orga_id": "1",
            "device_id": "602d0c19-cbfd-11ea-a978-00ffb1b25530",
            "device_name": "TEST_1",
            "device_connected": true,
            "device_type": "BCI_DEVICE",
            "session_ids": [
                {
                    "session_id": "zYAuLj7Mg",
                    "created_at": "2020-07-22T09:27:43.0097199Z"
                }
            ],
            "device_rules": [
                "UNKNOWN",
                "ROBOT"
            ],
            "receiver_device_id": "asdsadsad",
            "created_at": "2020-07-22T09:26:08.889Z",
            "updated_at": "2020-07-22T09:27:16.5658406Z"
        },
        {
            "orga_id": "1",
            "device_id": "64deace2-cbfd-11ea-a978-00ffb1b25530",
            "device_name": "TEST_2",
            "device_connected": false,
            "device_type": "ROBOT",
            "session_ids": [],
            "device_rules": [
                "BCI_DEVICE"
            ],
            "receiver_device_id": "",
            "created_at": "2020-07-22T09:26:16.769Z",
            "updated_at": "0001-01-01T00:00:00Z"
        }
    ]
}
```

###### ERROR Codes/Messages
-