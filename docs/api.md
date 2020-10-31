# REST API

## Registration

### Request token for user

```bash
$ curl http://${SERVER_ADDR}/auth/user/token \
-X POST -d '{"user_email":"user@mail"}'

{
  "token": "dXNlckBtYWls"
}
```

### Request token for device

```bash
$ curl http://${SERVER_ADDR}/auth/device/token \
-X POST -d '{"device_name":"Device-1", "user_email":"user@mail"}'

{
  "token": "Yjc5ZjhhMGMtY2Y2Mi00ZWMwLWExYWEtM2MzMzI0OTE5MTYw"
}
```

## Get user devices

Get all devices, associated with current authenticated user.

```bash
$ curl "http://${SERVER_ADDR}/devices/?token=${USER_TOKEN}"

{
  "items": [
    {
      "device": {
        "uuid": "08fc2dca-931c-4177-b87d-0f2f4da70051",
        "email": "2@email",
        "updatedAt": "2020-10-30T10:01:06.638Z"
      }
    }
  ],
  "total": 1,
  "next": null
}
```

## Get device info

Get device details.

```bash
$ curl "http://${SERVER_ADDR}/devices/byId?token=${USER_TOKEN}&id=${DEVICE_UUID}"

{
  "device": {
    "uuid": "08fc2dca-931c-4177-b87d-0f2f4da70051",
    "email": "2@email",
    "updatedAt": "2020-10-30T10:01:06.638Z"
  }
}
```

Get device temperature log.

```bash
$ curl "http://${SERVER_ADDR}/devices/byId/log?token=${USER_TOKEN}&id=${DEVICE_UUID}"

{
  "items": [
    {
      "ts": 1604311266498,
      "time": "2020-10-30T10:01:06.498Z",
      "temperature": 12.215
    },
    {
      "ts": 1604311266562,
      "time": "2020-10-30T10:01:06.638Z",
      "temperature": 27.991
    }
  ],
  "total": 2,
  "next": null
}
```

## Save temperature

Save temperature measurements for an authorized device.

```bash
$ curl http://${SERVER_ADDR}/measurements/temp?token=${DEVICE_TOKEN}" \
-X POST -d '[{"time":1604234663141, "value":12.34}]'
```
