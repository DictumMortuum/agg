# agg

Aggregates yml

```
~/temp.2020-01-29.LZlUdLzbL
❯ cat a.yml
a: b
c:
  - 1
  - 2
  - 3
f: e

~/temp.2020-01-29.LZlUdLzbL
❯ cat b.yaml
b: "hello"
d: "there"

~/temp.2020-01-29.LZlUdLzbL
❯ agg
a: b
b: hello
c:
- 1
- 2
- 3
d: there
f: e
```