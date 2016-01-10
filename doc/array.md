# 最大值与最小值

```js
	var arr = [
		{name: 'a', age: 14},
		{name: 'b', age: 20},
		{name: 'c', age: 9},
		{name: 'd', age: 88}
	];
	var r = d3.extent(arr, function (d) {
		return d.age;
	});
	console.log(r); // [9 88]
```