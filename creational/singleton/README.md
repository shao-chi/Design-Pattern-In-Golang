# Singleton Pattern

*"The singleton pattern is a software design pattern that restricts the instantiation of a class to one "single" instance."*

* 唯一實例
* when exactly one object is needed to coordinate actions across the system.

* 過度使用變 anti-pattern
    * singleton 裡面邏輯要單純
    * 不要到處使用 get instance, 應考慮使用 dependency injection 將需要的類別注入到將要使用的類別

![singleton](https://refactoring.guru/images/patterns/diagrams/singleton/structure-en-2x.png?id=cae4853e43f06db09f249668c35d61a1)

### Implementation in Golang
1. [Using sync.Mutex](./useMutex/useMutex.go)
2. [Using sync.Once](./useOnce/useOnce.go)
