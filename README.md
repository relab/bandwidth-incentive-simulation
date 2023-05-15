# go-incentive-simulation
This repository contains implementation of the swarm bandwidth incentive simulation written in Go.


# Instructions to Compile and Run System
Ensure Golang, preferably version 1.19.5 or later, is installed on the computer.

Clone the GitHub repository.
$ git clone git@github.com:Swarm-Bachelor/go-incentive-simulation.git

Change directory to ./go-incentive-simulation.
$ cd go-incentive-simulation

Configure the settings for the simulation by editing the config.yaml file.
$ nano config.yaml

Run the program.
$ go run main.go

View output in terminal and in ./results folder
$ cd results
$ cat *fileName*.*extension*

Generate new network files, using config.yaml for settings
$ cd data
$ go run generate_data.go
