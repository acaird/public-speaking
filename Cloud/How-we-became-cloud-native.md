# Understanding cloud native infrastructure, and the software that drives it.

By digging into succesful design patterns in many modern day cloud native applications we can begin to imagine a world where we take the next step into infrastructure as well.
Imagine a world where the paradigms we are seeing that run and orchestrate cloud native applications, begin to run and orchestrate our infrastructure as well.
In this talk we will explore the evolution of infrastructure management for humans over time.
We begin by discussing the first step of primitive automation by running selective scripts on infrastructure to provision it.
We then begin looking into infrastructure as code tooling such as terraform, or ansible. 
We take the leap into API centric infrastructure that we have seen in Kubernetes kops.
Then finally we paint a picture of "tomorrows" infrastructure using tools like kubicorn that are designed around a reconciliation enforcement model and intended to run as a control loop or operator pattern thus enforces infrastructure over time.
After the history is complete, and the user now has a good understanding of where we are, and where we want to go we can begin to fantasize about how infrastructure might look in the future.


We explore the software that would power these futuristic tools.
We lean how the Go programming language has made these leaps in infrastructure maturity possible. 
We also share the woes, and painful lessons learned along the way, and remind ourselves what traps we can fall in moving forward.
The user will walk away from the talk feeling excited about cloud native infrastructure, and hopefuly considering switching their own infrastructure management proccess over to the new paradigms presented in the speech.
