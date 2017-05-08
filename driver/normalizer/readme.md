# Problems

## Tokens are the original full text, not just the symbol

For code like `$A` the bash parser returns a leaf like this:

```
{
  "children": [],
  "elementType": "[Bash] variable",
  "startOffset": 0,
  "text": "$A",
  "textLength": 2
}
```

The text is "$A", not "A" as you would expect from parsers of other
languages like Java.

## Tokens of parents hold the text of their children

For code like `$A` the bash parser returns a tree like this:

```
{
  "children": [
  {
    "children": [
    {
      "children": [
      {
        "children": [
        {
          "children": [],
          "elementType": "[Bash] variable",
          "startOffset": 0,
          "text": "$A",
          "textLength": 2
        }
        ],
        "elementType": "var-use-element",
        "startOffset": 0,
        "text": "$A",
        "textLength": 2
      }
      ],
      "elementType": "[Bash] combined word",
      "startOffset": 0,
      "text": "$A",
      "textLength": 2
    }
    ],
    "elementType": "[Bash] generic bash command",
    "startOffset": 0,
    "text": "$A",
    "textLength": 2
  }
  ],
  "elementType": "simple-command",
  "startOffset": 0,
  "text": "$A",
  "textLength": 2
},
```

Note how every node contains the text of their children.

This is particularly painfull in the case of the "FILE" node that contains the
full sourced text of the file being analysed.

