# botholiday

`botholiday` is a Go package designed to help check the holidays of the Bank of Thailand.

### Install

```bash
go get github.com/oddsteam/bot-holiday
```

### Usage

You can register and obtain a ClientID from [Here](https://apiportal.bot.or.th/bot/public/start).

```go
import (
    "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

botholiday.Initialize("ClientID")
```

### Example

Get all holidays of the Bank of Thailand in this year.

```go
import (
    "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

inst := botholiday.Initialize("Client ID.")
inst.GetBOTHoliday()
```

Check if today is a holiday of the Bank of Thailand for this year or not.

```go
import (
    "time"
    "github.com/oddsteam/bot-holiday/pkg/botholiday"
)

inst := botholiday.Initialize("Client ID.")
if botholiday.CheckHoliday(inst.GetBOTHoliday(), time.Now()) {
    // Do something if today is a holiday.
}
```

### Maintainer

[ODDS](https://github.com/oddsteam)
