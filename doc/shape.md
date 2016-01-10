# area

```js
var area = d3.svg.area()
    .x(function(d) { return x(d.date); })
    .y0(height)
    .y1(function(d) { return y(d.close); });
```

x与y0指定面积图下界曲线,x与y1指定上界曲线.

面积图通过将上面的area指定给path的d属性完成绘图:

```js
	svg.append('path')
				.datum(data)
				.attr('class', 'area')
				.attr('d', area);
```

# line

```js
	var line = d3.svg.line()
			.x(function (d) {
				return x(d.date);
			})
			.y(function (d) {
				return y(d.close);
			});

	svg.append("path")
				.datum(data)
				.attr("class", "line")
				.attr("d", line);
```

## line.interpolate([interpolate])

如果指定了interpolate 参数，就会设置插值器模式为指定的字符串或者函数。

如果没有指定，就返回当前的插值器模式。支持下面的命名插值器模式：

* linear -分段线性片段，如折线。
* linear-closed –闭合直线段，以形成一个多边形。
* step - 水平和垂直段之间交替，如在step函数中。
* step-before -垂直和水平段之间交替，如在step函数中。
* step-after -水平和垂直段之间交替，如在step函数中。
* basis - B样条曲线，在两端的控制点的重复。
* basis-open – 开放B样条曲线，首尾不会相交。
* basis-closed -封闭B样条曲线，如在一个循环。
* bundle – 等价于basis, 除了使用tension参数拉直样条曲线。
* cardinal – 基本样条曲线，在末端控制点的重复。
* cardinal-open –开放的基本样条曲线，首尾不会相交，但和其他控制点相交。
* cardinal-closed -封闭基本样条曲线，如在一个循环。
* monotone - 三次插值，可以保留y的单调性。


                                                                                                                                                                                                                                                                                                                                                                                                                                                • linear -分段线性片段，如折线。 • linear-closed –闭合直线段，以形成一个多边形。 • step - 水平和垂直段之间交替，如在step函数中。 • step-before -垂直和水平段之间交替，如在step函数中。 • step-after -水平和垂直段之间交替，如在step函数中。 • basis - B样条曲线，在两端的控制点的重复。 • basis-open – 开放B样条曲线，首尾不会相交。 • basis-closed -封闭B样条曲线，如在一个循环。 • bundle – 等价于basis, 除了使用tension参数拉直样条曲线。 • cardinal – 基本样条曲线，在末端控制点的重复。 • cardinal-open –开放的基本样条曲线，首尾不会相交，但和其他控制点相交。 • cardinal-closed -封闭基本样条曲线，如在一个循环。 • monotone - 三次插值，可以保留y的单调性。