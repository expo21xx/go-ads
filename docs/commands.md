# Commands


## Invalid Command (`0x0000`)

Go: `CommandInvalid`  

## ReadDeviceInfo (`0x0001`)

Go: `CommandADSReadDeviceInfo`  

Reads the name and the version number of the ADS device.


### Request

No data required

### Response

```
+---+---+---+---+---------------+---------------+-------+-------+
| 0 | 1 | 2 | 3 |       4       |       5       |   6   |   7   |
+---+---+---+---+---------------+---------------+-------+-------+
|     Result    | Major Version | Minor Version | Version Build |
+---------------+---------------+---------------+---------------+
|                          Device Name                          |
+---------------------------------------------------------------+
|                                                               |
+---------------------------------------------------------------+
```

| Data array    | Size     | Description          |
|---------------|----------|----------------------|
| Result        | 4 bytes  | ADS error number.    |
| Major Version | 1 byte   | Major version number |
| Minor Version | 1 byte   | Minor version number |
| Version Build | 2 bytes  | Build number         |
| Device Name   | 16 bytes | Name of ADS device   |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115875851.html&id=8158832529229503828)


## Read (`0x0002`)

Go: `CommandADSRead`  

With ADS Read data can be read from an ADS device. The data are addressed by the Index Group and the Index Offset

### Request


```
+---+---+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+---+
|  Index Group  |  Index Offset |
+---------------+---+---+---+---+
|     Length    |   |   |   |   |
+---------------+---+---+---+---+
```

| Data array   | Size    | Description                                         |
|--------------|---------|-----------------------------------------------------|
| Index Group  | 4 bytes | Index Group of the data which should be read.       |
| Index Offset | 4 bytes | Index Offset of the data which should be read.      |
| Length       | 4 bytes | Length of the data (in bytes) which should be read. |

### Response

```
+---+---+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+---+
|     Result    |     Length    |
+---------------+---------------+
|              Data             |
+-------------------------------+
```

| Data array | Size    | Description                             |
|------------|---------|-----------------------------------------|
| Result     | 4 bytes | ADS error number                        |
| Length     | 4 bytes | Length of data which are supplied back. |
| Data       | n bytes | Data which are supplied back.           |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115876875.html&id=4960931295000833536)


## Write (`0x0003`)

Go: `CommandADSWrite`  

With ADS Write data can be written to an ADS device. The data are addressed by the Index Group and the Index Offset.

### Request

```
+---+---+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+---+
|  Index Group  |  Index Offset |
+---------------+---------------+
|     Length    |    Data...    |
+---------------+---------------+
|              ...              |
+-------------------------------+
```

| Data array   | Size    | Description                                       |
|--------------|---------|---------------------------------------------------|
| Index Group  | 4 bytes | Index Group of the data which should be written.  |
| Index Offset | 4 bytes | Index Offset of the data which should be written. |
| Length       | 4 bytes | Length of data in bytes which are written         |
| Data         | n bytes | Data which are written in the ADS device.         |


### Response

```
+---+---+---+---+
| 0 | 1 | 2 | 3 |
+---+---+---+---+
|     Result    |
+---------------+
```


| Data array | Size    | Description      |
|------------|---------|------------------|
| Result     | 4 bytes | ADS error number |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115877899.html&id=8845698684103663373)

## ReadState (`0x0004`)

Go: `CommandADSReadState`  

Reads the ADS status and the device status of an ADS device.


### Request

No data required.


### Response

```
+---+---+---+---+-----+-----+-------+------+
| 0 | 1 | 2 | 3 |  4  |  5  |   6   |   7  |
+---+---+---+---+-----+-----+-------+------+
|     Result    | ADS State | Device State |
+---------------+-----------+--------------+
```

| Data array   | Size    | Description                                     |
|--------------|---------|-------------------------------------------------|
| Result       | 4 bytes | Index Group of the data which should be written.|
| ADS State    | 2 bytes | ADS status (see ads_state.go).                  |
| Device State | 2 bytes | Device status.                                  |


