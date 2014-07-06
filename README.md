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

Consider this file containing multiple fields, semi aligned with spaces and tabs
and having a semicolon as separator as well:

<pre>
Num     Time   IsoTime  Other Separator
1 	    2h3m   3:01     ; 4
2 	    3h4m   2:01     ; 4
-3.5 	4h5m   4:01     ; 4
4.3 	5h6m   8:01     ; 4
5 	    6h7m   3:01     ; 4
foo     bar             ; 4
6.765 	7h8m   3:01     ; 4
7 	    8h9m   3:01     ; 4
8 	    9h10m  3:01     ; 4
9 	    10h11m 3:01     ; 4
10	    11h12m 12:01    ; 4
11	    12h13m 33:01:01    ; 4
12	    13h14m 3:01     ; 4
</pre>

What if you wanted to figure out the sums of the different columns?

Executing the summer command will sum the selected column.  The input will be
printed as well to make easy to select these lines in vim and retrieve the
result in an additional line below the selection.

### Let's sum the first column

<pre>
> summer < FILE
Num     Time   IsoTime  Other Separator
1 	    2h3m   3:01     ; 4
2 	    3h4m   2:01     ; 4
-3.5 	4h5m   4:01     ; 4
4.3 	5h6m   8:01     ; 4
5 	    6h7m   3:01     ; 4
foo     bar             ; 4
6.765 	7h8m   3:01     ; 4
7 	    8h9m   3:01     ; 4
8 	    9h10m  3:01     ; 4
9 	    10h11m 3:01     ; 4
10	    11h12m 12:01    ; 4
11	    12h13m 33:01:3600    ; 4
12	    13h14m 3:01     ; 4
72.57
</pre>
The result is 72.57.

### Let's sum the third column without printing the input

<pre>
> summer -f 3 -n < FILE
80h12m1s
</pre>
The result is 80h12m1s.

### Let's sum the fourth column

<pre>
> summer -f 4 -n < FILE
4.00
</pre>
Wait, the result is 4.00.  Shouldn't it be much higher?

### Let's sum the fourth (second) column by changing the separator

<pre>
> summer -f 2 -n -d \; < FILE
52.00
</pre>
This looks much better.  The result is 52.00.

## Todo

* Make line separator cross platform, function SumString
* Support binary, octal and hex and values
* Support negative fields that are counted from the right
* Return a proper error code when values couldn't be parsed
* Improve output of time values

# Done
* Add examples, 2014-07-06
* Support time values exceeding 24 hours, 2014-07-06
* Support negative times in -3:04 format, 2014-07-06
* Support summing time durations like 2:03, 2014-07-06
* Support summing time durations like 2hm3, 2014-07-06
