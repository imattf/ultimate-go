the Go garbage collector is a...
- non-generational
- non-compacting
- concurrent
- tri-color (white, grey, black); white = no longer in use; black = in use; grey = being marked
- mark and sweep 
...collector

mechanics...
Memory on the heap is not moving; fixed for the life-time of that program

But let's focus on the semantics and not the mechanics of the GC...

should GC be performant or 
application thru-put while GC happening or...

Focus on what Go takes as a priority
makes sure app has highest app thru-put
w/ today's cloud resources, try to keep the resource usage down
minimize latency; maximize thru put

When should the first GC happen?
There is a knob in Go GC: GOGC = 100 (100%)
First GC starts when first 4meg
STW = Stop the world: ~ < 100 micro secs
Mark start (STW) + Marking + Mark Term (STW)

Mental model...
When a Go program starts up we get an OS thread for every thread can run in parallel
4 thread go program
GC will take 25% of available CPU capacity

Application work goes from...
100% CPU = 4G -> 75% CPU 3G

Strategy: reduce need for GC in your app...
Tooling can show the amount of time we spend in GC