(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115878923.html&id=6874981934243835072)


## WriteControl (`0x0005`)

Go: `CommandADSWriteControl`  


Changes the ADS status and the device status of an ADS device. Additionally it is possible to send data to the ADS device to transfer
further information. These data were not analyzed from the current ADS devices (PLC, NC, ...).

### Request

```
+-----+-----+-------+------+---+---+---+---+
|  0  |  1  |   2   |   3  | 4 | 5 | 6 | 7 |
+-----+-----+-------+------+---+---+---+---+
| ADS State | Device State |     Length    |
+-----------+--------------+---------------+
|                   Data                   |
+------------------------------------------+
```

| Data array   | Size    | Description                                       |
|--------------|---------|---------------------------------------------------|
| ADS Statea   | 2 bytes | New ADS status (see ads_state.go).                |
| Device State | 2 bytes | New device status.                                |
| Length       | 4 bytes | Length of data in byte.                           |
| Data         | n bytes | Additional data which are sent to the ADS device. |


### Response

```
+---+---+---+---+
| 0 | 1 | 2 | 3 |
+---+---+---+---+
|     Result    |
+---------------+
```

| Data array | Size    | Description      |
|------------|---------|------------------|
| Result     | 4 bytes | ADS error number |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115879947.html&id=4720330147059483431)


## AddDeviceNotification (`0x0006`)

Go: `CommandADSAddDeviceNotification`  

A notification is created in an ADS device.

Note: We recommend to announce not more than 550 notifications per device.
Otherwise increase the payload by working with structures or use sum commands.



### Request

```
+---+---+---+---+----+----+----+----+
| 0 | 1 | 2 | 3 |  4 |  5 |  6 |  7 |
+---+---+---+---+----+----+----+----+
|  Index Group  |    Index Offset   |
+---------------+-------------------+
|     Length    | Transmission Mode |
+---------------+-------------------+
|   Max Delay   |     Cycle Time    |
+---------------+-------------------+
|                                   |
+            reserved               +
|                                   |
+-----------------------------------+
```

| Data array        | Size     | Description                                                                            |
|-------------------|----------|----------------------------------------------------------------------------------------|
| Index Group       | 4 bytes  | Index Group of the data, which should be sent per notification.                        |
| Index Offset      | 4 bytes  | Index Offset of the data, which should be sent per notification.                       |
| Length            | 4 bytes  | Length of data in bytes, which should be sent per notification.                        |
| Transmission Mode | 4 bytes  | See `ads_transmission_mode.go`                                                         |
| Max Delay         | 4 bytes  | At the latest after this time, the ADS Device Notification is called. The unit is 1ms. |
| Cycle Time        | 4 bytes  | The ADS server checks if the value changes in this time slice. The unit is 1ms         |
| reserved          | 16 bytes | Must be set to 0                                                                       |

### Response

```
+---+---+---+---+----+-----+-----+----+
| 0 | 1 | 2 | 3 | 4  |  5  |  6  |  7 |
+---+---+---+---+----+-----+-----+----+
|     Result    | Notification Handle |
```

| Data array          | Size    | Description            |
|---------------------|---------|------------------------|
| Result              | 4 bytes | ADS error number       |
| Notification Handle | 4 bytes | Handle of notification |


(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115880971.html&id=7388557527878561663)


## DeletDeviceNotification (`0x0007`)

Go: `CommandADSDeletDeviceNotification`  

One before defined notification is deleted in an ADS device.

### Request


```
+----+-----+-----+----+
| 0  |  1  |  2  |  3 |
+----+-----+-----+----+
| Notification Handle |
```

| Data array          | Size    | Description            |
|---------------------|---------|------------------------|
| Notification Handle | 4 bytes | Handle of notification |


### Response


```
+---+---+---+---+
| 0 | 1 | 2 | 3 |
+---+---+---+---+
|     Result    |
+---------------+
```

| Data array | Size    | Description      |
|------------|---------|------------------|
| Result     | 4 bytes | ADS error number |


(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115881995.html&id=6216061301016726131)


