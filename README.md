### Requirements
- [ ] You need to implement a cli tool that
  - [x] takes a list of urls as input
  - [x] visits each url and 
  - [x] prints to the output the list of pairs: 
    - [x] url and response body size. 
  - [x] the output must be sorted by the size of response body.

**Optional:**
- concurrent implementation

**Time limit:** 2 hours

**Programming Language:** go (only)

### Improvements
1. make the output look better:
- [go-pretty](https://github.com/jedib0t/go-pretty)
- [table](https://github.com/rodaine/table)

2. add tests

### Manual tests

```shell
./go-curl --urls hello.com blah http://artpav.dev google.com
./go-curl --urls youtube.com https://wodo.dev/ https://example.com 1
```
