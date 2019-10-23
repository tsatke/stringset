# stringset

Stringset is used to check whether a string is part of a string slice. At the moment, this is only efficient for large slices, approximately 10000+ elements. Also, this is very memory consuming. For a common english text with around 2400 words, the set takes up about 704kB of memory. You should probably rather use a `map[string]bool` (benchmarks below), which also only uses about 100kB for the same 2400 words (`Benchmark 3`).

However, it should be mentioned, that the performance of the map approach seems to decrease, the farther a word is towards the end of the text (`Benchmark 3`, `MapContains` decreases by `2ns` for different terms), which is why we approximate the feasability of this tool to >10000 elements (not benchmarked).

This tool was created to proof the idea, that such a task can be solved by creating a tree from all letters and checking whether a given path can be walked (exists).

## How to use it
First, go get it.
```
go get github.com/TimSatke/stringset
```

What you can do with `stringset`, is this.
```go
set := stringset.New([]string{"abc", "def", "abf"})
...
func Contains(term string) bool {
    return set.Contains(term)
}
```

## Performance and why you should use `stringset`

As you can see in `Benchmark 1`, depending on where the element is in a slice, the traditional approach (iterating over a string slice and checking for equality) can produce very different results.
In `Benchmark 2` you can see, that it is either way faster than `stringset`, or way slower.

The traditional approach is slower on average while having deltas of up to 125%, whereas the `stringset` is faster on average and more stable, independent of the location of a word in the slice or if it's even in the slice.

##### Benchmark 1
```
name                   time/op
SetContains-8           61.8ns ±13%
StringSliceContains-8  1.07µs ±106%
MapContains-8           10.2ns ±23%
```

##### Benchmark 2
* `Utterson` 2nd word in the text
* `Utterson;` 233rd word in the text
* `bargain` 2375th word in the text
* `question;` not in the text at all
* total length: 2394 words in the text
```
name                             time/op
SetContains/Utterson-8           60.5ns ± 2%
SetContains/Utterson;-8          67.0ns ± 2%
SetContains/bargain-8            54.9ns ± 3%
SetContains/question;-8          64.8ns ± 1%
StringSliceContains/Utterson-8   3.21ns ± 0%
StringSliceContains/Utterson;-8   180ns ± 0%
StringSliceContains/bargain-8    2.20µs ± 1%
StringSliceContains/question;-8  1.92µs ± 0%
MapContains/Utterson-8           8.72ns ± 1%
MapContains/Utterson;-8          9.08ns ±12%
MapContains/bargain-8            10.6ns ±18%
MapContains/question;-8          12.4ns ± 1%
```

##### Benchmark 3
```
name                             time/op
New-8                            1.00ms ± 0% // time it takes to create a set from the given words
NewMap-8                          139µs ± 0% // time it takes to create a map[string]bool from the given words

name                             alloc/op
New-8                             705kB ± 0%
NewMap-8                         96.5kB ± 0%

name                             allocs/op
New-8                             13.2k ± 0%
NewMap-8                           37.0 ± 0%
```
