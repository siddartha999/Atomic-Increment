Non-Atomic vs Atomic version of the increment operation.

<img width="1138" alt="Non-Atomic-Increment" src="https://github.com/siddartha999/Atomic-Increment/assets/37662337/a7987d9a-fde2-48b7-b74d-fdd3f1155d42">

<img width="1199" alt="AtomicIncrement" src="https://github.com/siddartha999/Atomic-Increment/assets/37662337/cd80821c-76c2-43ba-be6d-20bc55a2005c">


As the increment (++) operation is inherently non-atomic, a pessimistic locking mechasim has been utilized to make it thread-safe i.e, produces predictable results.

Reference: https://youtu.be/4F-WiPFrPsA?si=BOiR3vyPPYIGtiv5
