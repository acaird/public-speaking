# Over engineering cloud native applications with Go

### Abstract

As the cloud native realm of technology begins to grow, we will naturally find users and applications shifting into this new way of engineering and thinking.
Implicitly, we now have a large rise in newcomers to the field, and thus a high probability of mistakes in our software.
We also now have a new and powerful tool kit that has never been seen before in technology.

In this talk we explore cloud native design patterns, and the thinking that lead to them.
We explore overwhelming successes, and catastrophic failures in the new world of software that we enviably now find ourselves in.

The user will walk away with a clear understanding of my core values developed from the last few years of building cloud native applications:

> Make it work, Make it work well, Make it pretty

### Complexity

A modern day cloud can be fairly complex, but they're not THAT complex.
Most of the complexity lies in understanding a cloud's unique resources, and how the interface together.
For the most part, all clouds can be broken down into 4 main components that will make up the majority of cloud native applications.
These components listed in dependency order are:

 1. Network
 2. Compute
 3. Storage
 4. Authentication

Focusing on the big picture makes tackling these resources much easier.
Here I would give examples on networking resources, and ways to code powerful apps around them by lumping them together, instead of pulling them apart.
I would use the Kubernetes kops implementation of private network topology that I coded last year as one example.
I would show examples of Kubernetes source code that is pulled very far apart, and then show a working example of a much smaller piece of code that offers the same functionality and is easier to grok.

### Abstraction

Interfaces in Go are wonderful, especially for defining cloud implementation.
A user could have a DNS implementation for a number of different clouds and scenarios.
One of the major attractions to Golang is it's enforced implicit interface satisfaction.
That is to say, all interfaces in Go are implemented implicitly.

But ye be warned, over using interfaces can be a developer nightmare.



### Developer Empathy



