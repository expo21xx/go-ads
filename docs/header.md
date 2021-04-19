# AMS/ADS Headers

## AMS TCP Header

```
+-----+----+---+---+---+---+
| 0   | 1  | 2 | 3 | 4 | 5 |
+-----+----+---+---+---+---+
| reserved |     length    |
+----------+---------------+
```

Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115846283.html&id=5591912318145837195


## AMS Header

```
+------+-----+------+------+---+---+--------+-------+
| 0    | 1   | 2    | 3    | 4 | 5 | 6      | 7     |
+------+-----+------+------+---+---+--------+-------+
|          AMSNetId Target         | AMSPort Target |
+----------------------------------+----------------+
|          AMSNetId Source         | AMSPort Source |
+------------+-------------+-------+----------------+
| Command Id | State Flags |       Data Length      |
+------------+-------------+------------------------+
|        Error Code        |        Invoke Id       |
+--------------------------+------------------------+
|                        Data                       |
+---------------------------------------------------+
```

| Data array      | Size    | Description                                                                                                                                                     |
|-----------------|---------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| AMSNetId Target | 6 bytes | This is the AMSNetId of the station, for which the packet is intended. Remarks see below.                                                                       |
| AMSPort Target  | 2 Bytes | This is the AMSPort of the station, for which the packet is intended.                                                                                           |
| AMSNetId Source | 6 bytes | This contains the AMSPort of the station, from which the packet was sent.                                                                                       |
| AMSPort Source  | 2 Bytes | This contains the AMSPort of the station, from which the packet was sent.                                                                                       |
| Command Id      | 2 Bytes | See [docs/commadns.md].                                                                                                                                         |
| State Flags     | 2 Bytes | See below.                                                                                                                                                      |
| Data Length     | 4 bytes | Size of the data range. The unit is byte.                                                                                                                       |
| Error Code      | 4 bytes | AMS error number. See ADS Return Codes.                                                                                                                         |
| Invoke Id       | 4 bytes | Free usable 32 bit array. Usually this array serves to send an Id. This Id makes is possible to assign a received response to a request, which was sent before. |
| Data            | n bytes | Data range. The data range contains the parameter of the considering ADS commands.                                                                              |

Source: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/115847307.html&id=7738940192708835096

### AMS NetID

The AMSNetId consists of 6 bytes and addresses the transmitter or receiver. One possible AMSNetId would be e.g. `172.16.17.10.1.1.`.
The storage arrangement in this example is as follows:

|  0  |  1 |  2 |  3 | 4 | 5 |
|:---:|:--:|:--:|:--:|:-:|:-:|
| 172 | 16 | 17 | 10 | 1 | 1 |


The AMSNetId is purely logical and has usually no relation to the IP address. The AMSNetId is configurated at the target system.
At the PC for this the TwinCAT System Control is used. If you use other hardware, see the considering documentation for notes about settings of the AMS NetId.


### State Flags
The first bit marks, whether it´s a request or response. The third bit must be set to 1, to exchange data with ADS commands.
The other bits are not defined or were used for other internal purposes.

Therefore the other bits must be set to 0!

| Flag   | Description              |
|--------|--------------------------|
| 0x0001 | 0: Request / 1: Response |
| 0x0004 | ADS command              |

Bit number 7 marks, if it should be transferred with TCP or UDP.

| Flag   | Description  |
|--------|--------------|
| 0x000x | TCP Protocol |
| 0x004x | UDP Protocol |


## Notice

The information on this page is for reference only. All copyright and trademarks are with their respective holders.
This information is not covered by the repositories license. Please see [https://infosys.beckhoff.com](https://infosys.beckhoff.com) for
more information.
