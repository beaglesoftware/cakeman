# How cakeman works?
Others publish their module for C/C++ on Cakeman confectionary.

You will add it to your Cake with `cman add <packagename>` and install it with `cman install`

For example, we download a package called `sayhello`

`sayhello` includes these in sayhello.h:
```c
#include <stdio.h>

void sayhi(msg) {
    printf("<" + msg + ">\n");
    printf("\n");
    printf("\n");
    printf("|--------/------|"); // Example: An art of MacOS Finder logo
    printf("|   *   /    *  |");
    printf("|      /        |");
    printf("|  \   --\   /  |");
    printf("|   \     \ /   |");
    printf("|    \-----\    |");
    printf("|           \   |");
    printf("|------------\--|");
}
```

Then Cakeman downloads the header file to a directory called `headers`

You will include and use it like this:
```c
#include "headers/sayhello/sayhello.h" 

// Main function
int main() {
    sayhi("Hello, world!");
    // Output:
    // <Hello, world!>
    //  \
    //   \
    //    \
    // |----------|
    // | *     *  |
    // |          |
    // | \      / |
    // |  \    /  |
    // |   \--/   |
    // |          |
    // |----------|
}
```
