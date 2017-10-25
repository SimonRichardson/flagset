# flagset

Drop in replacement for golang flagset, that also reads environment vars.

## Introduction

This drop in replacement just augments what the go stdlib provides, by also 
reading in environment variables with the same name, but with the following 
rules: all flag names are uppercase and dots are replaced with underscores.

So a argument of `some.flag.name` becomes `SOME_FLAG_NAME`.

## Example

```go
flags := flagset.NewFlagSet("flags", flag.ExitOnError)
debug := flagset.Bool("debug", false, "debug flag")

if err := flags.Parse(args); err != nil {
  return nil
}

fmt.Println(*debug)
```

### Why?

This is a nonsense drop in replacement that provides environment look up without
the requirement of changing much code to just get things done.
