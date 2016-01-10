# stack

```js
	var stack = d3.layout.stack()
			.values(function (d) {
				return d.values;
			});
	var data = [
		{
			name: 'A',
			values: [ // 第一层,a占10, b占90
				{x: 'a', y: 10},
				{x: 'b', y: 30}
			]
		},
		{
			name: 'B',
			values: [ // 第一层,a占30, b占70
				{x: 'a', y: 90},
				{x: 'b', y: 70}
			]
		}
	];
	stack(data);
	data.map(function(d){
		d.values.map(function(v){
			console.log(v);
		});
	});
	// Object {x: "a", y: 10, y0: 0}
	// Object {x: "b", y: 30, y0: 0}
	// Object {x: "a", y: 90, y0: 10}
	// Object {x: "b", y: 70, y0: 30}
```

values中的d是data[i], 返回的是一个数组,代表这一层的信息,每个元素都表示这一层中某一块区域的大小(用y指定).

`stack(data)` 计算每一层各区域的基准(即y0), 上面的例子中:

x=a的位置: 有两层,10和90,所以第二层的y0就是10

x=b的位置: 30和70,所以第二层的y0就是30

## stack配合area

```js
	var area = d3.svg.area()
			.x(function (d) {
				return x(d.date);
			})
			.y0(function (d) {
				return y(d.y0);
			})
			.y1(function (d) {
				return y(d.y0 + d.y);
			});
```

例子见: StackedAreaChart.html


# tree

```js
	var tree = d3.layout.tree()
			.size([300, 200]);
	var data = {
		name: 'A',
		children: [ // children是必须的
			{name: 'A-1'},
			{name: 'A-2'},
			{name: 'A-3'}
		]
	};
	// 执行后,data每个元素都有depth,x,y(此时的树是从上到下,而不是从左到右)
	// nodes是包含所有节点的数组
	var nodes = tree.nodes(data);
	// links中存储每个节点之间的关系,包含字段source,target
	var links = tree.links(nodes);
```