## DeviceNotification (`0x0008`)

Go: `CommandADSDeviceNotification`  

Data will carry forward independently from an ADS device to a Client.


### Request

```
+---+---+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+---+
|     Length    |     Stamps    |
+---------------+---------------+
|       AdsStampHeader[0]       |
|              ...              |
+                               +
|    AdsStampHeader[Stamps-1]   |
+-------------------------------+
```

| Data array     | Size    | Description                                |
|----------------|---------|--------------------------------------------|
| Length         | 4 bytes | Size of data in byte.                      |
| Stamps         | 4 bytes | Number of elements of type AdsStampHeader  |
| AdsStampHeader | n bytes | Array with elements of type AdsStampHeader |


#### AdsStampHeader 

```
+----+----+----+---+---+---+---+---+
|  0 |  1 |  2 | 3 | 4 | 5 | 6 | 7 |
+----+----+----+---+---+---+---+---+
|             TimeStamp            |
+------------------+---------------+
|      Samples     |               |
+------------------+               +
|     AdsNotificationSample[0]     |
|                ...               |
| AdsNotificationSample[Samples-1] |
+----------------------------------+
```

| Data array            | Size    | Description                                                                                                                                                                                                                                                             |
|-----------------------|---------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| TimeStamp             | 8 bytes | The timestamp is coded after the Windos FILETIME format. I.e. the value contains the number of the nano seconds, which passed since 1.1.1601. In addition, the local time change is not considered. Thus the time stamp is present as universal Coordinated time (UTC). |
| Samples               | 4 bytes | Number of elements of type AdsNotificationSample                                                                                                                                                                                                                        |
| AdsNotificationSample | n bytes | Array with elements of type AdsNotificationSample                                                                                                                                                                                                                       |


#### AdsNotificationSample

```
+-----+-----+----+----+---+---+---+---+
|  0  |  1  |  2 |  3 | 4 | 5 | 6 | 7 |
+-----+-----+----+----+---+---+---+---+
| Notification Handle |  Sample Size  |
+---------------------+---------------+
|                 Data                |
|                 ...                 |
+-------------------------------------+
```

| Data array          | Size    | Description                  |
|---------------------|---------|------------------------------|
| Notification Handle | 4 bytes | Handle of notification.      |
| Sample Size         | 4 bytes | Size of data range in bytes. |
| Data                | n bytes | Data                         |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/9007199370624011.html&id=3423629729326333060)


## ReadWrite (`0x0009`)

Go: `CommandADSReadWrite`  

With ADS ReadWrite data will be written to an ADS device. Additionally, data can be read from the ADS device.
The data which can be read are addressed by the Index Group and the Index Offset

### Request

```
+---+---+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+---+
|  Index Group  |  Index Offset |
+---------------+---------------+
| Read Length   | Write Length  |
+---------------+---------------+
|              Data             |
|              ...              |
+-------------------------------+
```

| Data array   | Size    | Description                                       |
|--------------|---------|---------------------------------------------------|
| Index Group  | 4 bytes | Index Group, in which the data should be written. |
| Index Offset | 4 bytes | ndex Offset, in which the data should be written  |
| Read Length  | 4 bytes | Length of data in bytes, which should be read.    |
| Write Length | 4 bytes | Length of data in bytes, which should be written. |
| Data         | n bytes | Data which are written in the ADS device.         |


### Response

```
+---+---+---+---+---+---+---+---+
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
+---+---+---+---+---+---+---+---+
|     Result    |     Length    |
+---------------+---------------+
|              Data             |
|              ...              |
+-------------------------------+
```


| Data array | Size    | Description                             |
|------------|---------|-----------------------------------------|
| Result     | 4 bytes | ADS error number                        |
| Length     | 4 bytes | Length of data which are supplied back. |
| Data       | n bytes | Data which are supplied back.           |


## Notice

The information on this page is for reference only. All copyright and trademarks are with their respective holders.
This information is not covered by the repositories license. Please see [https://infosys.beckhoff.com](https://infosys.beckhoff.com) for
more information.
