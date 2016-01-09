# 基本使用

```js
var xAxis = d3.svg.axis()
    .scale(x)
    .orient("bottom");

chart.append("g")
    .attr("class", "x axis")
    .attr("transform", "translate(0," + height + ")")
    .call(xAxis);
```

其中x是一个scale,比如`d3.scale.ordinal()`.

xAxis单独放到一个g中,并把它定位到适当的位置.

样式1:

```css
		.axis text {
			font: 10px sans-serif;
		}

		.axis path,
		.axis line {
			fill: none;
			stroke: #000;
			shape-rendering: crispEdges;
		}

		.x.axis path { /* path 就是最长的那条线 */
			display: none;
		}
```

# 格式化输出

```js
var yAxis = d3.svg.axis()
    .scale(y)
    .orient("left")
    .ticks(10, "%");
```

ticks 调用的是对应比例尺的 tickFormat, 比如这里的y是 d3.linear, 那么将调用 linear.ticksFormat.

格式化参数说明参考 `d3.format`