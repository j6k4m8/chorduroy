# chorduroy

Command-line tool to generate chord diagrams for any stringed instrument of any number of frets and strings.

Guitar, ukulele, and mandolin-tested. (5s banjo support coming!)


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

`chorduroy` will also intelligently guess when a diagram can be "offset" to show chord patterns higher on the fretboard:

```
chorduroy -f 787X -o fretOffset.png
```

<img src="docs/fretOffset.png" width=150 />
