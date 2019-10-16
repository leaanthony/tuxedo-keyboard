# Go Tuxedo Keyboard

A simple library for interfacing with Tuxedo compatible keyboards.

## Requirements

This module relies on the fact that the [Tuxedo Keyboard Kernel Module](https://github.com/tuxedocomputers/tuxedo-keyboard) has been installed on your machine.

## Example

```
...
SetBrightness(128) // Half bright
SetKeyboardColourLeft(255,0,0) // Set left lights to Red
...
```

## Status

Works for a Metabox Alpha-X. 
`SetKeyboardColourRight`,`SetKeyboardColourCenter` and `SetMode` untested as my machine doesn't use them.