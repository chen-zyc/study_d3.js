# selections

多选

```js
d3.selectAll("p").style("color", "white");
```

单选

```js
d3.select("body").style("background-color", "black");
```


# Dynamic Properties

```js
d3.selectAll("p").style("color", function() {
  return "hsl(" + Math.random() * 360 + ",100%,50%)";
});
```

```js
d3.selectAll("p").style("color", function(d, i) {
  return i % 2 ? "#fff" : "#eee";
});
```

d是第i个节点上绑定的数据。

```js
d3.selectAll("p")
    .data([4, 8, 15, 16, 23, 42])
    .style("font-size", function(d) { return d + "px"; });
```


# append()添加元素

```js
var body = d3.select("body");
var div = body.append("div");
div.html("Hello, world!");
```

```js
var section = d3.selectAll("section");
var div = section.append("div");
div.html("Hello, world!");
```

Note: append() returns a new selection containing the new elements.(chaining methods)

# 通过数据来生成元素

```js
d3.select(".chart")
  .selectAll("div")
    .data(data)
  .enter().append("div")
    .style("width", function(d) { return d * 10 + "px"; })
    .text(function(d) { return d; });
```




# Enter and Exit

Enter用于创建新的节点

```js
d3.select("body").selectAll("p")
    .data([4, 8, 15, 16, 23, 42])
  .enter().append("p")
    .text(function(d) { return "I’m number " + d + "!"; });
```

data有6个元素，如果p没有6个，那么enter()指代的就是没有数据的新创建的节点。

data()的结果就是选中了那些已匹配的节点。

```js
// Update…
var p = d3.select("body").selectAll("p")
    .data([4, 8, 15, 16, 23, 42])
    .text(function(d) { return d; });

// Enter…
p.enter().append("p")
    .text(function(d) { return d; });

// Exit…
p.exit().remove();
```

Exit用于删除节点，就是p的个数比data还多，就把多的删除掉。

## 用法1：

比如条形图，初始化时使用的是缩放比1，然后data添加了一些新的数据，此时就可以用enter()选中新的数据使用缩放比2.


# Transition

```js
d3.select("body").transition()
    .style("background-color", "black");
```

```js
d3.selectAll("circle").transition()
    .duration(750)
    .delay(function(d, i) { return i * 10; })
    .attr("r", function(d) { return Math.sqrt(d * scale); });
```


