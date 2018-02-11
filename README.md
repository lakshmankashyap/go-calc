# Calc.go
Given a JSON dictionary containing a `method` and a nested `point` dictionary containing `x` and `y`, outputs the result of running the method on x and y, as an expression.

Extra keys in the dictionary will be ignored.  Missing keys in the `point` will default to zero.  Missing `method` will default to `add`.

## Example

```
> calc someinput.json
5.500000 x -1.542000 = -8.481000
```

## Usage
```
calc <input file path>
```
The program will also read directly from standard input, e.g.:
```
echo '{"method":"add", "point": { "x": 5.5, "y": -1.542 } }' | calc
```
```
calc < somefile.json
```

#### Allowed methods

The allowed values for `method` are `add`, `subtract`, `multiply` and `divide`.

#### Example input

```
{
  "method": "add",
  "point": {
    "x": 1,
    "y": 2
  }
}
```
