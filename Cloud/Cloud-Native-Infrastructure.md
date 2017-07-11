# Cloud Native Infrastructure - Achieving a type 1 civilization

## Abstract

In this talk we will look at the state of Cloud Native Infrastructure as it stands today.
We will explore lessons learned in the DevOps discipline over the past decade that ultimately lead us to wear we are now.
We explore concrete similarities and differences between traditional infrastructure and cloud infrastructure, and dicuss the value that hosting in the cloud can offer, as well as warn about the dangers.
We will deep dive into coding patterns that helo bring this new philosophy into existence.
Upon conclusion we will understand how we are not replacing humans with automation tooling, but rather opening exciting new doors for projects to help us achieve beautiful and powerful infrastructure

## What is Cloud Native

Over the past decade in my career I have developed many applications for many different reasons.
Everything from accounting softare, to ecommerce APIs, to automation tooling for scalabale web based systems.
Regardless of the type of application I was building they all shared one common fundamental.
They all were developed with the intention on being "deployed" somewhere.

In some cases the deployment would be to a data center, and in other cases the deployment would be straight to a virtual machine in the cloud.
Regardless of the intended end state, there was always a bridge.
I like to call it "The Infrastructure Bridge".
It was the philosophical boundry between the application as it existed in development, and the application as it was intended to exist in production.

Crossing the bridge meant clearly defining one or more of these "environments", so that some tooling could span them in an attempt to bridge the connection.

Cloud native infrastrucutre completely liberates itself from these old paradigms.
The infrastructure layers now are represented by software.
Good software offers powerful opinions, this is ultimately any software's true value.
The software that represents Cloud Native Infrastructure offers opinions, and helps to manage virtual hardware.

Which means we now are developing applications in possibly the most exciting point in the history of technological evolution.
The infrastructure that used to govern our application development is now falling by the wayside.
As our software is now able to manage it's own infrastructure. 

Infrastructure that we can now treat as software has and will continue to change the scope of software development for the course of humanity. This is Cloud Native.

## Infrastructure Software

Throughout history software engineers have developed brialliant patterns for abstracting complex ideas into code.
Now that infrastructure can be represented as software, we are watching the pattern of mapping the complex ideas of infrastructure into elegant software patterns.
Perhaps the most notable pattern is the reconciliation and controll loops found in Kubernetes core components.
This pattern has it's origin in advanced robotics, and has been battle tested over time to demonstrate stability.

The pattern is quite simple, yet extremely powerful.
The pattern is defined by 2 sub components, the intended state, and a control loop that will audit the intended state.
The state can be defined simply by updating a record in some data store.
The control loop will run in parallel and independent of the record updating mechanism.
The control loop will parse the record, and reconcile it against the state of the world.
If a delta is detected between the two states, the control loop will take action and alter the state of the world.

Imagine a robot with a directive that says "put your left foot here".
This would be comparable to a directive that says "move your left foot forward 6 inches".
Notice how the former will be independently succesful regardless of the previous position of the robot's foot, while the latter has a vulnerability in that knowing the previous location is critical before changing to the new location.
This introduces complexity, and high risk of failure.
We do not treat our infrastructure in the way of the latter.
We define intended state, and that is all that software will need to concern itself with.

Now imagine all of the infrastructure components that make up a modern web application.
Virtual machines, DNS, routing, a network layer and network topology.
There could be load balancing, storage and backup concerns, authorization and access control.
The list goes on and on.

Now imagine each of these infrastructure components represented in the form of an API object.
The software to create and manage these infrastructure components is now fairly simple.
Between the state enforcement model we borrowed from robotics, and the new representation of infrastructure as API objects, an operator can define an intended state, and almost as if by magic it will manifest itself into reality and become so.

As all software evolves, very elegeant and clean solutions typically stand out.
This is the case of CNCF's Kubernetes project. A clean, and simple abstraction for powerful infrastructure ideas.

## Software as Software

The infrastructure layer, that at one time was physically separated from the software (or application) layers of the traditional OSI model now runs on the same layer.
This creates a unique environment where applications can now manage and control their own infrastructure as needed.
This beautiful new masterwork of infrastructure as software, and software running in the same layers is what Cloud Native means to me.

The infrastructure does not care about what the software does, or advertises to do.
The software does not care about how the infrastructure layers control what they control.

Thus, creating a harmonious environment where both infrastructure and application software can live side by side, in a mutally exclusive relationship.

On that note, the software layers of the system haven't changed drastically at all.
DNS is still DNS, networking is still networking, and servers are still servers.

So adopting and moving an application to the cloud is fundamentally no different than it has been for many many years.

## Conclusion

Infrastructure as code now means so much more than it ever has in the past.
We desparetly need to free ourselves mentally from the old paradigms of developing software ON infrastructure.
We need to begin developing our software WITH infrastructure.

According to Dr. Kaku of MIT there are many types of potential civilization milestones a civilization can achieve.
The human race is slowly approaching a Type 1 civilization, but is at risk of failure.

It is commonly believed that the crux of achieving a type 1 civilization can ultimately be the breaking point for a species' evolution.

As it stands today, the global network connecting our planet together is the internet.
The circulatory system of our planet.
On a biological level, being around to watch the evolution of this system is astounding.
Imagine an anatomical circulatory system in that was able to control itself, and grow itself, according to the needs of the host.
This complex evolution has happened to our global network within the past few years.

Being able to contribute and push the limits of this new evolution will be imperative in developing our species towards a Type 1 civilization.

So what is Cloud Native infrastructure?
It's the ability for software to manage it's own infrastructure.
It's an exciting new evolution that we are lucky to be a part of.
It's possibly one of the biggest hurdles in helping our species achieve global harmony.




