# jsonconsul

[![Build Status](https://travis-ci.org/vsco/jsonconsul.svg)](https://travis-ci.org/vsco/jsonconsul)

`jsonconsul` allows the creation of json files from the values that
are located in Consul. It also contains the ability to import KV keys
into Consul.

## Usage

### Import

A json config can be used to update key values in Consul.

Here is an example file named `example.json`
```
{"foo":{"bar":"test","blah":"Test","do":"TEST","loud":{"asd":{"bah":"test"}}}}
```

Now if we want to import this into the root prefix then we'd do the following:

```sh
jsonconsul import -json-file example.json
```

To import into an alternate prefix the following needs to be done:

```sh
jsonconsul import -prefix='vsco/buzz' -json-file example.json
```

### Export

There are five ways to run `jsonconsul export`. There are:

 - Output to STDOUT
 - Output to file
 - Output to file but timestamp the output file and symlink to name.
 - Poll and output to file after a duration.
 - Poll and output to file after a duration but with timestamped output file.

#### Output to STDOUT
```sh
jsonconsul export -prefix="foo"
```

#### Output to file
```sh
jsonconsul export -prefix="foo" -config=foo.json
```

#### Output to file with timestamp
```sh
jsonconsul export -prefix="foo" -config=foo.json -timestamp
```

This generates a file called `foo.json.<unixtimestamp>`. `foo.json`
will then be a symbolic link to `foo.json.<unixtimestamp>`.


#### Poll and output to file
```sh
jsonconsul watch -prefix="foo" -config=foo.json -poll
```

This polls consul every minute for changes and outputs those values to
json. If an alternate frequency is preferred then include the
`-poll_frequency` flag.

#### Poll and output to file with timestamp
```sh
jsonconsul watch -prefix="foo" -config=foo.json -timestamp -poll
```


### Diff

If you want to see how values in Consul differ from a json file that
you have this command will show you the differences between the local
json values and consul.

```sh
jsonconsul diff -prefix="foo" file.json
```

Or if you want to diff from global:

```sh
jsonconsul diff file.json
```

## License

The MIT License (MIT)

Copyright (c) 2015 Visual Supply, Co.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
