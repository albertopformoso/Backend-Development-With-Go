# Sync Memory Access

## Wait Groups
Example diagrams:

<img src="../media/wg1.png" width="400px"/>

<img src="../media/wg2.png" width="400px"/>

<img src="../media/wg3.png" width="400px"/>


## Mutual Exclusion (Mutex)

Prevents from entering to more than one process at the same time into the _critic section_.
The _critic section_ is a fragment of code where a shared resource can be modified.

## Data race

Occurs when two goroutines access the same variable at the same time and at least one of the access is of write type.