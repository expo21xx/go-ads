# Index Groups

## Specification "Index-Group" of the PLC

| Index-Group (0x = hex) | Index Group description                                                                                                    |
|------------------------|----------------------------------------------------------------------------------------------------------------------------|
| 0x00000000 0x00000FFF  | reserved                                                                                                                   |
| 0x00001000             | PLC ADS parameter range                                                                                                    |
| 0x00002000             | PLC ADS status range                                                                                                       |
| 0x00003000             | PLC ADS unit function range                                                                                                |
| 0x00004000             | [PLC ADS services (includes services to access PLC memory range (%M field))](https://infosys.beckhoff.com/content/1033/tc3_ads_intro/9007199371984395.html)                                               |
| 0x00006000 0x0000EFFF  | reserved for PLC ADS extension                                                                                             |
| 0x0000F000 0x0000FFFF  | [general TwinCAT ADS system services (includes services to access PLC process diagram of the physical inputs and outputs)](https://infosys.beckhoff.com/content/1033/tc3_ads_intro/18014398626945547.html) |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/117241867.html&id=1944752650545554679)



## Specification of the PLC services

| Index Group | Index Offset           | Access | Data Type  | Description                                                                                                                                                                                        |
|-------------|------------------------|--------|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0x00004020  | 0x00000000- 0x0000FFFF | R/W    | `UINT8[n]` | **READ_M - WRITE_M** PLC memory range(%M field).Offset is byte offset.                                                                                                                             |
| 0x00004021  | 0x00000000- 0xFFFFFFFF | R/W    | `UINT8`    | **READ_MX - WRITE_MX** PLC memory range (%MX field).The low word of the index offset is the byte offset. The index offset contains the bit address calculated from the byte number *8 + bit number |
| 0x00004025  | 0x00000000             | R      | `ULONG`    | **PLCADS_IGR_RMSIZE** Byte length of the process diagram of the memory range                                                                                                                       |
| 0x00004030  | 0x00000000- 0xFFFFFFFF | R/W    | `UINT8`    | **PLCADS_IGR_RWRB** Retain data range. The index offset is byte offset                                                                                                                             |
| 0x00004035  | 0x00000000             | R      | `ULONG`    | **PLCADS_IGR_RRSIZE** Byte length of the retain range                                                                                                                                              |
| 0x00004040  | 0x00000000- 0xFFFFFFFF | R/W    | `UINT8`    | **PLCADS_IGR_RWDB** Data range. The index offset is byte offset.                                                                                                                                   |
| 0x00004045  | 0x00000000             | R      | `ULONG`    | **PLCADS_IGR_RDSIZE** Byte length of the data range                                                                                                                                                |


(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/9007199371984395.html&id=5492819642433960109)

## Specification of the ADS system services

| Index Group | Index Offset                                                            | Access | Data Type                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Description                                                                                                                                                                                                                                          |
|-------------|-------------------------------------------------------------------------|--------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0x0000F003  | 0x00000000                                                              | R/W    | W: `UINT8[N]` R: `UINT32`                                                                                                                                                                                                                                                                                                                                                                                                                                    | **GET_SYMHANDLE_BYNAME** A handle (code word) is assigned to the name contained in the write data and is returned to the caller as a result.                                                                                                         |
| 0x0000F004  | 0x00000000                                                              |        |                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Reserved.                                                                                                                                                                                                                                            |
| 0x0000F005  | 0x00000000- 0xFFFFFFFF=symHandle                                        |        | `UINT8[N]`                                                                                                                                                                                                                                                                                                                                                                                                                                                   | **READ_/WRITE_SYMVAL_BYHANDLE** Reads the value of the variable identified by ‚symHdl' or assigns a value to the variable. The ‚symHdl' must first have been determined by the GET_SYMHANDLE_BYNAME services.                                        |
| 0x0000F006  | 0x00000000                                                              | W      | `UINT32`                                                                                                                                                                                                                                                                                                                                                                                                                                                     | **RELEASE_SYMHANDLE** The code (handle) contained in the write data for an interrogated, named PLC variable is released.                                                                                                                             |
| 0x0000F020  | 0x0001F400- 0xFFFFFFFF                                                  | R/W    | `UINT8[N]`                                                                                                                                                                                                                                                                                                                                                                                                                                                   | **READ_I - WRITE_I** PLC process diagram of the physical inputs (%I field). Offset is byte offset.                                                                                                                                                   |
| 0x0000F021  | 0x000FA000- 0xFFFFFFFF                                                  | R/W    | `UINT8`                                                                                                                                                                                                                                                                                                                                                                                                                                                      | **READ_IX - WRITE_IX** PLC process diagram of the physical inputs (%IX field). The index offset contains the bit address which is calculated from base offset (0xFA000) + byte number +8 + bit number                                                |
| 0x0000F025  | 0x00000000                                                              | R      | `ULONG`                                                                                                                                                                                                                                                                                                                                                                                                                                                      | **ADSIGRP_IOIMAGE_RISIZE** Byte length of the PLC process diagram of the physical inputs.                                                                                                                                                            |
| 0x0000F030  | 0x0003E800- 0xFFFFFFFF                                                  | R/W    | `UINT8[N]`                                                                                                                                                                                                                                                                                                                                                                                                                                                   | **READ_Q - WRITE_Q** PLC process diagram of the physical outputs (%Q field). Offset is byte offset.                                                                                                                                                  |
| 0x0000F031  | 0x001F4000- 0xFFFFFFFF                                                  | R/W    | `UINT8`                                                                                                                                                                                                                                                                                                                                                                                                                                                      | **READ_QX - WRITE_QX** PLC process diagram of the physical outputs(%QX field). The index offset contains the bit address which is calculated from the base offset (0x1F4000) + byte number *8 + bit number.                                          |
| 0x0000F035  | 0x00000000                                                              | R      | `ULONG`                                                                                                                                                                                                                                                                                                                                                                                                                                                      | **ADSIGRP_IOIMAGE_ROSIZE** Byte length of the PLC process diagram of the physical outputs.                                                                                                                                                           |
| 0x0000F080  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `n * ULONG[3] :=IG1, IO1, Len1,IG2, IO2, Len2,...,IG(n), IO(n), Len(n)` R: `n * ULONG+ UINT8[Len1]+ UINT8[Len2]+ ...,+ UINT8[Len(n)] := Result1, Result2, ..., Result(n),Data1, Data2, ..., Data(n)`                                                                                                                                                                                                                                                      | **ADSIGRP_SUMUP_READ** The write-data contains a list of multiple, separate AdsReadReq(IG, IO, Len, Data) sub-commands. The read-data contains a list of return codes followed by the requested data.                                                |
| 0x0000F081  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `(n * ULONG[3]) + UINT8[Len1] + UINT8[Len2] + ..., + UINT8[Len(n)] := IG1, IO1, Len1, IG2, IO2, Len2, ..., IG(n), IO(n), Len(n), Data1, Data2, ..., Data(n)`  R: `n * ULONG := Result1, Result2, ..., Result(n)`                                                                                                                                                                                                                                          | **ADSIGRP_SUMUP_WRITE** The write-data contains a list of multiple, separate AdsWriteReq(IG, IO, Len, Data) sub-commands. The read-data contains a list of return codes.                                                                             |
| 0x0000F082  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `(n * ULONG[4]) + UINT8[WriteLen1] + UINT8[WriteLen2] + ..., + UINT8[WriteLen(n)] := IG1, IO1, ReadLen1, WriteLen1, IG2, IO2, ReadLen2, WriteLen2, ..., IG(n), IO(n), ReadLen(n), ..., WriteLen(n), WriteData1, WriteData2, ..., WriteData(n)`  R: `(n * ULONG[2])+ UINT8[ReturnLen1] + UINT8[ReturnLen2] + ..., + UINT8[ReturnLen(n)] := Result1, ReturnLen1, Result2, ReturnLen2, ..., Result(n), ReturnLen(n), ReadData1, ReadData2, ..., ReadData(n)` | **ADSIGRP_SUMUP_READWRITE** The write-data contains a list of multiple, separate AdsReadWriteReq(IG, IO, readLen, writeLen, Data) sub-commands. The read-data contains a list of return codes and return data length followed by the requested data. |
| 0x0000F083  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `n * ULONG[3] := IG1, IO1, Len1, IG2, IO2, Len2, ..., IG(n), IO(n), Len(n)`  R: `n * ULONG + UINT8[Len1] + UINT8[Len2] + ..., + UINT8[Len(n)] := Result1, Result2, ..., Result(n), Data1, Data2, ..., Data(n)`                                                                                                                                                                                                                                            | **ADSIGRP_SUMUP_READEX** The write-data contains a list of multiple, separate AdsReadReq(IG, IO, Len, Data) sub-commands.The read-data contains a list of return codes followed by the requested data.                                               |
| 0x0000F084  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `n * ULONG[3] := IG1, IO1, Len1, IG2, IO2, Len2, ..., IG(n), IO(n), Len(n)`  R: `n * ULONG + UINT8[Len1] + UINT8[Len2] + ..., + UINT8[Len(n)] := Result1, Result2, ..., Result(n), Data1, Data2, ..., Data(n)`                                                                                                                                                                                                                                            | **ADSIGRP_SUMUP_READEX2** The write-data contains a list of multiple, separate AdsReadReq(IG, IO, Len, Data) sub-commands.The read-data contains a list of return codes followed by the requested data.                                              |
| 0x0000F085  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `(n * ULONG[3]) := IG1, IO1, Len1, IG2, IO2, Len2, ..., IG(n), IO(n), Len(n)`  R: `(n * ULONG) + UINT8[Len1] + UINT8[Len2] + ..., + UINT8[Len(n)] := Result1, Result2, ..., Result(n), Handle1, Handle2,..., Handle(n)`                                                                                                                                                                                                                                   | **ADSIGRP_SUMUP_ADDDEVNOTE** The write-data contains a list of multiple, separate AdsAddDeviceNotifications(IG, IO, Len, Data) sub-commands.The read-data contains a list of return codes followed by the requested notification handles.            |
| 0x0000F086  | 0x00000000- 0xFFFFFFFF= n (number of internal sub-commands)n(max) = 500 | R/W    | W: `Handle1, Handle2,..., Handle(n)`  R: `(n * ULONG) + UINT8[Len1] + UINT8[Len2] + ..., + UINT8[Len(n)] := Result1, Result2, ..., Result(n)`                                                                                                                                                                                                                                                                                                                | **ADSIGRP_SUMUP_DELDEVNOTE** The write-data contains a list of multiple handles.The read-data contains a list of return codes.                                                                                                                       |

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/18014398626945547.html&id=8490003811267189798)


## Further

Further optional and not implemented index groups can be found [here](https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/171136786553140747.html&id=5388918862860483324).

## Notice

The information on this page is for reference only. All copyright and trademarks are with their respective holders.
This information is not covered by the repositories license. Please see [https://infosys.beckhoff.com](https://infosys.beckhoff.com) for
more information.