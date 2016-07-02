# grpczap

Under development.

## Usage

```golang
import (
	"github.com/uber-go/zap"
	"github.com/kazegusuri/grpczap/zapctx"
	"github.com/kazegusuri/grpczap"
	"golang.org/x/net/context"
)

func main() {
	logger = zap.NewJSON()
	uIntOpt := grpc.UnaryInterceptor(grpczap.UnaryZapHandler(logger))
	grpc.NewServer(uIntOpt)
}

func Foo(ctx context.Context) {
	logger := zapctx.MustFromContext(ctx)
	logger.Info("message")
}
```

## Copyright

<table>
  <tr>
    <td>Author</td><td>Masahiro Sano <sabottenda@gmail.com></td>
  </tr>
  <tr>
    <td>Copyright</td><td>Copyright (c) 2016- Masahiro Sano</td>
  </tr>
  <tr>
    <td>License</td><td>MIT License</td>
  </tr>
</table>
