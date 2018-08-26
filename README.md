# geojsondif

### What is it?

Writing tests to compare geospatial data is kind of shitty especially when you have a few representations that you may shuffle between. Geojson is always a format that is used somewhere in the stack at least for my workflows. So this library essentially compares two geojson features from different representations and ensures there equal or returns a reason for difference. 

Currently this library assumes a 6 digit decimal espg 4326 precision.

**Basically two geojson features go in -> an error explaining the delta comes out.**

### Usage 

```golang
package main

import (
  "github.com/paulmach/go.geojson"
  "github.com/murphy214/geojsondif"
  "fmt"
)

func main() {
  feat1 := // your geojson feature
  feat2 := // your geojson feature
  err := geojsondif.CompareFeatures(feat1,feat2)
  if err != nil {
    fmt.Println(err)
  }
}
```

### Caveats

Doesn't support geometry collections, and currently error returned is rather generic could easily add in more specific errors messages. 
