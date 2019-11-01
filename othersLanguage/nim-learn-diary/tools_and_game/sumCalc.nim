import strutils, sequtils

proc sum(a: seq): float =
    var res = 0f
    for i in a:
        res += i
    res

while true:
    readLine(stdin).splitWhitespace.filterIt(it in "0".."9").mapIt(it.parseFloat).sum.echo