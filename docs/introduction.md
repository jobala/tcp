## Introduction

This project is about a user-space networking stack written in Go. It uses a tap device to access ethernet frames and pass them up the protocol suite for processing. 

## What Is A Tap Device ?

To understand a tap device, you first need to understand how a computer accesses the internet which is usually through a router. When a computer makes a request, the data is first sent to a router and then the router sends it to the internet. When a response is sent back, the router receives it first and then sends it to the computer.

Data between the computer and router is exchanged using ethernet frames which contain among other properties the destination MAC address. When the computer receives an ethernet frame the network interface chip checks if the frame's destination address matches it's MAC address. If it does, it accepts the frame and passes it to the network stack for processing.

In our case, the tap device plays the same role as the network interface accepting ethernet frames and passing it to the network stack. This project implements the network stack that receives and processes these frames.


