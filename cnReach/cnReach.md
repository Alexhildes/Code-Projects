# cnReach Testing

## Test Setup - Solcomm Office - Point to Point Network

```mermaid
graph LR;
L1[Laptop 1: 192.168.15.12]
L2[Laptop 2: 192.168.15.14]

C1[cnReach 1: 192.168.15.191]
C2[cnReach 2: 192.168.15.190]
subgraph Server
L1--Eth--> C1
end

subgraph Client
L2--Eth-->C2
end

subgraph Link
EndP-cnReach2--Attenuator-->AccessP-cnReach1
end
```
## iperf

iperf is used to test the throughput of the radios. This tests the Client to Server communication link.

Cd into the iperf directory

For the Server computer (Laptop 1)

```
iperf -s
```

For the Client computer (Laptop 2)

```
iperf -c <Server IP> 
```

## Preliminary Testing

### PTP Network

@50 dB attenuation

- Rx Only
- Tx Only
- Bi-directional

Testing from cnReach 2 Web UI

![image](images/throughput_PTP.png)

```
------------------------------------------------------------
Server listening on TCP port 5001
TCP window size: 64.0 KByte (default)
------------------------------------------------------------
[  4] local 192.168.15.12 port 5001 connected with 192.168.15.14 port 49279
[ ID] Interval       Transfer     Bandwidth
[  4]  0.0-14.4 sec   512 KBytes   292 Kbits/sec
```

### PTM Network

```mermaid
graph LR;
L1[Laptop 1: 192.168.15.12]
L2[Laptop 2: 192.168.15.14]

C1[cnReach 1: 192.168.15.191]
C2[cnReach 2: 192.168.15.190]
C3[cnReach 3: 192.168.15.192]
S1[Splitter]
R1[50 Ohm]

subgraph Server Side
L1--Eth--> C1
end

subgraph Client Side
L2--Eth-->C2
L2--Eth-->C3
end

subgraph Radio Link
EndP-cnReach2-->S1
S1-->AccessP-cnReach1
EndP-cnReach3-->S1
R1-->S1
end
```