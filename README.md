# gutool
golang cli tools for manipulating text.

## Introduction

In the spirit of small and focused command-line programs that comprise
the (\*)nix shell,
these utilities provide functions that I often find myself
cobbling together on-the-fly with `awk` and `sed`.

I already did some of this in C and python, so here it is again.

## upar

Format text into paragraphs.

### example

Here is some test data:

```
cat test_data.txt

When in the Course of human events, it
becomes necessary for one people to
dissolve the political bands which have
connected them with another, and to
assume among the powers of the earth,
the separate and equal station to which
the Laws of Nature and of Nature's God
entitle them, a decent respect to the
opinions of mankind requires that they
should declare the causes which impel
them to the separation.
```

Format the test data into paragraphs of lines not exceeding 80 characters:

```
upar < test_data.txt

When in the Course of human events, it becomes necessary for one people to
dissolve the political bands which have connected them with another, and to
assume among the powers of the earth, the separate and equal station to which
the Laws of Nature and of Nature's God entitle them, a decent respect to the
opinions of mankind requires that they should declare the causes which impel
them to the separation.
```

Format the test data into paragraphs of lines not exceeding 60 characters:

```
upar -width 60 < test_data.txt

When in the Course of human events, it becomes necessary for
one people to dissolve the political bands which have
connected them with another, and to assume among the powers
of the earth, the separate and equal station to which the
Laws of Nature and of Nature's God entitle them, a decent
respect to the opinions of mankind requires that they should
declare the causes which impel them to the separation.
```

Add an indent:

```
upar -indent 5 < test_data.txt

     When in the Course of human events, it becomes necessary for one people to
     dissolve the political bands which have connected them with another, and to
     assume among the powers of the earth, the separate and equal station to
     which the Laws of Nature and of Nature's God entitle them, a decent respect
     to the opinions of mankind requires that they should declare the causes
     which impel them to the separation.
```

Multi-paragraph:

```
upar -indent 5 -width 60 < test_data_2.txt

     When in the Course of human events, it becomes
     necessary for one people to dissolve the political
     bands which have connected them with another, and to
     assume among the powers of the earth, the separate and
     equal station to which the Laws of Nature and of
     Nature's God entitle them, a decent respect to the
     opinions of mankind requires that they should declare
     the causes which impel them to the separation.

     We hold these truths to be self-evident, that all men
     are created equal, that they are endowed by their
     Creator with certain unalienable Rights, that among
     these are Life, Liberty and the pursuit of Happiness.
```

If `indent` is not specified, then the indent of the
first line is used:

```
cat test_data_3.txt

    We hold these truths to be self-evident,
that all men are created equal,
        that they are endowed by their Creator with certain unalienable Rights,
  that among these are Life,
Liberty and the pursuit of Happiness.

upar < test_data_3.txt

    We hold these truths to be self-evident, that all men are created equal,
    that they are endowed by their Creator with certain unalienable Rights, that
    among these are Life, Liberty and the pursuit of Happiness.
```

### use with vi

From `command` mode in `vi`, the next 10 lines can be formatted
into a paragraph with this command:

```
:.,+9!upar -width 75
```

This is helpful for cleaning up comment blocks or formatting simple text files.

### syntax
```
upar [-indent indent] [-width width]
```

### options
```
Usage of upar:
  -indent int
    	left indent (default -1)
  -width int
    	max width of line (default 80)
 ```
