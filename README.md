### Overview

The test file contains tests for a FlattenMap and a FlattenJson method. Starting with FlattenMap,
write the code to make the test pass. If you get to a solution, uncomment the FlattenJson test
and make that pass to.

```shell
go test
```

### Flatten Logic

Take a map of the form:

```
{ "a": { "b": { "c": 1, "f": 2 }, "g": 3 }, "d": 2 }
```

And return a map where properties are only one level deep, with their depth represented in dot notation.

```
{ "a.b.c": 1, "a.b.f": 2, "a.g": 3, "d": 2 }
```

### Implementation Notes

- Do not concern yourself with primitives, edge cases, security, or performance;
- Assume you have safe, valid (but arbitrarily nested) objects;
- You can change everything you want except the test itself;
- Assume there are no circular references in the objects you receive;
- You can search things online, as long as you're not searching for a solution;
- Do not use Github Copilot, Amazon CodeWhisperer, or similar tools;
