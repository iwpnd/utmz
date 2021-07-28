# utmz

Methods to get utm zone number, epsg code or proj4 string from latitude and longitude.

## installation

```
go get -u github.com/iwpnd/utmz
```

## usage

### func Zone(p Point)

Return the UTM zone number for a given point.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/utmz"
  )

func main() {
  p := utmz.Point{Lat: 52.25, Lng: 13.37}
  z, err := utmz.Zone(p)

  if err != nil {
      fmt.Println(err)
      return
    }

  fmt.Printf(z)
}
```

Results in

```
33
```

### func Epsg(p Point)

Return a UTM zones EPSG code at a given point.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/utmz"
  )

func main() {
  p := utmz.Point{Lat: 52.25, Lng: 13.37}
  epsg, err := utmz.Epsg(p)

  if err != nil {
      fmt.Println(err)
      return
    }

  fmt.Printf(epsg)
}
```

Results in

```
32633
```

### func Proj4(p Point)

Return an UTM zones proj4 string from a given point.

```go
package main

import (
  "fmt"

  "github.com/iwpnd/utmz"
  )

func main() {
  p := utmz.Point{Lat: 52.25, Lng: 13.37}
  proj4, err := utmz.Proj4(p)

  if err != nil {
      fmt.Println(err)
      return
    }

  fmt.Printf(proj4)
}
```

Results in

```
+proj=utm +zone=33 +datum=WGS84 +units=m +no_def +ellps=WGS84 +towgs84=0,0,0
```

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/utmz](https://github.com/iwpnd/utmz)
