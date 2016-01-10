# linear

```js
        var linear = d3.scale.linear()
                .domain([0, 10]) // 定义域，x轴范围
                .range([0, 100]); // 值域，y轴范围
        console.log(linear(0)); // 0
        console.log(linear(5)); // 50
        console.log(linear(10)); // 100
        console.log(linear(20)); // 200
```

类似于： `y = ax + b`

x超出了定义的范围也是可以的。

# ordinal

```js
        var ordinal = d3.scale.ordinal()
                .domain([1, 3, 5, 7])
                .range(['a', 'b', 'c', 'd']);
        console.log(ordinal(1)); // a
        console.log(ordinal(3)); // b
        console.log(ordinal(5)); // c
        console.log(ordinal(9)); // a
        console.log(ordinal(2)); // b
        console.log(ordinal(10)); // c
        console.log(ordinal(2)); // b
```

x,y 都是离散的.

当遇到不存在的x时,会将不存在的x加入到定义域中,值域为第一个y,之后再遇到不存在的,取下一个y.

if len(x) < len(y) :

```js
        var ordinal = d3.scale.ordinal()
                .domain([1, 3, 5])
                .range(['a', 'b', 'c', 'd']);
        console.log(ordinal(1)); // a
        console.log(ordinal(3)); // b
        console.log(ordinal(5)); // c
        console.log(ordinal(7)); // d
```

多出来的y自动分配给第一次出现的未定义x.

if len(y) < len(x) :

```js
        var ordinal = d3.scale.ordinal()
                .domain([1, 3, 5, 7, 9])
                .range(['a', 'b', 'c', 'd']);
        console.log(ordinal(1)); // a
        console.log(ordinal(3)); // b
        console.log(ordinal(5)); // c
        console.log(ordinal(7)); // d
        console.log(ordinal(9)); // a
        console.log(ordinal(1)); // a
        console.log(ordinal(2)); // b
```

多出来的9被分配了'a',再出现未定义的x则从'a'之后分配.

总结：未定义的x会被隐式的加入到定义域中，输出域是循环使用的。


## ordinal.rangePoints(interval[, padding])

```js
var o = d3.scale.ordinal()
    .domain([1, 2, 3, 4])
    .rangePoints([0, 100]);

o.range(); // [0, 33.333333333333336, 66.66666666666667, 100]

o.rangePoints([0, 120], 1);
o.range()   // [15, 45, 75, 105]    前面被空了 30*1/2=15 后面也被空了 15；其中 30 是间隔宽度

o.rangePoints([0, 120], 2);
o.range()   // [24, 48, 72, 96]    前面被空了 24*2/2=24 后面也被空了 24；其中 24 是间隔宽度

o.rangePoints([0, 120], 3);
o.range()   // [30, 50, 70, 90]    前面被空了 20*3/2=30 后面也被空了 30；其中 20 是间隔宽
```

间隔宽度的计算：

![](pic/01.png)

前后有 `step*padding/2`, 每个元素两两相隔 `step`, 总宽度为interval, 所以：

```
step * padding/2 * 2 + (dataLength - 1) * step = interval
=> step = interval / (padding + dataLength - 1)
```


## ordinal.rangeBands(interval[, padding[, outerPadding]])

```js
var o = d3.scale.ordinal()
    .domain([1, 2, 3])
    .rangeBands([0, 100]);

o.rangeBand(); // 33.333333333333336
o.range(); // [0, 33.333333333333336, 66.66666666666667]
o.rangeExtent(); // [0, 100]
```

![](pic/02.png)

rangeBand的计算：

首尾各一个 `step * outerPadding`, 除最后一个元素外其余元素都有一个 step, 最后一个元素只有rangeBand: `step*dataLength-step*padding`.

所以：

```
step * outerPadding * 2 + step * dataLength - step * padding = interval
=> step = interval / (outerPadding * 2 + dataLength - padding)

rangeBand = step * (1 - padding)
```

## color

颜色比例尺是序数比例尺的特例.

```js
	var color = d3.scale.category10();
	color.domain(['a','b','c','d']);
	console.log(color('a')); // #1f77b4
```


# 时间比例尺

```js
	var parseDate = d3.time.format('%Y-%m-%d').parse;
	var x = d3.time.scale()
			.domain([parseDate('2016-01-10'), parseDate('2016-02-01')])
			.range([0, 100]);
	['2016-01-10', '2016-01-20', '2016-01-31', '2016-02-01'].map(function (d) {
		console.log(x(parseDate(d)));
	});
	// 0
	// 45.45454545454545
	// 95.45454545454545
	// 100
```