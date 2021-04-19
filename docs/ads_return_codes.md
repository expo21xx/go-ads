# ADS Return Codes

(Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tf6701_tc3_iot_communication_mqtt/99079192176427915.html&id=)

## Global error codes

| Hex  | Dec | HRESULT     | Name                      | Description                                                                                                |
|------|-----|-------------|---------------------------|------------------------------------------------------------------------------------------------------------|
| 0x0  | 0   | 0x9811 0000 | ERR_NOERROR               | No error.                                                                                                  |
| 0x1  | 1   | 0x9811 0001 | ERR_INTERNAL              | Internal error.                                                                                            |
| 0x2  | 2   | 0x9811 0002 | ERR_NORTIME               | No real-time.                                                                                              |
| 0x3  | 3   | 0x9811 0003 | ERR_ALLOCLOCKEDMEM        | Allocation locked – memory error.                                                                          |
| 0x4  | 4   | 0x9811 0004 | ERR_INSERTMAILBOX         | Mailbox full – the ADS message could not be sent. Reducing the number of ADS messages per cycle will help. |
| 0x5  | 5   | 0x9811 0005 | ERR_WRONGRECEIVEHMSG      | Wrong HMSG.                                                                                                |
| 0x6  | 6   | 0x9811 0006 | ERR_TARGETPORTNOTFOUND    | Target port not found – ADS server is not started or is not reachable.                                     |
| 0x7  | 7   | 0x9811 0007 | ERR_TARGETMACHINENOTFOUND | Target computer not found – AMS route was not found.                                                       |
| 0x8  | 8   | 0x9811 0008 | ERR_UNKNOWNCMDID          | Unknown command ID.                                                                                        |
| 0x9  | 9   | 0x9811 0009 | ERR_BADTASKID             | Invalid task ID.                                                                                           |
| 0xA  | 10  | 0x9811 000A | ERR_NOIO                  | No IO.                                                                                                     |
| 0xB  | 11  | 0x9811 000B | ERR_UNKNOWNAMSCMD         | Unknown AMS command.                                                                                       |
| 0xC  | 12  | 0x9811 000C | ERR_WIN32ERROR            | Win32 error.                                                                                               |
| 0xD  | 13  | 0x9811 000D | ERR_PORTNOTCONNECTED      | Port not connected.                                                                                        |
| 0xE  | 14  | 0x9811 000E | ERR_INVALIDAMSLENGTH      | Invalid AMS length.                                                                                        |
| 0xF  | 15  | 0x9811 000F | ERR_INVALIDAMSNETID       | Invalid AMS Net ID.                                                                                        |
| 0x10 | 16  | 0x9811 0010 | ERR_LOWINSTLEVEL          | Installation level is too low –TwinCAT 2 license error.                                                    |
| 0x11 | 17  | 0x9811 0011 | ERR_NODEBUGINTAVAILABLE   | No debugging available.                                                                                    |
| 0x12 | 18  | 0x9811 0012 | ERR_PORTDISABLED          | Port disabled – TwinCAT system service not started.                                                        |
| 0x13 | 19  | 0x9811 0013 | ERR_PORTALREADYCONNECTED  | Port already connected.                                                                                    |
| 0x14 | 20  | 0x9811 0014 | ERR_AMSSYNC_W32ERROR      | AMS Sync Win32 error.                                                                                      |
| 0x15 | 21  | 0x9811 0015 | ERR_AMSSYNC_TIMEOUT       | AMS Sync Timeout.                                                                                          |
| 0x16 | 22  | 0x9811 0016 | ERR_AMSSYNC_AMSERROR      | AMS Sync error.                                                                                            |
| 0x17 | 23  | 0x9811 0017 | ERR_AMSSYNC_NOINDEXINMAP  | No index map for AMS Sync available.                                                                       |
| 0x18 | 24  | 0x9811 0018 | ERR_INVALIDAMSPORT        | Invalid AMS port.                                                                                          |
| 0x19 | 25  | 0x9811 0019 | ERR_NOMEMORY              | No memory.                                                                                                 |
| 0x1A | 26  | 0x9811 001A | ERR_TCPSEND               | TCP send error.                                                                                            |
| 0x1B | 27  | 0x9811 001B | ERR_HOSTUNREACHABLE       | Host unreachable.                                                                                          |
| 0x1C | 28  | 0x9811 001C | ERR_INVALIDAMSFRAGMENT    | Invalid AMS fragment.                                                                                      |
| 0x1D | 29  | 0x9811 001D | ERR_TLSSEND               | TLS send error – secure ADS connection failed.                                                             |
| 0x1E | 30  | 0x9811 001E | ERR_ACCESSDENIED          | Access denied – secure ADS access denied.                                                                  |

