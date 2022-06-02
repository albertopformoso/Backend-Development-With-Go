# 2.1 Structures Instead of Classes

Before explaining how to implement Object Oriented Programming (OOP), first we need to define if Go is a OOP language. The answer as the creators says is yes and no.

First We need to explain the four pilars of the OOP and explain how to implement them in Go:

1. **Abstraction**: is the mental process that we do to isolate an element from the real of ficticious world. For instance if you think of a car, what you will do is this mental process of abstraction to think on what are the  characteristics or actions that the car have like the color, brand, how to start the engine, etc. You can also make an abstraction of elements that aren't real but are ficticious. For example a superhero, despite this element doesn't exist you will think on some characteristics or actions that the superhero will have, like the superpowers or suit he/she will have.

    This mental process is called abstraction, but the abstraction on a OOP language is this characteristic that this languages allows so we can portray these actions or characteristics that our elements or objects have.

    In Go we don't have classes as in other programming languages, but it have structures (_structs_) wehere we can portray this characeristics through the fields of the structure and the actions through the methods (_structs and receivers_).

    In Go the methods aren't exclusive for _structs_ you can apply methods to _maps_, _fucntions_, _arrays_, etc.

2. **Encapsulation**: is the characteristic that a language provide to protect the properties or mehtods of the user that is working with the classes.

    Go have its own way to encapsulate with **imported and exported modifiers** where at the package level you can indicate what you want and don't want to export so the user that will work with the package knows what he can and can't use.

3. **Inheritance**: this concept is easier to understand since we can watch it in the real world, for example we inherit the eye color, skin tone and hair and also you can inherit these characteristics to your son and some more characteristics.

    In Go there's no inheritance, since you have a type hierarchy which you can implement to have types integrated to types, similar to subclasses in other programming languages.

4. **Polymorphism**: is the characteristic that a programming language provide to send messages syntactically similar and that objects of different types can respond to these messages.

