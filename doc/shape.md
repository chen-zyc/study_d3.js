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

