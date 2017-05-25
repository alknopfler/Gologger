# Gologger
Logger module for golang based on logrus module.
This module was created to reduce the logrus abstraction in order to use
just 2 functions:
- Init() : initialize the logrus params
- Print(): Print an error with the base fields

**Usage:**

** Import package **
```
import (

	"github.com/alknopfler/Gologger/gologger"
	)
```

**Init with stdout and info level:**

  ``` 
  gologger.Init(os.Stdout, logrus.InfoLevel)
  ```

**Print to show the error message:**

```
gologger.Print("WARN",7,"Description to show","filename.go")
```






