# stringset

Stringset is used to check whether a string is part of a string slice.

For this task, you would usually create a string slice and use a `for`-loop, like this.
```go
mySlice := []string{"abc", "def", "abf"}
...
func Contains(term string) bool {
    for _, elem := range mySlice {
        if elem == term {
            return true
        }
    }
    return false
}
```

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
SetContains-8           62.4ns ±12%
StringSliceContains-8  1.23µs ±125%
```

##### Benchmark 2
* `Utterson` 2nd word in the text
* `Utterson;` 233rd word in the text
* `bargain` 2375th word in the text
* `question;` not in the text at all
* total length: 2394 words in the text
```
name                             time/op
SetContains/Utterson-8           60.9ns ± 1%
SetContains/Utterson;-8          67.6ns ± 2%
SetContains/bargain-8            55.2ns ± 1%
SetContains/question;-8          66.0ns ± 0%
StringSliceContains/Utterson-8   3.26ns ± 0%
StringSliceContains/Utterson;-8   184ns ± 0%
StringSliceContains/bargain-8    2.74µs ± 1%
StringSliceContains/question;-8  1.98µs ± 1%
```