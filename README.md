# botholiday

`botholiday` is a Go package designed to help check the holidays of the Bank of Thailand.

### Install

```bash
go get github.com/oddsteam/bot-holiday
```

### Usage

You can register and obtain a Autherized token from [Here](https://portal.api.bot.or.th/).

```go
import (
    "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

botholiday.Initialize("AccessToken")
```

### Example

Get all holidays of the Bank of Thailand in this year.

```go
import (
    "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

inst := botholiday.Initialize("eyJvcmciOiI2NzM1...")
inst.GetBOTHoliday()
```

Check if today is a holiday of the Bank of Thailand for this year or not.

```go
import (
    "time"
    "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

inst := botholiday.Initialize("eyJvcmciOiI2NzM1...")
if botholiday.CheckHoliday(inst.GetBOTHoliday(), time.Now()) {
    // Do something if today is a holiday.
}
```

### Maintainer

[ODDS](https://github.com/oddsteam)

### License

MIT
