# bandwidth-incentive-simulation
This repository contains implementation of the swarm bandwidth incentive simulation written in Go.

# Instructions to Compile and Run System
Ensure Golang, preferably version 1.19.5 or later, is installed on the computer.

Clone the GitHub repository:
```$ git clone git@github.com:Swarm-Bachelor/go-incentive-simulation.git```

Change directory to ./go-incentive-simulation:
```$ cd go-incentive-simulation```

Configure the settings for the simulation by editing the config.yaml file:
```$ nano config.yaml```

Run the program:
```$ go run main.go```

View output in terminal and in ./results folder:
```$ cd results```
```$ cat *fileName*.*extension*```

Generate new network files, using config.yaml for settings:
```$ cd data```
```$ go run generate_data.go```



## Repository Transition Notice

**This repository is the new home of the project.** 

The original project was previously hosted in a separate repository, which can be found at [link-to-old-repo](https://github.com/Swarm-Bachelor/go-incentive-simulation). The old repository contains the earlier versions, historical information, and previous contributions.

Going forward, all development, bug fixes, and new features will be implemented in this repository. Please refer to this repository for the latest updates and active development.
