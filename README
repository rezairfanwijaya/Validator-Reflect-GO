
# Reflect Validator

Project ini adalah project iseng membuat validator pada sebuah struct menggunakan package [Reflect](https://pkg.go.dev/reflect). Contoh penggunakan validator ini bisa dipakai ketika mengharuskan user mengisi sebuah field, seperti username ataupun password. Anda hanya perlu menambahkan tag `required:"true"` pada filed yang diwajibkan diisi oleh user.


## Documentation

#### Installation

```bash
 go get github.com/rezairfanwijaya/Validator-Reflect-GO
```

#### Example
```bash
import (
	"fmt"

	validate "github.com/rezairfanwijaya/Validator-Reflect-GO"
)

type SuperAdmin struct {
	Name    string `required:"true"`
	Age     int    `required:"true"`
	Address string `required:"true"`
}

func main() {
	SA := SuperAdmin{
		Name:    "SuperAdmin",
		Age:     20,
		Address: "Jakarta",
	}

	valid, err := validate.Validate(SA)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(valid)
	}
}
```

```bash
  output : true
```



