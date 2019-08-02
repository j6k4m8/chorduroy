# chorduroy

Command-line tool to generate chord diagrams for any stringed instrument of any number of frets and strings.


```shell
chorduroy -f X554X5 -o Am6.png
```

<img src="docs/demo.png" width=150 />


## Usage

```
  -f string
    	Fingering (from highest) (default "X554X5")
  -o string
    	Path at which to save diagram (default "diagram.png")
  -s float
    	Number of frets to include (default 6)
```


`chorduroy` guesses the number of strings on your instrument based upon the fingering you provide. For example, if there are four characters in your fingering, the output will have four strings:

```shell
chorduroy -f 232X -o ukeG.png
```

<img src="docs/ukeG.png" width=150 />

