# CompleSAT
A Simple SAT Solver

I was wondering how hard it could be to make a SAT solver which is (more or less) functional.
Speed isn't a top priority of this project, however readablility of the code is.
I hope that, it can be used as a toy SAT solver to get a general grasp of what it's supposed to do.

## What if I want a real SAT solver
As you can tell by the simplicity of the code, CompleSAT isn't really meant as a
state of the art SAT solver (however having it be somewhat fast is on the wish list).

If you are looking for one in Go that'll solve your problems in acceptable time:
take a look at [Gophersat](https://github.com/crillab/gophersat) by the CRIL.

If you're just looking for SAT solvers in general there are too many to be listed
however you could take a look at [MiniSat](http://minisat.se) or an extention of that
solver called [MinisatID](https://dtai.cs.kuleuven.be/software/minisatid) developed
at the Computer Science department of the KU Leuven.