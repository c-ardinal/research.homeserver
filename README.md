# 研究用簡易スマートホームブリッジ

## build&run
Golang version: 1.8.5
```
git clone https://github.com/c-ardinal/research.homeserver
cd research.homeserver
go get -d -v
go build
./research.homeserver example.json
```

## REST APIs
|Implemented|Method|Path|Description|Example URL|
|:---------:|:----:|:--:|:---------:|:-----:|
|○| GET  |/devices            |機器情報一覧を取得|http://localhost:8080/devices|
|○| GET  |/device/{deviceName}|機器情報を取得|http://localhost:8080/device/living-aircon|
|  | POST |/device/{deviceName}|機器の追加|http://localhost:8080/device/living-aircon|
|  | PUT  |/device/{deviceName}|機器情報の修正|http://localhost:8080/device/living-aircon|
|  |DELETE|/device/{deviceName}|機器の削除|http://localhost:8080/device/living-aircon|
|○| POST |/control            |機器の操作|http://localhost:8080/control|
|  | POST |/scan               |機器の自動スキャン|http://localhost:8080/scan|

## Example request body
 - POST /control
 ```json
 {
     "operations": [
         {
             "target": "living-aircon",
             "operation": "power_on"
         },
         {
             "target": "living-light-garden",
             "operation": "light_off"
         }
     ]
 }
 ```

## Recommended Tools
 - Restlet client ... https://chrome.google.com/webstore/detail/restlet-client-rest-api-t/aejoelaoggembcahagimdiliamlcdmfm
