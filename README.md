Env
---
A simple wrapper to get environment variables.

## Usage
```
env.Prefix = "TEST"

env.Get("VARIABLE_NAME", "default") //actually gets TEST_VARIABLE_NAME
env.GetInt("VARIABLE_NAME", 1) //returns a int
```