## Router error codes

| Hex   | Dec  | HRESULT     | Name                       | Description                                                            |
|-------|------|-------------|----------------------------|------------------------------------------------------------------------|
| 0x500 | 1280 | 0x9811 0500 | ROUTERERR_NOLOCKEDMEMORY   | Locked memory cannot be allocated.                                     |
| 0x501 | 1281 | 0x9811 0501 | ROUTERERR_RESIZEMEMORY     | The router memory size could not be changed.                           |
| 0x502 | 1282 | 0x9811 0502 | ROUTERERR_MAILBOXFULL      | The mailbox has reached the maximum number of possible messages.       |
| 0x503 | 1283 | 0x9811 0503 | ROUTERERR_DEBUGBOXFULL     | The Debug mailbox has reached the maximum number of possible messages. |
| 0x504 | 1284 | 0x9811 0504 | ROUTERERR_UNKNOWNPORTTYPE  | The port type is unknown.                                              |
| 0x505 | 1285 | 0x9811 0505 | ROUTERERR_NOTINITIALIZED   | The router is not initialized.                                         |
| 0x506 | 1286 | 0x9811 0506 | ROUTERERR_PORTALREADYINUSE | The port number is already assigned.                                   |
| 0x507 | 1287 | 0x9811 0507 | ROUTERERR_NOTREGISTERED    | The port is not registered.                                            |
| 0x508 | 1288 | 0x9811 0508 | ROUTERERR_NOMOREQUEUES     | The maximum number of ports has been reached.                          |
| 0x509 | 1289 | 0x9811 0509 | ROUTERERR_INVALIDPORT      | The port is invalid.                                                   |
| 0x50A | 1290 | 0x9811 050A | ROUTERERR_NOTACTIVATED     | The router is not active.                                              |
| 0x50B | 1291 | 0x9811 050B | ROUTERERR_FRAGMENTBOXFULL  | The mailbox has reached the maximum number for fragmented messages.    |
| 0x50C | 1292 | 0x9811 050C | ROUTERERR_FRAGMENTTIMEOUT  | A fragment timeout has occurred.                                       |
| 0x50D | 1293 | 0x9811 050D | ROUTERERR_TOBEREMOVED      | The port is removed.                                                   |


## General ADS error codes


