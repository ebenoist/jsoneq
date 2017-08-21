Assert JSON Equal
---

Assert on two json strings that the lhs is equal to the rhs. Ignores extra keys in the rhs.

```
=== RUN   Test_ReturnsFalse
--- FAIL: Test_ReturnsFalse (0.00s)
	jsoneq.go:42:
		Expected:
		{
		  "a": [
		    2,
		    3,
		    1
		  ],
		  "bar": "foo",
		  "baz": "zed",
		  "foo": "zed",
		  "l": 2,
		  "persona": {
		    "id": 10,
		    "name": "Erik"
		  }
		}

		Actual:
		{
		  "a": [
		    3,
		    2,
		    1
		  ],
		  "foo": "bar",
		  "fooze": "baz",
		  "id": 12,
		  "l": 2,
		  "persona": {
		    "id": 12
		  }
		}

		Difference:
		- bar
		- baz
		- foo: zed
		+ foo: bar
		- a: [2 3 1]
		+ a: [3 2 1]
		- persona.id: 10
		+ persona.id: 12
		- persona.name
```
