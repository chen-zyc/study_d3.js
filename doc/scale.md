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