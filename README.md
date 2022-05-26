
# Vitals CLI

[ 🔋 ] Vitals is a CLI written in Go made to quickly retrieve Web Vitals
from websites using the PageSpeed Insights API & Lighthouse Scores from Google.

[ 🐍 ] Powered by Cobra: https://github.com/spf13/cobra

## Available Commands

### `vitals [command]`

#### `vitals check  -u`
Retrieve lighthouse information on the specified domain name


For example:

`vitals check  -u example.com`

wil return:

```bash
Lighthouse Version: 9.3.0
Lighthouse Fetch Time: 2022-05-26T14:24:07.609Z
Lighthouse Seo Score: 0.91
Lighthouse Best Practices Score: 1
Lighthouse Performance Score: 1
Lighthouse Accessibility Score: 0.92
```


## Contributing

Contributions are always welcome!

