# Abstraction

> Noun: The process of considering something independently of its associations or attributes.

## Abstractions in technology

We all remember the gang of four.
The four fundamental philosophies behind object oriented programming.
Encapsulation, Polymorphism, Abstraction, and Inheritance.

Of the 4, the one word I continue to hear routinely in the office is *abstraction*.
Abstraction is a simple and powerful idea that can be found everywhere throughout our existence.

The word represents the idea that technical, and powerful decisions can be made without knowing intimate details of the problem.
In technology this can be a good thing.
It allows an engineer to take simple action, without having to worry about the underlying implementation of how it's being done.
Let's take a function `GetRecord()` and say that it returns some encoded record.
The engineer now can focus on 2 sides of the contract.
The first being the underlying implementation and ensuring that it does what it advertises to do, that is, it will get a record from somewhere.
The second being the layer above the implementation that deals with records after they have been retrieved.

While thinking and working in the mental space of the latter of the two concepts, the engineer is liberated from the noise and complexity that comes from getting a record.
The lucky engineer now can focus on the part of the program that does something with said record.

This pattern was made popular by Object Oriented programming scholars, and is still widely seen in programming today.

## Abstractions in the workplace

I am transgender, and I am openly gay.

I am also a maintainer of the most successful opens source cloud project in history.
I have also contributed to the Go standard library.

These are facts about me, but should have little relevance in your ability to reason about me and judge my ability to perform in a given situation.

Let's say I needed to implement a Heap in a program.
Work asked this of me, and it was my job as a software engineer to be able to do this.
I would honestly struggle with this assignment.
I haven't implemented a Heap in years, and have never actually coded one in production before.
My prestegious credentials aside, there is a high chance that the new grad student sitting next to me could accomplish this task much better than me.

So the sitting here is simple, and it might be hard to miss.
Which is why I find it meaningful to compare the workplace, to a code repostiory.

The task is simple, implement a Heap.

The work should be abstracted far away from the implementation of the people who will be solving the problem.

Imagine the function `ImplementHeap()`.
One one hand, this could be implemented by the queer senior engineer in the room.
On the other hand, this could be implemented by the grad student who started on Monday.
The fact is, that having the clear boundry of abstraction to pull the calling layer away from the details is important.
Having an environment that is structured in such a way at first map appear to be blindly, but in the long run is quite powerful.

It could very well be one of the ways we begin to balance out the career tree, in favor of skill and not that of merit.

Which as a woman in technology, seems like a very nice idea to me.


## The dangers of over abstraction

Let's take our `ImplementHeap()` function again that will call out into the pool of engineers and ask them to code a Heap.

But this time, let's make our request even more ambiguous.
Let's change the function from `ImplementHeap()` to `ImplementSolution()`.
This way the engineer tasked with the problem can solve it as they wish, without the burden of being told how to do something.

At first this seems like a nice model for solving problems.
Don't be fooled, this model is extremely dangerous.

The engineers would inevitably come up with varying solutions, based on the way they perceive the problems.
Or by which bit of technology they want to play with.
This introduces a fuzzy layer of knowledge transfer where the calling function (the manager) makes the assumption that the implementation (the engineer) understands the problem.

While this can be solved by cleverly transferring the knowledge using documents, and graphs, it still introduces complexity and risk.
The Heap implementation however does not have this risk. As the engineer is tasked with building a finite tool.
Either it's a Heap, or it's not.

So being able to offer guidance on what is and isn't the right amounts of abstraction in a decision making process is a bit of an art.
Some are better at it than others, and some favor more abstraction over less.
Ultimately it's an opinion you will get to spend your lifetime developing.

## Magic

With assumption comes a bit of magic.
Magic in this case can be defined as an unexplainable action that takes place and can be measured.

With all abstraction comes some level of assumption from the calling perspective, and thus there will also be some degree of magic involved.
The key is to limit the magic that happens, and to keep the magic confined into a single responsibility.

So even though the abstraction introduces magic, a bit of magic is okay if properly handled.
In fact, that is part of the beauty of well implemented abstraction.

## Conscience abstraction is hard

Abstraction is easy to reason about and conceptualize philosophically.
Practicing abstraction in your day to day life is risky, and hard.
It is a task that requires work and effort as a human.

Imagine an engineer called Stephen.
He is one of the most brilliant minds on the team.
He has written books, and solved many of the teams problems in beautiful and elegant ways.
In fact, he is notably the brightest mind on the team.

Now imagine I told you can't speak.
And that he drools on himself a lot.
Also that not only is he one of the brightest minds on the team, but he is debatably the brightest mind in the world.

His name is Stephen Hawking.

Changing your thought patterns based on tiny bits of knowledge is one of the things that has helped the human race evolve and thrive over time.
We are conditioned since birth to be a data processing machine, and to constantly evaluate and measure stimulus as it comes in.

So abstracting yourself above these tiny implementation details actually goes AGAINST what our bodies and minds will instinctually do.
Thus making it one of the hardest and most fascinating disciplines a human can practice.

As a thought experiment try to consciously change your behavior.
As with all skills a human can learn, one must practice.

Practicing pulling yourself a layer above the implementation, and judging your environment without concern for it's trivial details.

## Conclusion

In a weird way, one must begin mentally dehumanizing their peers in order to truly love and respect them.
As the only way to reach pure balance, is take unimportant details (gender, skin color, race, religion) and make technical decisions based on skill and ability alone.



