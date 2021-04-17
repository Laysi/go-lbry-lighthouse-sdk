# go-lbry-lighthouse-sdk
A golang sdk client for LBRY Lighthouse server

## example:
### AutoComplete

```go
import	lighthouse "github.com/Laysi/go-lbry-lighthouse-sdk/sdk"

func main(){
    client:=lighthouse.NewAPIClient(lighthouse.NewConfiguration())
	result, _, err := client.LighthouseApi.SearchAutoComplete(context.Background()).Nsfw(false).S("dragon").Execute()
	if err != nil {
		panic(err)
	}
	print(result);
}

```