# Abstract

In this talk we look at how Go has forever changed the cloud for better and for worse. We draw on real world examples and meaningful data sets. We look at concrete examples in OSS that reflect the constraints and features of Go, and explore how it has implicitly shaped the cloud and tech forever.

# Description

As an engineer who has worked intimately with the public clouds of the world I began to notice a trend in designing and building the clouds. In this talk I will present concrete data sets that provide meaningful statistics on the Go programming language and its use in building out the cloud as we observe it today.

The talk will look at tooling I have observed both internally and publicly in various public clouds. We will look at how Go is overwhelming the majority of code and look at the patterns we see in solving fundamental problems between the clouds. We will observe the whimsical nature of technologies coming to life at the same time independent of each other, and draw connections between them and the constraints and features of Go.

The talk will be highly philosophical, but also very concrete with its data. The talk will introduce doubt in using Go in the cloud, as well as reassure users we are heading in the right direction. 

I ultimately hope the audience walks away questioning their technical decisions and reevaluating them as they move forward.

## Overview

- We begin by defining the core constructs of Go at a high level, and reminding the audience of the importance of interfacing and concurrent architecture.
- We look at several successful open source Go projects in the cloud space (Kubernetes, Terraform, etc) and explore their design patterns.
    - We look at various encoding avenues, Go channels, and implicit satisfaction of interfaces and how it has shaped the user experience in the cloud
    - We compare the strong typing in Go to the overwhelming dynamic typing of commonly used encoding languages in the cloud (JSON, YAML, etc)
    - We join data between the Go standard library issue tracker, and open source features in the commit logs of these projects
- We question what the cloud would be like if another language with different constraints had "won" the cloud engineer's hearts
- We amalgamate data, user experience, and opinion together to deliver a powerful experience report for engineering in the cloud with Go


## Notes

I speak for the software, and spend most of my waking hours architecting and fixing application and infrastructure level concerns in the cloud. At the time of writing this I am a senior engineer on Azure, but have spent countless hours on architecting infrastructure and application tooling for AWS and Google Cloud.

I plan on working closely with my colleagues across the clouds to gain realistic, truthful, and powerful data for the presentation. There will be graphs, and the graphs will be powerful. I will do my best to consolidate data and experience reports and deliver them in a linear and easily digestible format for audience members of all levels.
