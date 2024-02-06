# map-reduce-golang

# Disclaimer

This project is meant as a personal learning experience and does not represent professional quality code.

Best effort is made to keep good coding standards whilst working at speed.

Constructive feedback is ecnouraged and welcomed.

# Motivation

The motivation behind this implementation of MapReduce is to dig deeper into the Go programming language and learn how concurrency, synchronisation (mutual exclusion, signalling, rendezvous etc) and parallelism works

# Overview

Tis is a simple implementation of the [MapReduce](https://en.wikipedia.org/wiki/MapReduce) programming model for handling and processing large data sets concurrently.

As the name suggest, the user specifies a `Map` function that processes input data and produces an indermidiate set of key/value pairs.

Then we have a `Reduce` function which takes the intermidiate result produced from the `Map` and merges them to produce the final result.

This programming model is really powerfull when parallelised and distributed across many machines - that way we can process big data quicker.

# Implementation

TODO: explain current implementation goals

# Fault Tolerance

TODO: explain fault tolerance techniques used or available or planned

# Contribution Guidelines