| Hex   | Dec  | HRESULT     | Name                               | Description                                                                                                                                                                               |
|-------|------|-------------|------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0x700 | 1792 | 0x9811 0700 | ADSERR_DEVICE_ERROR                | General device error.                                                                                                                                                                     |
| 0x701 | 1793 | 0x9811 0701 | ADSERR_DEVICE_SRVNOTSUPP           | Service is not supported by the server.                                                                                                                                                   |
| 0x702 | 1794 | 0x9811 0702 | ADSERR_DEVICE_INVALIDGRP           | Invalid index group.                                                                                                                                                                      |
| 0x703 | 1795 | 0x9811 0703 | ADSERR_DEVICE_INVALIDOFFSET        | Invalid index offset.                                                                                                                                                                     |
| 0x704 | 1796 | 0x9811 0704 | ADSERR_DEVICE_INVALIDACCESS        | Reading or writing not permitted.                                                                                                                                                         |
| 0x705 | 1797 | 0x9811 0705 | ADSERR_DEVICE_INVALIDSIZE          | Parameter size not correct.                                                                                                                                                               |
| 0x706 | 1798 | 0x9811 0706 | ADSERR_DEVICE_INVALIDDATA          | Invalid data values.                                                                                                                                                                      |
| 0x707 | 1799 | 0x9811 0707 | ADSERR_DEVICE_NOTREADY             | Device is not ready to operate.                                                                                                                                                           |
| 0x708 | 1800 | 0x9811 0708 | ADSERR_DEVICE_BUSY                 | Device is busy.                                                                                                                                                                           |
| 0x709 | 1801 | 0x9811 0709 | ADSERR_DEVICE_INVALIDCONTEXT       | Invalid operating system context. This can result from use of ADS function blocks in different tasks. It may be possible to resolve this through multitasking synchronization in the PLC. |
| 0x70A | 1802 | 0x9811 070A | ADSERR_DEVICE_NOMEMORY             | Insufficient memory.                                                                                                                                                                      |
| 0x70B | 1803 | 0x9811 070B | ADSERR_DEVICE_INVALIDPARM          | Invalid parameter values.                                                                                                                                                                 |
| 0x70C | 1804 | 0x9811 070C | ADSERR_DEVICE_NOTFOUND             | Not found (files, ...).                                                                                                                                                                   |
| 0x70D | 1805 | 0x9811 070D | ADSERR_DEVICE_SYNTAX               | Syntax error in file or command.                                                                                                                                                          |
| 0x70E | 1806 | 0x9811 070E | ADSERR_DEVICE_INCOMPATIBLE         | Objects do not match.                                                                                                                                                                     |
| 0x70F | 1807 | 0x9811 070F | ADSERR_DEVICE_EXISTS               | Object already exists.                                                                                                                                                                    |
| 0x710 | 1808 | 0x9811 0710 | ADSERR_DEVICE_SYMBOLNOTFOUND       | Symbol not found.                                                                                                                                                                         |
| 0x711 | 1809 | 0x9811 0711 | ADSERR_DEVICE_SYMBOLVERSIONINVALID | Invalid symbol version. This can occur due to an online change. Create a new handle.                                                                                                      |
| 0x712 | 1810 | 0x9811 0712 | ADSERR_DEVICE_INVALIDSTATE         | Device (server) is in invalid state.                                                                                                                                                      |
| 0x713 | 1811 | 0x9811 0713 | ADSERR_DEVICE_TRANSMODENOTSUPP     | AdsTransMode not supported.                                                                                                                                                               |
| 0x714 | 1812 | 0x9811 0714 | ADSERR_DEVICE_NOTIFYHNDINVALID     | Notification handle is invalid.                                                                                                                                                           |
| 0x715 | 1813 | 0x9811 0715 | ADSERR_DEVICE_CLIENTUNKNOWN        | Notification client not registered.                                                                                                                                                       |
| 0x716 | 1814 | 0x9811 0716 | ADSERR_DEVICE_NOMOREHDLS           | No further notification handle available.                                                                                                                                                 |
| 0x717 | 1815 | 0x9811 0717 | ADSERR_DEVICE_INVALIDWATCHSIZE     | Notification size too large.                                                                                                                                                              |
| 0x718 | 1816 | 0x9811 0718 | ADSERR_DEVICE_NOTINIT              | Device not initialized.                                                                                                                                                                   |
| 0x719 | 1817 | 0x9811 0719 | ADSERR_DEVICE_TIMEOUT              | Device has a timeout.                                                                                                                                                                     |
| 0x71A | 1818 | 0x9811 071A | ADSERR_DEVICE_NOINTERFACE          | Interface query failed.                                                                                                                                                                   |
| 0x71B | 1819 | 0x9811 071B | ADSERR_DEVICE_INVALIDINTERFACE     | Wrong interface requested.                                                                                                                                                                |
| 0x71C | 1820 | 0x9811 071C | ADSERR_DEVICE_INVALIDCLSID         | Class ID is invalid.                                                                                                                                                                      |
| 0x71D | 1821 | 0x9811 071D | ADSERR_DEVICE_INVALIDOBJID         | Object ID is invalid.                                                                                                                                                                     |
| 0x71E | 1822 | 0x9811 071E | ADSERR_DEVICE_PENDING              | Request pending.                                                                                                                                                                          |
| 0x71F | 1823 | 0x9811 071F | ADSERR_DEVICE_ABORTED              | Request is aborted.                                                                                                                                                                       |
| 0x720 | 1824 | 0x9811 0720 | ADSERR_DEVICE_WARNING              | Signal warning.                                                                                                                                                                           |
| 0x721 | 1825 | 0x9811 0721 | ADSERR_DEVICE_INVALIDARRAYIDX      | Invalid array index.                                                                                                                                                                      |
| 0x722 | 1826 | 0x9811 0722 | ADSERR_DEVICE_SYMBOLNOTACTIVE      | Symbol not active.                                                                                                                                                                        |
| 0x723 | 1827 | 0x9811 0723 | ADSERR_DEVICE_ACCESSDENIED         | Access denied.                                                                                                                                                                            |
| 0x724 | 1828 | 0x9811 0724 | ADSERR_DEVICE_LICENSENOTFOUND      | Missing license.                                                                                                                                                                          |
| 0x725 | 1829 | 0x9811 0725 | ADSERR_DEVICE_LICENSEEXPIRED       | License expired.                                                                                                                                                                          |
| 0x726 | 1830 | 0x9811 0726 | ADSERR_DEVICE_LICENSEEXCEEDED      | License exceeded.                                                                                                                                                                         |
| 0x727 | 1831 | 0x9811 0727 | ADSERR_DEVICE_LICENSEINVALID       | Invalid license.                                                                                                                                                                          |
| 0x728 | 1832 | 0x9811 0728 | ADSERR_DEVICE_LICENSESYSTEMID      | License problem: System ID is invalid.                                                                                                                                                    |
| 0x729 | 1833 | 0x9811 0729 | ADSERR_DEVICE_LICENSENOTIMELIMIT   | License not limited in time.                                                                                                                                                              |
| 0x72A | 1834 | 0x9811 072A | ADSERR_DEVICE_LICENSEFUTUREISSUE   | License problem: Time in the future.                                                                                                                                                      |
| 0x72B | 1835 | 0x9811 072B | ADSERR_DEVICE_LICENSETIMETOLONG    | License period too long.                                                                                                                                                                  |
| 0x72C | 1836 | 0x9811 072C | ADSERR_DEVICE_EXCEPTION            | Exception at system startup.                                                                                                                                                              |
| 0x72D | 1837 | 0x9811 072D | ADSERR_DEVICE_LICENSEDUPLICATED    | License file read twice.                                                                                                                                                                  |
| 0x72E | 1838 | 0x9811 072E | ADSERR_DEVICE_SIGNATUREINVALID     | Invalid signature.                                                                                                                                                                        |
| 0x72F | 1839 | 0x9811 072F | ADSERR_DEVICE_CERTIFICATEINVALID   | Invalid certificate.                                                                                                                                                                      |
| 0x730 | 1840 | 0x9811 0730 | ADSERR_DEVICE_LICENSEOEMNOTFOUND   | Public key not known from OEM.                                                                                                                                                            |
| 0x731 | 1841 | 0x9811 0731 | ADSERR_DEVICE_LICENSERESTRICTED    | License not valid for this system ID.                                                                                                                                                     |
| 0x732 | 1842 | 0x9811 0732 | ADSERR_DEVICE_LICENSEDEMODENIED    | Demo license prohibited.                                                                                                                                                                  |
| 0x733 | 1843 | 0x9811 0733 | ADSERR_DEVICE_INVALIDFNCID         | Invalid function ID.                                                                                                                                                                      |
| 0x734 | 1844 | 0x9811 0734 | ADSERR_DEVICE_OUTOFRANGE           | Outside the valid range.                                                                                                                                                                  |
| 0x735 | 1845 | 0x9811 0735 | ADSERR_DEVICE_INVALIDALIGNMENT     | Invalid alignment.                                                                                                                                                                        |
| 0x736 | 1846 | 0x9811 0736 | ADSERR_DEVICE_LICENSEPLATFORM      | Invalid platform level.                                                                                                                                                                   |
| 0x737 | 1847 | 0x9811 0737 | ADSERR_DEVICE_FORWARD_PL           | Context – forward to passive level.                                                                                                                                                       |
| 0x738 | 1848 | 0x9811 0738 | ADSERR_DEVICE_FORWARD_DL           | Context – forward to dispatch level.                                                                                                                                                      |
| 0x739 | 1849 | 0x9811 0739 | ADSERR_DEVICE_FORWARD_RT           | Context – forward to real-time.                                                                                                                                                           |
| 0x740 | 1856 | 0x9811 0740 | ADSERR_CLIENT_ERROR                | Client error.                                                                                                                                                                             |
| 0x741 | 1857 | 0x9811 0741 | ADSERR_CLIENT_INVALIDPARM          | Service contains an invalid parameter.                                                                                                                                                    |
| 0x742 | 1858 | 0x9811 0742 | ADSERR_CLIENT_LISTEMPTY            | Polling list is empty.                                                                                                                                                                    |
| 0x743 | 1859 | 0x9811 0743 | ADSERR_CLIENT_VARUSED              | Var connection already in use.                                                                                                                                                            |
| 0x744 | 1860 | 0x9811 0744 | ADSERR_CLIENT_DUPLINVOKEID         | The called ID is already in use.                                                                                                                                                          |
| 0x745 | 1861 | 0x9811 0745 | ADSERR_CLIENT_SYNCTIMEOUT          | Timeout has occurred – the remote terminal is not responding in the specified ADS timeout. The route setting of the remote terminal may be configured incorrectly.                        |
| 0x746 | 1862 | 0x9811 0746 | ADSERR_CLIENT_W32ERROR             | Error in Win32 subsystem.                                                                                                                                                                 |
| 0x747 | 1863 | 0x9811 0747 | ADSERR_CLIENT_TIMEOUTINVALID       | Invalid client timeout value.                                                                                                                                                             |
| 0x748 | 1864 | 0x9811 0748 | ADSERR_CLIENT_PORTNOTOPEN          | Port not open.                                                                                                                                                                            |
| 0x749 | 1865 | 0x9811 0749 | ADSERR_CLIENT_NOAMSADDR            | No AMS address.                                                                                                                                                                           |
| 0x750 | 1872 | 0x9811 0750 | ADSERR_CLIENT_SYNCINTERNAL         | Internal error in Ads sync.                                                                                                                                                               |
| 0x751 | 1873 | 0x9811 0751 | ADSERR_CLIENT_ADDHASH              | Hash table overflow.                                                                                                                                                                      |
| 0x752 | 1874 | 0x9811 0752 | ADSERR_CLIENT_REMOVEHASH           | Key not found in the table.                                                                                                                                                               |
| 0x753 | 1875 | 0x9811 0753 | ADSERR_CLIENT_NOMORESYM            | No symbols in the cache.                                                                                                                                                                  |
| 0x754 | 1876 | 0x9811 0754 | ADSERR_CLIENT_SYNCRESINVALID       | Invalid response received.                                                                                                                                                                |
| 0x755 | 1877 | 0x9811 0755 | ADSERR_CLIENT_SYNCPORTLOCKED       | Sync Port is locked.                                                                                                                                                                      |



