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