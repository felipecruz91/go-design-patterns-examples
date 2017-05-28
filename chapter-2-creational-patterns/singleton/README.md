# Singleton creational pattern

It will provide you with a single instance, it is created and then reused between all the parts in the application that need to use that particular behavior.

# Situation

We consider using the `Singleton` pattern when the following rule applies:

- We need a *single*, *shared* value, of some particular type.
- We need to *restrict object creation* of some type to a *single unit* along the entire program.

# Cases

- When you want to use the same connection to a database to make every query.

- When you open a Secure Shell (SSH) connection to a server to do a few tasks, and reopen the connection for each task.

- If you need to limit the access to some variable or space, you use a `Singleton` as the door to this variable.

# Example

We will write a counter that holds the number of times it has been called during program execution. It should not matter how many instances we have of the counter, all of the must *count* the same value and it must be consistent between the instances.

# Acceptance criteria

- When no counter has been created before, a new one is created with the value 0.
- If a counter has already been created, return this instance that holds the actual count.
- If we call the method `AddOne`, the count must be incremented by 1.