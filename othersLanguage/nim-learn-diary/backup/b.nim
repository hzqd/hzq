import sequtils
proc lam(it: int): int = 
    if it <= 2: 1 else: lam(it-1)+lam(it-2)

(1..20).toSeq.map(lam).foldl(a+b).echo