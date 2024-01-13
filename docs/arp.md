In a IPv4 network 32 bit IP addresses are used as the source or destination of information but the exchange of ethernet frames use 48-bit MAC addresses. Address resolution is the process of discovering mapping from one address to another. ARP provides a dynamic mapping between IP and MAC addresses and operates only when reaching local servers in the same IP subnet. If a computer tries to reach a remote server then a router is required to reach the destination.

## Direct Delivery And ARP

Assuming there exists a server in a network at address 10.0.0.1 and we have a client with the same IP prefix -- local to the network -- in this case a browser. When the browser sends a request to the server, the browser makes a tcp connection to the server then sends a TCP segment to the server by sending a IPv4 datagram. Since the server has the same IP prefix, the datagram can be sent directly without a router.

The machine running the browser must convert the 32-bit IP address to the 48-bit MAC address to send the datagram directly to the machine running the server.ARP is responsible for the conversion of 32-bit IP addresses to 48-bit MAC addresses. 

ARP works in broadcast networks where the link layer delivers a message to all attached networked devices and to the conversion of IP to MAC addresses ARP sends a frame called ARP request to all attached network interfaces, the frame contains the destination IP address. ARP request tells devices on the network to respond with their MAC address if their IP matches the destination IP address. 

Only the system with the matching IP address responds with an ARP reply which contains its MAC address. After the resply is received by the sender, the datagram is sent directly to the machine running the server.

## ARP Cache

ARP Cache is maintained on each host and contains recent IP-MAC address mappings for each interface using address resolution. Normally cache entries expire after 20 minutes.
