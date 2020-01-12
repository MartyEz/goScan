# goScan

This repo is a simple scanner in Go

It uses semaphore to limit the number of opened sockets.

It uses goroutine to speed up scanning

The scanner communicates throught a tcp stream as it's used in a revershell