## RTime error codes


| Hex    | Dec  | HRESULT     | Name                      | Description                                                                               |
|--------|------|-------------|---------------------------|-------------------------------------------------------------------------------------------|
| 0x1000 | 4096 | 0x9811 1000 | RTERR_INTERNAL            | Internal error in the real-time system.                                                   |
| 0x1001 | 4097 | 0x9811 1001 | RTERR_BADTIMERPERIODS     | Timer value is not valid.                                                                 |
| 0x1002 | 4098 | 0x9811 1002 | RTERR_INVALIDTASKPTR      | Task pointer has the invalid value 0 (zero).                                              |
| 0x1003 | 4099 | 0x9811 1003 | RTERR_INVALIDSTACKPTR     | Stack pointer has the invalid value 0 (zero).                                             |
| 0x1004 | 4100 | 0x9811 1004 | RTERR_PRIOEXISTS          | The request task priority is already assigned.                                            |
| 0x1005 | 4101 | 0x9811 1005 | RTERR_NOMORETCB           | No free TCB (Task Control Block) available. The maximum number of TCBs is 64.             |
| 0x1006 | 4102 | 0x9811 1006 | RTERR_NOMORESEMAS         | No free semaphores available. The maximum number of semaphores is 64.                     |
| 0x1007 | 4103 | 0x9811 1007 | RTERR_NOMOREQUEUES        | No free space available in the queue. The maximum number of positions in the queue is 64. |
| 0x100D | 4109 | 0x9811 100D | RTERR_EXTIRQALREADYDEF    | An external synchronization interrupt is already applied.                                 |
| 0x100E | 4110 | 0x9811 100E | RTERR_EXTIRQNOTDEF        | No external sync interrupt applied.                                                       |
| 0x100F | 4111 | 0x9811 100F | RTERR_EXTIRQINSTALLFAILED | Application of the external synchronization interrupt has failed.                         |
| 0x1010 | 4112 | 0x9811 1010 | RTERR_IRQLNOTLESSOREQUAL  | Call of a service function in the wrong context                                           |
| 0x1017 | 4119 | 0x9811 1017 | RTERR_VMXNOTSUPPORTED     | Intel VT-x extension is not supported.                                                    |
| 0x1018 | 4120 | 0x9811 1018 | RTERR_VMXDISABLED         | Intel VT-x extension is not enabled in the BIOS.                                          |
| 0x1019 | 4121 | 0x9811 1019 | RTERR_VMXCONTROLSMISSING  | Missing function in Intel VT-x extension.                                                 |
| 0x101A | 4122 | 0x9811 101A | RTERR_VMXENABLEFAILS      | Activation of Intel VT-x fails.                                                           |


## TCP Winsock error codes


| Hex                                         | Dec   | Name            | Description                                                                                                                                                                                                                                                                              |
|---------------------------------------------|-------|-----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 0x274C                                      | 10060 | WSAETIMEDOUT    | A connection timeout has occurred - error while establishing the connection, because the remote terminal did not respond properly after a certain period of time, or the established connection could not be maintained because the connected host did not respond.                      |
| 0x274D                                      | 10061 | WSAECONNREFUSED | Connection refused - no connection could be established because the target computer has explicitly rejected it. This error usually results from an attempt to connect to a service that is inactive on the external host, that is, a service for which no server application is running. |
| 0x2751                                      | 10065 | WSAEHOSTUNREACH | No route to host - a socket operation referred to an unavailable host.                                                                                                                                                                                                                   |
| More Winsock error codes: Win32 error codes |


## Notice

The information on this page is for reference only. All copyright and trademarks are with their respective holders.
This information is not covered by the repositories license. Please see [https://infosys.beckhoff.com](https://infosys.beckhoff.com) for
more information.
