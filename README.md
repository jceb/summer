summer
======

<pre>
Usage of summer:
	Sum values in column and print result
Options:
  -f 1  --field=1     Selected field
  -d    --delimiter=  Use delimiter instead of space-like characters
  -n    --no-print    Don't print input
  -p    --print       Print input
  -s 2  --scale=2     Scale to number of digits after the decimal point
        --help        show usage message
</pre>

## Examples

to come

## Todo

* Make line separator cross platform, function SumString
* Support binary, octal and hex and values
* Support negative fields that are counted from the right
* Return a proper error code when values couldn't be parsed

# Done
* Support negative times in -3:04 format, 2014-07-06
* Support summing time durations like 2:03, 2014-07-06
* Support summing time durations like 2hm3, 2014-07-06
