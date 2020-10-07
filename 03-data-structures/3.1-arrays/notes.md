# class notes

Slices are the most important data structure in Go

Today's clocks run at 3Ghz

Arrays are most important data structure as it related to hardware
- contiguous memory is more performant that the pre-fectcher can use to find cadence to access memory
- if so predicatble strides can easily be picked up by pre-fectcher

Cache Line grabs the memory values before you need them when data stored in arrays
Cache ines are ~64mb in size

Go has to be mechanically sympathetic because go doesn't use a vitrual machine like Java or C# 
- where the vm can align data structures to the metal.

With Go, the machine is the model, so go only uses arrays, slices and maps 
- so Go has to data structures that align with the metal (memory)

So, contiguous data, predictable access patterns over fancy algorythms

