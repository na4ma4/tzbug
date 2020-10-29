# tzbug

Testing [timezone bug](https://github.com/golang/go/issues/42138) in golang

## How to use

```shell
git clone git@github.com:na4ma4/tzbug.git
cd tzbug
make
```

## Expected Output

### Alpine 3.8

```text
2020/10/29 01:05:41 Expected time stamp returned for Europe/Berlin
2020/10/29 01:05:41     Expected: 2020-10-29 15:30:00 +0100 CET
2020/10/29 01:05:41     Actual:   2020-10-29 15:30:00 +0100 CET
2020/10/29 01:05:41 Expected time stamp returned for Australia/Brisbane
2020/10/29 01:05:41     Expected: 2020-10-29 15:30:00 +1000 AEST
2020/10/29 01:05:41     Actual:   2020-10-29 15:30:00 +1000 AEST
2020/10/29 01:05:41 Expected time stamp returned for Australia/Sydney
2020/10/29 01:05:41     Expected: 2020-10-29 15:30:00 +1100 AEDT
2020/10/29 01:05:41     Actual:   2020-10-29 15:30:00 +1100 AEDT
```

### Alpine 3.12

```text
2020/10/29 01:05:43 Expected time stamp returned for Australia/Brisbane
2020/10/29 01:05:43     Expected: 2020-10-29 15:30:00 +1000 AEST
2020/10/29 01:05:43     Actual:   2020-10-29 15:30:00 +1000 AEST
2020/10/29 01:05:43 Incorrect time stamp returned for Australia/Sydney
2020/10/29 01:05:43     Expected: 2020-10-29 15:30:00 +1100 AEDT
2020/10/29 01:05:43     Actual:   2020-10-29 15:30:00 +1000 AEST
2020/10/29 01:05:43 Incorrect time stamp returned for Europe/Berlin
2020/10/29 01:05:43     Expected: 2020-10-29 15:30:00 +0100 CET
2020/10/29 01:05:43     Actual:   2020-10-29 15:30:00 +0200 CEST
```

## References

Based on [this gist](https://gist.github.com/hlubek/f46a73bc9d150cf1f2af585b0849e3d9)
