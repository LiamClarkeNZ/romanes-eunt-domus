[![Build Status](https://dev.azure.com/steelskynz/RomanesEuntDomus/_apis/build/status/LiamClarkeNZ.romanes-eunt-domus?branchName=main)](https://dev.azure.com/steelskynz/RomanesEuntDomus/_build/latest?definitionId=3&branchName=main)


I may have taken Postel's law too far, but I did have fun doing so.

```shell
~/d/romanes-eunt-domus (main)> ./roman xiiiiiiv
9
IX
```

To add a little sanity, use the flag `-less-liberal` to prevent accidentally implementing parts of a roman numeral calculator:

```shell
./roman -less-liberal xiiiiiiv
IIIIIIV converts to a number <= 0, which is most likely not intended
```

References used:

* https://pkg.go.dev
* https://go.dev/blog/intro-generics
* https://en.wikipedia.org/wiki/Roman_numerals