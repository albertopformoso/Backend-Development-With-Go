# 2.3 Composition instead of Inheritance

## Inheritance vs Composition
+ Inheritance has an "is a" relationship and composition **"has/uses"**.
+ Simple inheritance has 1 to 1 relationship, in composition the related objects is **1 to many**.
+ Inheritance can have many levels of type hierarchy, in composition **there is no hierarchy, there are relationships**, where the **_single responsibility principle_** and **dependency injection** can be applied.
+ In inheritance, there is polymorphism through inheritance, in composition **there is polymorphism through interfaces**.
+ Testing is made easier when composition is used, since, if dependency injection is used and it receives an interface, there are many possibilities to perform unit tests easily. In inheritance it is more difficult to do and manage.

## Other features with advantages in Go
+ Methods can be in different declared types, it's not unique to structs.
+ It has implicit Interfaces, which do not need to use the `implements` keyword. A type can satisfy multiple interfaces and need not be specified within the type or structure.
+ Allows field embedding, allows embedding types into other types (analogous to subclassing but not identical).